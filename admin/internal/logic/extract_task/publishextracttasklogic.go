package extract_task

import (
	"context"
	"fmt"
	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/golang/glog"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishExtractTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishExtractTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishExtractTaskLogic {
	return &PublishExtractTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishExtractTaskLogic) PublishExtractTask(req *types.PublishExtractTaskReq) (err error) {
	taskModel, err := dao.SelectExtractTaskByID(l.svcCtx.DB, int(req.ID))
	if err != nil {
		glog.Error("查询抽取任务失败：", err)
		return
	}

	workspaceModel, err := dao.SelectKnowledgeGraphWorkspaceByID(l.svcCtx.DB, taskModel.WorkSpaceID)
	if err != nil {
		glog.Error("查询工作空间失败：", err)
		return
	}

	stmt := fmt.Sprintf("USE %s;", workspaceModel.WorkSpaceName)
	glog.Info("使用图空间:", stmt)
	_, err = l.svcCtx.Nebula.Execute(stmt)
	if err != nil {
		glog.Error("使用图空间失败:", err)
		return err
	}

	entityModels, _, err := dao.SelectEntitiesByTaskID(l.svcCtx.DB, int(req.ID), -1, 10)
	if err != nil {
		glog.Error(err)
		return
	}

	entityIDs := []int{}
	for _, em := range entityModels {
		entityIDs = append(entityIDs, int(em.ID))
	}

	propModels, err := dao.SelectPropsByEntityIDs(l.svcCtx.DB, entityIDs)
	if err != nil {
		glog.Error(err)
		return
	}

	propsMap := make(map[int][]types.EntityProp)
	for _, pm := range propModels {
		propsMap[int(pm.EntityID)] = append(propsMap[int(pm.EntityID)], types.EntityProp{
			ID:        int64(pm.ID),
			TaskID:    int64(pm.TaskID),
			EntityID:  int64(pm.EntityID),
			PropName:  pm.PropName,
			PropValue: pm.PropValue,
		})
	}

	// 创建点
	for _, em := range entityModels {
		props, ok := propsMap[int(em.ID)]
		if !ok {
			props = []types.EntityProp{}
		}

		ontology, e := dao.SelectSchemaOntologyByID(l.svcCtx.DB, int64(em.OntologyID))
		if e != nil {
			glog.Error(e)
			return e
		}

		tagName := ontology.OntologyName

		propNameList := "name"
		for _, prop := range props {
			propNameList = propNameList + ", " + prop.PropName
		}

		VID := em.ID

		propValueList := "\"" + em.EntityName + "\""
		for _, prop := range props {
			propValueList = propValueList + ", \"" + prop.PropValue + "\""
		}

		stmt := fmt.Sprintf("INSERT VERTEX %s (%s) VALUES %d:(%s)", tagName, propNameList, VID, propValueList)
		glog.Info("创建点:", stmt)
		_, err := l.svcCtx.Nebula.Execute(stmt)
		if err != nil {
			glog.Error("发布点到 Nebula 失败:", err)
			return err
		}
	}

	// 创建边
	relModels, _, err := dao.SelectRelationshipsByTaskID(l.svcCtx.DB, int(req.ID), -1, 10)
	if err != nil {
		glog.Error(err)
		return
	}

	for _, rm := range relModels {
		edgeType := rm.RelationshipName
		srcVID := rm.SourceEntityID
		dstVID := rm.TargetEntityID

		stmt := fmt.Sprintf("INSERT EDGE %s () VALUES %d->%d:()", edgeType, srcVID, dstVID)
		glog.Info("创建边:", stmt)
		_, err = l.svcCtx.Nebula.Execute(stmt)
		if err != nil {
			glog.Error("发布边到 Nebula 失败:", err)
			return err
		}
	}

	err = dao.UpdateExtractTaskPublished(l.svcCtx.DB, int(req.ID), true)
	if err != nil {
		glog.Error(err)
		return
	}

	return
}
