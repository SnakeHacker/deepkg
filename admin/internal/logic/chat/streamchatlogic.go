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

// 推理过程中的状态信息
type ReasoningState struct {
	Steps      []ReasoningStep `json:"steps"`      // 已执行的推理步骤
	Question   string          `json:"question"`   // 原始问题
	IsComplete bool            `json:"isComplete"` // 是否已完成推理
}

// 单个推理步骤
type ReasoningStep struct {
	StepNumber  int    `json:"stepNumber"`  // 步骤序号
	Description string `json:"description"` // 步骤描述
	Query       Query  `json:"query"`       // 执行的查询
	Result      string `json:"result"`      // 查询结果
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
	// 判断问题是否需要多步骤推理
	isComplex, reason, err := l.IsComplexQuestion(req.Content)
	if err != nil {
		glog.Error(err)
		return err
	}

	if isComplex {
		// 使用动态推理处理复杂问题
		glog.Info("使用动态推理处理复杂问题。原因：", reason)
		return l.DynamicReasoningStreamChat(req, ch)
	}

	glog.Info("使用单步查询处理简单。原因：", reason)
	// 原有的单步查询逻辑，适用于简单问题
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

func (l *StreamChatLogic) IsComplexQuestion(question string) (isComplex bool, reason string, err error) {
	sysPrompt := `你是知识图谱推理助手。你现在的任务是判断问题是否需要多步骤推理才能回答。
输出格式：{"isComplex": true/false, "reason": "原因说明"}`

	msgs := []llm.Message{
		{
			Role:    "system",
			Content: sysPrompt,
		},
	}

	result, err := l.svcCtx.LLM.Infer(question, llm.History{Messages: msgs})
	if err != nil {
		glog.Error(err)
		isComplex = false
		return
	}

	result = io_util.CleanJsonStr(result)
	var response struct {
		IsComplex bool   `json:"isComplex"`
		Reason    string `json:"reason"`
	}

	err = json.Unmarshal([]byte(result), &response)
	if err != nil {
		glog.Error(err)
		isComplex = false
		return
	}

	isComplex = response.IsComplex
	reason = response.Reason

	return
}

func (l *StreamChatLogic) DynamicReasoningStreamChat(req *types.StreamChatReq, ch chan<- string) error {
	// 获取知识图谱模式信息
	ontologies, triples, err := l.GetSchemaStr(req.WorkspaceID)
	if err != nil {
		return err
	}

	// 初始化推理状态
	state := ReasoningState{
		Question: req.Content,
		Steps:    []ReasoningStep{},
	}

	// 执行第一步推理
	firstStepQuery, firstStepDesc, err := l.PlanInitialStep(req.Content, ontologies, triples)
	if err != nil {
		return err
	}

	// 执行第一步查询
	firstStepResult, err := l.QuerySchema(req.WorkspaceID, firstStepQuery)
	if err != nil {
		return err
	}

	// 记录第一步结果
	state.Steps = append(state.Steps, ReasoningStep{
		StepNumber:  1,
		Description: firstStepDesc,
		Query:       firstStepQuery,
		Result:      firstStepResult,
	})

	// 最多执行5步推理，防止无限循环
	maxSteps := 5

	// 用于检测重复查询
	var lastQuery Query

	for i := 1; i < maxSteps; i++ {
		// 判断是否需要继续推理
		isComplete, nextQuery, nextDesc, err := l.PlanNextStep(state, ontologies, triples)
		if err != nil {
			return err
		}

		// 检测重复查询
		if i > 1 && queriesEqual(nextQuery, lastQuery) {
			glog.Info("检测到重复查询，停止推理")
			state.IsComplete = true
			break
		}

		// 如果推理完成，生成最终答案
		if isComplete || (len(nextQuery.Entities) == 0 && len(nextQuery.Relations) == 0) {
			state.IsComplete = true
			break
		}

		// 执行下一步查询
		nextResult, err := l.QuerySchema(req.WorkspaceID, nextQuery)
		if err != nil {
			return err
		}

		// 记录本步骤
		state.Steps = append(state.Steps, ReasoningStep{
			StepNumber:  i + 1,
			Description: nextDesc,
			Query:       nextQuery,
			Result:      nextResult,
		})

		// 如果查询结果为空，可能无法继续推理
		if strings.TrimSpace(nextResult) == "" {
			break
		}
	}

	// 生成最终答案
	return l.GenerateFinalAnswer(state, ch)
}

func (l *StreamChatLogic) PlanInitialStep(question string, ontologies string, triples string) (query Query, description string, err error) {
	sysPrompt := `你是一个知识图谱推理助手。你现在的任务是分析用户的问题并规划第一步查询。

实体：问题中明确提到的对象（如“姚明”“特斯拉”），需标注其在知识图谱中所属的本体（如人物、组织，需与知识图谱Schema一致）。
关系：问题中描述的实体间关联（如配偶、创始人），需标注关系类型。

输出格式为JSON:
{
  "isComplete": true/false,
  "query": {
    "entities": [{"ontology": "本体名称", "props": {"name": "实体名称", ...}}, ...],
    "relations": ["关系名称", ...]
  },
  "description": "下一步查询的目的描述"
}
字段说明：
isComplete是否已能回答问题的标志。如果为true，表示已能回答问题；如果为false，表示需要继续查询。
如果需要继续查询，query字段包含下一步查询的实体和关系信息。如果不需要继续查询，则query字段可以为空对象{}。
entities字段是一个数组，每个元素为要查询的实体对象，包含所属本体名称。若要根据实体的属性进行查询，则props字段包含属性键值对。若不需要查询实体属性，则props可以为空对象{}。
relations字段是一个数组，包含要查询的关系名称。如果不需要查询关系，则该数组可以为空。
description字段是对下一步查询目的的简要描述。
输出示例：
{
	"entities": [{"ontology": "人物", "props": {"name": "姚明"}}],
	"relations": ["配偶"]
}
`
	msgs := []llm.Message{
		{Role: "system", Content: sysPrompt},
	}

	prompt := fmt.Sprintf(`用户问题: %s
当前已有的本体以及属性，格式为：本体(属性名, 属性名, ...) 
%s
当前已有的关系三元组，格式为：源本体 -> 关系 -> 目标本体
%s

请分析该问题并设计第一步查询。`, question, ontologies, triples)

	result, err := l.svcCtx.LLM.Infer(prompt, llm.History{Messages: msgs})
	if err != nil {
		return Query{}, "", err
	}

	// 解析JSON响应
	var response struct {
		Query       Query  `json:"query"`
		Description string `json:"description"`
	}

	result = io_util.CleanJsonStr(result)
	err = json.Unmarshal([]byte(result), &response)
	if err != nil {
		return Query{}, "", err
	}

	return response.Query, response.Description, nil
}

func (l *StreamChatLogic) PlanNextStep(state ReasoningState, ontologies string, triples string) (isComplete bool, nextQuery Query, description string, err error) {
	sysPrompt := `你是一个知识图谱推理助手。你现在的任务是根据已有查询结果，决定是否需要继续查询，并规划下一步。

实体：问题中明确提到的对象（如“姚明”“特斯拉”），需标注其在知识图谱中所属的本体（如人物、组织，需与知识图谱Schema一致）。
关系：问题中描述的实体间关联（如配偶、创始人），需标注关系类型。

重要规则：
1. 如果连续两次执行相同或类似查询得到相似结果，说明数据库中没有更多信息，请标记为完成(isComplete=true)
2. 如果查询结果为空或很少，不要反复执行相同查询，应尝试不同查询或标记为完成
3. 当获得的结果已足够回答问题时，立即标记为完成
4. 最多执行3步查询，避免无限循环

输出格式为JSON:
{
  "isComplete": true/false,
  "query": {
    "entities": [{"ontology": "本体名称", "props": {"name": "实体名称", ...}}, ...],
    "relations": ["关系名称", ...]
  },
  "description": "下一步查询的目的描述"
}
字段说明：
isComplete是否已能回答问题的标志。如果为true，表示已能回答问题；如果为false，表示需要继续查询。
如果需要继续查询，query字段包含下一步查询的实体和关系信息。如果不需要继续查询，则query字段可以为空对象{}。
entities字段是一个数组，每个元素为要查询的实体对象，包含所属本体名称。若要根据实体的属性进行查询，则props字段包含属性键值对。若不需要查询实体属性，则props可以为空对象{}。
relations字段是一个数组，包含要查询的关系名称。如果不需要查询关系，则该数组可以为空。
description字段是对下一步查询目的的简要描述。
输出示例：
{
	"entities": [{"ontology": "人物", "props": {"name": "姚明"}}],
	"relations": ["配偶"]
}
`
	// 构造已执行步骤的描述
	stepsDesc := ""
	for _, step := range state.Steps {
		stepsDesc += fmt.Sprintf("步骤%d(%s):\n查询: %+v\n结果: %s\n\n",
			step.StepNumber, step.Description, step.Query, step.Result)
	}

	msgs := []llm.Message{
		{Role: "system", Content: sysPrompt},
	}

	prompt := fmt.Sprintf(`用户问题: %s
当前已有的本体以及属性，格式为：本体(属性名, 属性名, ...) 
%s
当前已有的关系三元组，格式为：源本体 -> 关系 -> 目标本体
%s
已执行的步骤:
%s

根据上述信息，分析是否已能回答问题，或需要继续查询。如需继续，规划下一步查询。`,
		state.Question, ontologies, triples, stepsDesc)

	result, err := l.svcCtx.LLM.Infer(prompt, llm.History{Messages: msgs})
	if err != nil {
		return false, Query{}, "", err
	}

	// 解析JSON响应
	var response struct {
		IsComplete  bool   `json:"isComplete"`
		Query       Query  `json:"query"`
		Description string `json:"description"`
	}

	result = io_util.CleanJsonStr(result)
	err = json.Unmarshal([]byte(result), &response)
	if err != nil {
		return false, Query{}, "", err
	}

	return response.IsComplete, response.Query, response.Description, nil
}

func (l *StreamChatLogic) GenerateFinalAnswer(state ReasoningState, ch chan<- string) error {
	sysPrompt := `你是一个知识图谱推理助手。你现在的任务是基于多步推理结果回答用户问题，并详细解释推理过程。如果无法回答问题，请明确说明。`

	// 构建推理步骤记录
	stepsDesc := ""
	for _, step := range state.Steps {
		stepsDesc += fmt.Sprintf("步骤%d(%s):\n查询结果: %s\n\n",
			step.StepNumber, step.Description, step.Result)
	}
	glog.Info(stepsDesc)

	msgs := []llm.Message{
		{Role: "system", Content: sysPrompt},
		{Role: "user", Content: fmt.Sprintf("问题: %s\n\n推理过程:\n%s\n\n请基于以上推理步骤回答问题。",
			state.Question, stepsDesc)},
	}

	// 流式返回最终答案
	return l.svcCtx.LLM.InferStream(l.ctx, state.Question, llm.History{Messages: msgs}, ch)
}

// 判断两个查询是否相同
func queriesEqual(q1, q2 Query) bool {
	// 检查实体数量是否相同
	if len(q1.Entities) != len(q2.Entities) {
		return false
	}

	// 检查关系数量是否相同
	if len(q1.Relations) != len(q2.Relations) {
		return false
	}

	// 检查关系是否相同
	for i, r := range q1.Relations {
		if r != q2.Relations[i] {
			return false
		}
	}

	// 简单比较每个实体的本体名称
	for _, e1 := range q1.Entities {
		found := false
		for _, e2 := range q2.Entities {
			if e1.Ontology == e2.Ontology {
				// 如果属性name存在，也比较name
				if name1, ok1 := e1.Props["name"]; ok1 {
					if name2, ok2 := e2.Props["name"]; ok2 && name1 == name2 {
						found = true
						break
					}
				} else {
					found = true
					break
				}
			}
		}
		if !found {
			return false
		}
	}

	return true
}
