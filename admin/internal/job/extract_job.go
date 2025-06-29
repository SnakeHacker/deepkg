package job

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/SnakeHacker/deepkg/admin/common/ai/llm"
	"github.com/SnakeHacker/deepkg/admin/common/knowledge_graph"
	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/SnakeHacker/deepkg/admin/internal/model/gorm_model"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/utils/io_util"
	"github.com/golang/glog"
)

func DoExtractTask(svcCtx *svc.ServiceContext, taskID int) (err error) {

	// 载入文本内容
	contents, err := GetContents(svcCtx, taskID)
	if err != nil {
		glog.Error(err)
		return
	}

	// 载入三元组定义
	triples, err := GetTriples(svcCtx, taskID)
	if err != nil {
		glog.Error(err)
		return
	}

	glog.Infof("taskID: %v, triples size: %v", taskID, len(triples))

	// 载入本体
	ontologies, ontologyMap, ontologyNameMap, err := GetOntologies(svcCtx, triples)
	if err != nil {
		glog.Error(err)
		return
	}

	// 载入属性定义
	propsMap := map[int64][]gorm_model.SchemaOntologyProp{}
	for _, ontology := range ontologies {
		// 抽取本体属性
		props, err := GetOntologyProps(svcCtx, ontology)
		if err != nil {
			glog.Error(err)
			return err
		}
		propsMap[int64(ontology.ID)] = props
	}

	// 载入三元组字符串
	tripleContent, err := GetTriplesStr(svcCtx, triples, ontologyMap)
	if err != nil {
		glog.Error(err)
		return
	}

	// 开始抽取文本
	for _, content := range contents {
		tx := svcCtx.DB.Begin()
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
			}
		}()

		ontoligiesStrList := []string{}
		for _, ontology := range ontologies {
			ontoligiesStrList = append(ontoligiesStrList, ontology.OntologyName)
		}

		// 抽取实体
		prompt := ExtractEntityPrompt(content, ontoligiesStrList)
		// glog.Info(prompt)
		result, err := svcCtx.LLM.Infer(prompt, llm.History{})
		if err != nil {
			glog.Error(err)
			return err
		}

		glog.Info(result)

		result = io_util.CleanJsonStr(result)

		entities := knowledge_graph.Entities{}
		err = json.Unmarshal([]byte(result), &entities)
		if err != nil {
			glog.Error(err)
			return err
		}

		entitiesMap := make(map[string]gorm_model.Entity)
		// 实体入库
		for idx, et := range entities.Entities {
			etModel := &gorm_model.Entity{
				EntityName: et.EntityName,
				TaskID:     taskID,
			}
			err = dao.CreateEntity(tx, etModel)
			if err != nil {
				glog.Error(err)
				tx.Rollback()
				return err
			}

			entities.Entities[idx].ID = int(etModel.ID)
			entitiesMap[etModel.EntityName] = *etModel
		}

		// 抽取实体属性
		for _, entity := range entities.Entities {
			ontology, ok := ontologyNameMap[entity.Type]
			if !ok {
				glog.Warningf("ontology:[%v] is not existed", entity.Type)
				continue
			}
			glog.Infof("Extracting Entity: %v - Ontology: %v props...", entity.EntityName, entity.Type)

			propModels := propsMap[int64(ontology.ID)]
			propNameList := []string{}
			for _, prop := range propModels {
				propNameList = append(propNameList, prop.PropName)
			}

			if len(propNameList) == 0 {
				glog.Infof("Ontology: %v has not props...", entity.Type)
				continue
			}

			glog.Infof("Entity: %v, props: ", entity.EntityName, propNameList)

			prompt = ExtractPropsPrompt(content, entity.EntityName, propNameList)
			result, err = svcCtx.LLM.Infer(prompt, llm.History{})
			if err != nil {
				glog.Error(err)
				return err
			}

			glog.Info(result)

			result = io_util.CleanJsonStr(result)

			props := knowledge_graph.Props{}
			err = json.Unmarshal([]byte(result), &props)
			if err != nil {
				glog.Error(err)
				return err
			}

			glog.Info(props)

			for _, prop := range props.Props {
				propModel := &gorm_model.Prop{
					PropName:  prop.PropName,
					PropValue: prop.Value,
					EntityID:  entity.ID,
					TaskID:    taskID,
				}

				err = dao.CreateProp(tx, propModel)
				if err != nil {
					glog.Error(err)
					tx.Rollback()
					return err
				}
			}

		}

		// 抽取三元组
		prompt = ExtractRelationshipPrompt(content, entities.Entities, tripleContent)
		glog.Info(prompt)
		result, err = svcCtx.LLM.Infer(prompt, llm.History{})
		if err != nil {
			glog.Error(err)
			return err
		}
		glog.Info(result)

		result = io_util.CleanJsonStr(result)

		relationships := knowledge_graph.Relationships{}
		err = json.Unmarshal([]byte(result), &relationships)
		if err != nil {
			glog.Error(err)
			return err
		}

		// 关系入库
		for _, relationship := range relationships.Relationships {
			sourceEntity, ok := entitiesMap[relationship.Source]
			if !ok {
				err = errors.New("source entity not found")
				glog.Error(err)
				tx.Rollback()
				return err
			}

			targetEntity, ok := entitiesMap[relationship.Target]
			if !ok {
				err = errors.New("target entity not found")
				glog.Error(err)
				tx.Rollback()
				return err
			}

			relModel := &gorm_model.Relationship{
				SourceEntityID:   int(sourceEntity.ID),
				TargetEntityID:   int(targetEntity.ID),
				RelationshipName: relationship.Rel,
				TaskID:           taskID,
			}
			err = dao.CreateRelationship(tx, relModel)
			if err != nil {
				glog.Error(err)
				tx.Rollback()
				return err
			}

		}

		err = tx.Commit().Error
		if err != nil {
			glog.Error(err)
			tx.Rollback()
			return err
		}

	}

	return
}

