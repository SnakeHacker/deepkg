package extract_task_result

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/golang/glog"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetExtractTaskResultLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetExtractTaskResultLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetExtractTaskResultLogic {
	return &GetExtractTaskResultLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetExtractTaskResultLogic) GetExtractTaskResult(req *types.GetExtractTaskResultReq) (resp *types.GetExtractTaskResultResp, err error) {

	entityModels, _, err := dao.SelectEntitys(l.svcCtx.DB, int(req.TaskID), -1, 10)
	if err != nil {
		glog.Error(err)
		return
	}

	entityIDs := []int{}
	for _, em := range entityModels {
		entityIDs = append(entityIDs, int(em.ID))
	}

	propModels, err := dao.SelectPropsByIDs(l.svcCtx.DB, entityIDs)
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

	nodes := []types.Entity{}
	for _, em := range entityModels {
		props, ok := propsMap[int(em.ID)]
		if !ok {
			props = []types.EntityProp{}
		}
		nodes = append(nodes, types.Entity{
			ID:         int64(em.ID),
			TaskID:     req.TaskID,
			EntityName: em.EntityName,
			Props:      props,
		})
	}

	edges := []types.Relationship{}

	relModels, _, err := dao.SelectRelationships(l.svcCtx.DB, int(req.TaskID), -1, 10)
	if err != nil {
		glog.Error(err)
		return
	}

	for _, rm := range relModels {
		edges = append(edges, types.Relationship{
			ID:               int64(rm.ID),
			RelationshipName: rm.RelationshipName,
			SourceEntityID:   int64(rm.SourceEntityID),
			TargetEntityID:   int64(rm.TargetEntityID),
		})
	}

	resp = &types.GetExtractTaskResultResp{
		ExtractTaskResult: types.ExtractTaskResult{
			TaskID:        req.TaskID,
			Entities:      nodes,
			Relationships: edges,
		},
	}

	return
}
