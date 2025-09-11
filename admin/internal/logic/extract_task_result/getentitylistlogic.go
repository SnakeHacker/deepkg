package extract_task_result

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEntityListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetEntityListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEntityListLogic {
	return &GetEntityListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetEntityListLogic) GetEntityList(req *types.GetEntityListReq) (resp *types.GetEntityListResp, err error) {

	entitiyModels, total, err := dao.SelectEntitiesByTaskID(l.svcCtx.DB, int(req.TaskID), req.PageNumber, req.PageSize)

	entities := []types.Entity{}
	for _, em := range entitiyModels {
		entity := types.Entity{
			ID:         int64(em.ID),
			EntityName: em.EntityName,
			TaskID:     int64(em.TaskID),
		}
		entities = append(entities, entity)
	}

	resp = &types.GetEntityListResp{
		Total:      total,
		Entities:   entities,
		PageSize:   req.PageSize,
		PageNumber: req.PageNumber,
	}

	return
}
