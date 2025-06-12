package extract_task_result

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/golang/glog"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateEntityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateEntityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateEntityLogic {
	return &UpdateEntityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateEntityLogic) UpdateEntity(req *types.UpdateEntityReq) (err error) {
	entity := req.Entity
	entityModel, err := dao.SelectEntityByID(l.svcCtx.DB, int(entity.TaskID))
	if err != nil {
		glog.Error(err)
		return
	}

	entityModel.EntityName = entity.EntityName
	entityModel.TaskID = int(entity.TaskID)

	err = dao.UpdateEntity(l.svcCtx.DB, &entityModel)
	if err != nil {
		glog.Error(err)
		return
	}

	return
}
