package chat

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/SnakeHacker/deepkg/admin/internal/model/gorm_model"
	"github.com/SnakeHacker/deepkg/admin/internal/utils/io_util"
	"github.com/golang/glog"
	nebula_go "github.com/vesoft-inc/nebula-go/v3"
	"strings"

	"github.com/SnakeHacker/deepkg/admin/common/ai/llm"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Query struct {
	Entities  []EntityQuery `json:"entities"`
	Relations []string      `json:"relations"` // 关系名称列表
}

type EntityQuery struct {
	Ontology string            `json:"ontology"` // 本体名称
	Props    map[string]string `json:"props"`    // 属性键值对
}

type StreamChatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStreamChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StreamChatLogic {
	return &StreamChatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StreamChatLogic) StreamChat(req *types.StreamChatReq, ch chan<- string) (err error) {
	ontologies, triples, err := l.GetSchemaStr(req.WorkspaceID)
	if err != nil {
		glog.Error(err)
		return err
	}
	glog.Infof("ontologies: %s", ontologies)
	glog.Infof("triples: %s", triples)

	requirements, err := l.ExtractRequirements(req.Content, ontologies, triples)
	if err != nil {
		glog.Error(err)
		return err
	}
	requirements = io_util.CleanJsonStr(requirements)
	glog.Infof("requirements: %s", requirements)

	var query Query
	err = json.Unmarshal([]byte(requirements), &query)
	if err != nil {
		glog.Error(err)
		return
	}

	queryResult, err := l.QuerySchema(req.WorkspaceID, query)
	if err != nil {
		glog.Error(err)
		return
	}
	glog.Infof("queryResult: %s", queryResult)

	sysPrompt := `你是一个知识图谱推理助手。你现在的任务是根据查询到的实体和关系结果，回答用户的问题。如果查询到的结果无法回答用户的问题，请直接回复“无法回答”。`
	msgs := []llm.Message{}

	msgs = append(msgs, llm.Message{
		Role:    "system",
		Content: sysPrompt,
	})

	for _, msg := range req.History {
		msgs = append(msgs, llm.Message{
			Role:    msg.Role,
			Content: msg.Content,
		})
	}

	msgs = append(msgs, llm.Message{
		Role:    "user",
		Content: fmt.Sprintf("用户问题：%s\n查询到的实体和关系结果：%s", req.Content, queryResult),
	})

	return l.svcCtx.LLM.InferStream(l.ctx, req.Content, llm.History{
		Messages: msgs,
	}, ch)
}

func (l *StreamChatLogic) ExtractRequirements(question string, ontologies string, triples string) (result string, err error) {
	sysPrompt := `你是一个知识图谱推理助手。你现在的任务是从用户的问题中总结出涉及到的、要查询的实体及关系。

任务说明：
从用户的问题中识别关键实体（具体对象）和关系（实体间的关联），输出为JSON格式。
实体：问题中明确提到的对象（如“姚明”“特斯拉”），需标注其在知识图谱中所属的本体（如人物、组织，需与知识图谱Schema一致）。
关系：问题中描述的实体间关联（如配偶、创始人），需标注关系类型。

输出格式要求：
输出严格的JSON对象，包含以下字段：
entities：数组，每个元素为要查询的属于给出本体的实体对象，包含所属本体名称、实体名称、其他属性，格式：{"ontology": "所属本体", "props": {"name": "实体名称", "属性名称": "属性值", ...}}。
relations：数组，每个元素为要查询的关系名称，若不需要查询关系则数组为空。

输入输出示例：

用户问题：姚明的妻子是谁？
当前已有的本体以及属性，格式为：本体(属性名, 属性名, ...)
人物(name, 年龄, 性别)
组织(name, 成立时间, 创始人)
当前已有的关系三元组，格式为：源本体 -> 关系 -> 目标本体
人物 -> 配偶 -> 人物
人物 -> 创始人 -> 组织
输出：
{
	"entities": [{"ontology": "人物", "props": {"name": "姚明"}}],
	"relations": ["配偶"]
}

用户问题：当前图空间里有哪些人物？
当前已有的本体以及属性，格式为：本体(属性名, 属性名, ...)
人物(name, 年龄, 性别)
组织(name, 成立时间, 创始人)
当前已有的关系三元组，格式为：源本体 -> 关系 -> 目标本体
人物 -> 配偶 -> 人物
人物 -> 创始人 -> 组织
输出：
{
	"entities": [{"ontology": "人物", "props": {}}],
	"relations": []
}
`
	msgs := []llm.Message{}

	msgs = append(msgs, llm.Message{
		Role:    "system",
		Content: sysPrompt,
	})

	query := fmt.Sprintf(`用户问题：%s
当前已有的本体以及属性，格式为：本体(属性名, 属性名, ...)
%s
当前已有的关系三元组，格式为：源本体 -> 关系 -> 目标本体
%s
`, question, ontologies, triples)

	return l.svcCtx.LLM.Infer(query, llm.History{
		Messages: msgs,
	})
}

