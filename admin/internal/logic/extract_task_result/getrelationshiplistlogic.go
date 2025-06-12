package extract_task_result

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/SnakeHacker/deepkg/admin/internal/model/gorm_model"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/golang/glog"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRelationshipListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRelationshipListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRelationshipListLogic {
	return &GetRelationshipListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRelationshipListLogic) GetRelationshipList(req *types.GetRelationshipListReq) (resp *types.GetRelationshipListResp, err error) {
	rlModels, total, err := dao.SelectRelationships(l.svcCtx.DB, int(req.TaskID), req.PageNumber, req.PageSize)

	entityIDs := []int64{}
	for _, rlModel := range rlModels {
		entityIDs = append(entityIDs, int64(rlModel.SourceEntityID), int64(rlModel.TargetEntityID))
	}

	entityModels, err := dao.SelectEntityModelsByIDs(l.svcCtx.DB, entityIDs)
	if err != nil {
		glog.Error(err)
		return
	}
	entityMap := map[int64]gorm_model.Entity{}
	for _, entityModel := range entityModels {
		entityMap[int64(entityModel.ID)] = *entityModel
	}

	relationships := []types.Relationship{}
	for _, rlModel := range rlModels {
		relationship := types.Relationship{
			ID:               int64(rlModel.ID),
			RelationshipName: rlModel.RelationshipName,
			SourceEntityID:   int64(rlModel.SourceEntityID),
			SourceEntityName: entityMap[int64(rlModel.SourceEntityID)].EntityName,
			TargetEntityID:   int64(rlModel.TargetEntityID),
			TargetEntityName: entityMap[int64(rlModel.TargetEntityID)].EntityName,
		}

		relationships = append(relationships, relationship)
	}

	resp = &types.GetRelationshipListResp{
		Total:         total,
		Relationships: relationships,
		PageSize:      req.PageSize,
		PageNumber:    req.PageNumber,
	}

	return
}
