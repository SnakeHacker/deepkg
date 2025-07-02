package chat

import (
	"context"
	"errors"
	"fmt"
	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/SnakeHacker/deepkg/admin/internal/model/gorm_model"
	"github.com/golang/glog"

	"github.com/SnakeHacker/deepkg/admin/common/ai/llm"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

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
	triples, err := l.GetSchemaTripleList(req.WorkspaceID)
	if err != nil {
		glog.Error(err)
		return err
	}

	requirements, err := l.ExtractRequirements(req.Content, triples)
	if err != nil {
		glog.Error(err)
		return err
	}

	sysPrompt := ``
	msgs := []llm.Message{}

	msgs = append(msgs, llm.Message{
		Role:    "system",
		Content: sysPrompt,
	}, llm.Message{
		Role:    "user",
		Content: req.Content,
	})

	for _, msg := range req.History {
		msgs = append(msgs, llm.Message{
			Role:    msg.Role,
			Content: msg.Content,
		})
	}

	return l.svcCtx.LLM.InferStream(l.ctx, req.Content, llm.History{
		Messages: msgs,
	}, ch)
}

func (l *StreamChatLogic) ExtractRequirements(question string, triples string) (result string, err error) {
	sysPrompt := `你是一个知识图谱推理助手。你的任务是从用户的问题中总结出涉及到的、要查询的实体及关系。

任务说明：
从用户的问题中识别关键实体（具体对象）和关系（实体间的关联），输出为JSON格式。
实体：问题中明确提到的对象（如“姚明”“特斯拉”），需标注其在知识图谱中所属的本体（如人物、组织，需与知识图谱Schema一致）。
关系：问题中描述的实体间关联（如配偶、创始人），需标注关系类型。

输出格式要求：
输出严格的JSON对象，包含以下字段：
entities：数组，每个元素为实体对象，格式：{"name": "实体名称", "type": "实体类型"}。
relations：数组，每个元素为关系对象，格式：{"relation": "关系类型"}。

输入输出示例：
用户问题：姚明的妻子是谁？
当前已有的关系三元组（本体 -> 关系 -> 本体）：
人物 -> 配偶 -> 人物
人物 -> 创始人 -> 组织
输出：
{
	"entities": [{"name": "姚明", "type": "人物"}],
	"relations": [{"relation": "配偶"}]
}
`
	msgs := []llm.Message{}

	msgs = append(msgs, llm.Message{
		Role:    "system",
		Content: sysPrompt,
	})

	query := fmt.Sprintf(`用户问题：%s
当前已有的关系三元组（本体 -> 关系 -> 本体）：
%s
`, question, triples)

	return l.svcCtx.LLM.Infer(query, llm.History{
		Messages: msgs,
	})
}

func (l *StreamChatLogic) GetSchemaTripleList(workSpaceId int) (result string, err error) {
	tripleModels, _, err := dao.SelectSchemaTriples(l.svcCtx.DB, workSpaceId, 1, 100)
	if err != nil {
		glog.Error(err)
		return
	}

	ontologyIDs := []int64{}
	for _, tripleModel := range tripleModels {
		ontologyIDs = append(ontologyIDs, int64(tripleModel.SourceOntologyID), int64(tripleModel.TargetOntologyID))
	}

	ontologyModels, err := dao.SelectSchemaOntologiesByIDs(l.svcCtx.DB, ontologyIDs)
	if err != nil {
		glog.Error(err)
		return
	}

	ontologyMap := make(map[int64]gorm_model.SchemaOntology)
	for _, ontologyModel := range ontologyModels {
		ontologyMap[int64(ontologyModel.ID)] = *ontologyModel
	}

	result = ""
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

		result += fmt.Sprintf("%s -> %s -> %s\n", sourceOntology.OntologyName, tripleModel.Relationship, targetOntology.OntologyName)
	}

	return
}