func (l *StreamChatLogic) GetSchemaStr(workSpaceId int) (ontologies string, triples string, err error) {
	ontologyModels, _, err := dao.SelectSchemaOntologys(l.svcCtx.DB, workSpaceId, -1, -1)
	if err != nil {
		glog.Error(err)
		return
	}

	ontologies = ""
	var ontologyPropModels []*gorm_model.SchemaOntologyProp
	for _, ontologyModel := range ontologyModels {
		ontologyPropModels, _, err = dao.SelectSchemaOntologyProps(l.svcCtx.DB, int(ontologyModel.ID), -1, -1)
		if err != nil {
			glog.Error(err)
			return
		}

		ontologies += fmt.Sprintf("%s(name", ontologyModel.OntologyName)
		for _, propModel := range ontologyPropModels {
			ontologies += fmt.Sprintf(", %s", propModel.PropName)
		}
		ontologies += ")\n"
	}

	tripleModels, _, err := dao.SelectSchemaTriples(l.svcCtx.DB, workSpaceId, -1, -1)
	if err != nil {
		glog.Error(err)
		return
	}

	ontologyIDs := []int64{}
	for _, tripleModel := range tripleModels {
		ontologyIDs = append(ontologyIDs, int64(tripleModel.SourceOntologyID), int64(tripleModel.TargetOntologyID))
	}

	ontologyMap := make(map[int64]gorm_model.SchemaOntology)
	for _, ontologyModel := range ontologyModels {
		ontologyMap[int64(ontologyModel.ID)] = *ontologyModel
	}

	triples = ""
	for _, tripleModel := range tripleModels {
		sourceOntology, ok := ontologyMap[int64(tripleModel.SourceOntologyID)]
		if !ok {
			err = errors.New("source ontology not found")
			glog.Error(err)
			return
		}

		targetOntology, ok := ontologyMap[int64(tripleModel.TargetOntologyID)]
		if !ok {
			err = errors.New("target ontology not found")
			glog.Error(err)
			return
		}

		triples += fmt.Sprintf("%s -> %s -> %s\n", sourceOntology.OntologyName, tripleModel.Relationship, targetOntology.OntologyName)
	}

	return
}

