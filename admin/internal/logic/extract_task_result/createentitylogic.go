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

type CreateEntityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateEntityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateEntityLogic {
	return &CreateEntityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateEntityLogic) CreateEntity(req *types.CreateEntityReq) (err error) {
	entity := req.Entity
	entityModel := gorm_model.Entity{
		EntityName: entity.EntityName,
		TaskID:     int(entity.TaskID),
	}

	err = dao.CreateEntity(l.svcCtx.DB, &entityModel)
	if err != nil {
		glog.Error(err)
		return
	}

	return
}