func GetTriples(svcCtx *svc.ServiceContext, taskID int) (triples []*gorm_model.SchemaTriple, err error) {
	tripleModels, err := dao.SelectExtractTaskTriples(svcCtx.DB, taskID)
	if err != nil {
		glog.Error(err)
		return
	}

	tripleIDs := []int64{}
	for _, tripleModel := range tripleModels {
		tripleIDs = append(tripleIDs, int64(tripleModel.TripleID))
	}

	triples, err = dao.SelectSchemaTriplesByIDs(svcCtx.DB, tripleIDs)
	if err != nil {
		glog.Error(err)
		return
	}

	return
}

func GetContents(svcCtx *svc.ServiceContext, taskID int) (contents []string, err error) {
	// 抽取文档
	extractTaskDocModels, err := dao.SelectExtractTaskDocuments(svcCtx.DB, taskID)
	if err != nil {
		glog.Error(err)
		return
	}
	// Get Document Content
	docIDs := []int64{}
	for _, etdModel := range extractTaskDocModels {
		docIDs = append(docIDs, int64(etdModel.DocID))
	}

	docModels, err := dao.SelectDocumentModelByIDs(svcCtx.DB, docIDs)
	if err != nil {
		glog.Error(err)
		return
	}

	for _, docModel := range docModels {
		resp, err := http.Get(docModel.DocPath)
		if err != nil {
			glog.Error(err)
			continue
		}
		defer resp.Body.Close()

		docContent, err := io.ReadAll(resp.Body)
		if err != nil {
			glog.Error(err)
			continue
		}
		contents = append(contents, string(docContent))
	}

	return
}

func GetOntologies(svcCtx *svc.ServiceContext, triples []*gorm_model.SchemaTriple) (ontologies []gorm_model.SchemaOntology, ontologyMap map[int64]gorm_model.SchemaOntology, ontologyNameMap map[string]gorm_model.SchemaOntology, err error) {
	ontologyMap = make(map[int64]gorm_model.SchemaOntology)
	ontologyNameMap = make(map[string]gorm_model.SchemaOntology)
	ontologyIDs := []int64{}
	for _, triple := range triples {
		ontologyIDs = append(ontologyIDs, int64(triple.SourceOntologyID), int64(triple.TargetOntologyID))
	}

	ontologyModels, err := dao.SelectSchemaOntologiesByIDs(svcCtx.DB, ontologyIDs)
	if err != nil {
		glog.Error(err)
		return
	}
	for _, ontologyModel := range ontologyModels {
		ontologyMap[int64(ontologyModel.ID)] = *ontologyModel
		ontologyNameMap[ontologyModel.OntologyName] = *ontologyModel
	}

	for _, v := range ontologyMap {
		ontologies = append(ontologies, v)
	}

	return
}

func GetOntologyProps(svcCtx *svc.ServiceContext, ontology gorm_model.SchemaOntology) (props []gorm_model.SchemaOntologyProp, err error) {

	propModels, _, err := dao.SelectSchemaOntologyProps(svcCtx.DB, int(ontology.ID), -1, 0)
	if err != nil {
		glog.Error(err)
		return
	}

	for _, propModel := range propModels {
		props = append(props, *propModel)
	}

	return

}

func GetTriplesStr(svcCtx *svc.ServiceContext, triples []*gorm_model.SchemaTriple, ontologyMap map[int64]gorm_model.SchemaOntology) (tripleContent []string, err error) {

	for _, triple := range triples {
		sourceOntology, ok := ontologyMap[int64(triple.SourceOntologyID)]
		if !ok {
			err = errors.New("ontology not found")
			glog.Error(err)
			return
		}
		targetOntology, ok := ontologyMap[int64(triple.TargetOntologyID)]
		if !ok {
			err = errors.New("ontology not found")
			glog.Error(err)
			return
		}
		tripleContent = append(tripleContent,
			fmt.Sprintf("%s -> %s -> %s", sourceOntology.OntologyName, triple.Relationship, targetOntology.OntologyName))
	}

	return
}