func (l *StreamChatLogic) QuerySchema(workspaceID int, query Query) (result string, err error) {
	workspaceModel, err := dao.SelectKnowledgeGraphWorkspaceByID(l.svcCtx.DB, int64(workspaceID))
	if err != nil {
		glog.Error(err)
		return
	}

	// 切换到指定图空间
	stmt := fmt.Sprintf("USE %s;", workspaceModel.WorkSpaceName)
	glog.Infof("选择图空间：%s", stmt)
	_, err = l.svcCtx.Nebula.Execute(stmt)
	if err != nil {
		glog.Error(err)
		return
	}

	// 查询实体
	entitiesResult := ""
	if len(query.Entities) > 0 {
		for _, entity := range query.Entities {
			stmt = fmt.Sprintf("MATCH (v:`%s` ", entity.Ontology)
			if len(entity.Props) > 0 {
				stmt += "{"
				propsStr := []string{}
				for propName, propValue := range entity.Props {
					propsStr = append(propsStr, fmt.Sprintf("`%s`: '%s'", propName, propValue))
				}
				stmt += strings.Join(propsStr, ", ")
				stmt += "}"
			}
			stmt += ") RETURN v;"
			glog.Infof("查询实体：%s", stmt)

			var resultSet *nebula_go.ResultSet
			resultSet, err = l.svcCtx.Nebula.Execute(stmt)
			if err != nil {
				glog.Error(err)
				return
			}
			rowSize := resultSet.GetRowSize()
			for i := 0; i < rowSize; i++ {
				var record *nebula_go.Record
				record, err = resultSet.GetRowValuesByIndex(i)
				if err != nil {
					glog.Error(err)
					return
				}
				var value *nebula_go.ValueWrapper
				value, err = record.GetValueByIndex(0)
				if err != nil {
					glog.Error(err)
					return
				}
				var valueVertex *nebula_go.Node
				valueVertex, err = value.AsNode()
				if err != nil {
					glog.Error(err)
					return
				}
				entitiesResult += fmt.Sprintf("%s\n", valueVertex.String())
			}
		}
	}
	glog.Infof("查询到的实体结果：%s", entitiesResult)

	// 查询关系
	relationsResult := ""
	if len(query.Relations) > 0 && len(query.Entities) > 0 {
		for _, entity := range query.Entities {
			// 作为源实体
			stmt = fmt.Sprintf("MATCH p=(v:`%s` ", entity.Ontology)
			if len(entity.Props) > 0 {
				stmt += "{"
				propsStr := []string{}
				for propName, propValue := range entity.Props {
					propsStr = append(propsStr, fmt.Sprintf("`%s`: '%s'", propName, propValue))
				}
				stmt += strings.Join(propsStr, ", ")
				stmt += "}"
			}
			stmt += fmt.Sprintf(")-[e:`%s`", query.Relations[0])
			if len(query.Relations) > 1 {
				for _, relation := range query.Relations[1:] {
					stmt += fmt.Sprintf("|`%s`", relation)
				}
			}
			stmt += fmt.Sprintf("]->() RETURN p;")
			glog.Infof("查询关系：%s", stmt)

			var resultSet *nebula_go.ResultSet
			resultSet, err = l.svcCtx.Nebula.Execute(stmt)
			if err != nil {
				glog.Error(err)
				return
			}
			rowSize := resultSet.GetRowSize()
			for i := 0; i < rowSize; i++ {
				var record *nebula_go.Record
				record, err = resultSet.GetRowValuesByIndex(i)
				if err != nil {
					glog.Error(err)
					return
				}
				var value *nebula_go.ValueWrapper
				value, err = record.GetValueByIndex(0)
				if err != nil {
					glog.Error(err)
					return
				}
				var valuePath *nebula_go.PathWrapper
				valuePath, err = value.AsPath()
				if err != nil {
					glog.Error(err)
					return
				}
				relationsResult += fmt.Sprintf("%s\n", valuePath.String())
			}

			// 作为目标实体
			stmt = fmt.Sprintf("MATCH p=()-[e:`%s`", query.Relations[0])
			if len(query.Relations) > 1 {
				for _, relation := range query.Relations[1:] {
					stmt += fmt.Sprintf("|`%s`", relation)
				}
			}
			stmt += fmt.Sprintf("]->(v:`%s`", entity.Ontology)
			if len(entity.Props) > 0 {
				stmt += "{"
				propsStr := []string{}
				for propName, propValue := range entity.Props {
					propsStr = append(propsStr, fmt.Sprintf("`%s`: '%s'", propName, propValue))
				}
				stmt += strings.Join(propsStr, ", ")
				stmt += "}"
			}
			stmt += ") RETURN p;"
			glog.Infof("查询关系：%s", stmt)

			resultSet, err = l.svcCtx.Nebula.Execute(stmt)
			if err != nil {
				glog.Error(err)
				return
			}
			rowSize = resultSet.GetRowSize()
			for i := 0; i < rowSize; i++ {
				var record *nebula_go.Record
				record, err = resultSet.GetRowValuesByIndex(i)
				if err != nil {
					glog.Error(err)
					return
				}
				var value *nebula_go.ValueWrapper
				value, err = record.GetValueByIndex(0)
				if err != nil {
					glog.Error(err)
					return
				}
				var valuePath *nebula_go.PathWrapper
				valuePath, err = value.AsPath()
				if err != nil {
					glog.Error(err)
					return
				}
				relationsResult += fmt.Sprintf("%s\n", valuePath.String())
			}
		}
	} else if len(query.Relations) > 0 && len(query.Entities) == 0 {
		// 如果没有指定实体，则查询所有关系
		stmt = fmt.Sprintf("MATCH p=()-[e:`%s`", query.Relations[0])
		if len(query.Relations) > 1 {
			for _, relation := range query.Relations[1:] {
				stmt += fmt.Sprintf("|`%s`", relation)
			}
		}
		stmt += "]->() RETURN p;"

		var resultSet *nebula_go.ResultSet
		resultSet, err = l.svcCtx.Nebula.Execute(stmt)
		if err != nil {
			glog.Error(err)
			return
		}
		rowSize := resultSet.GetRowSize()
		for i := 0; i < rowSize; i++ {
			var record *nebula_go.Record
			record, err = resultSet.GetRowValuesByIndex(i)
			if err != nil {
				glog.Error(err)
				return
			}
			var value *nebula_go.ValueWrapper
			value, err = record.GetValueByIndex(0)
			if err != nil {
				glog.Error(err)
				return
			}
			var valuePath *nebula_go.PathWrapper
			valuePath, err = value.AsPath()
			if err != nil {
				glog.Error(err)
				return
			}
			relationsResult += fmt.Sprintf("%s\n", valuePath.String())
		}
	}

	result = ""
	if entitiesResult != "" {
		result += fmt.Sprintf("查询到的实体：\n%s", entitiesResult)
		result += "\n"
	}
	if relationsResult != "" {
		result += fmt.Sprintf("查询到的关系：\n%s", relationsResult)
	}

	return
}
