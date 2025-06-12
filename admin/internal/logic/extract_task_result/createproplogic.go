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

type CreatePropLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreatePropLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePropLogic {
	return &CreatePropLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreatePropLogic) CreateProp(req *types.CreatePropReq) (err error) {
	prop := req.Prop
	propModel := gorm_model.Prop{
		PropName:  prop.PropName,
		PropValue: prop.PropValue,
		EntityID:  int(prop.EntityID),
		TaskID:    int(prop.TaskID),
	}

	err = dao.CreateProp(l.svcCtx.DB, &propModel)
	if err != nil {
		glog.Error(err)
		return
	}

	return
}
