package extract_task_result

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/golang/glog"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePropLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePropLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePropLogic {
	return &UpdatePropLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePropLogic) UpdateProp(req *types.UpdatePropReq) (err error) {
	prop := req.Prop
	propModel, err := dao.SelectPropByID(l.svcCtx.DB, int(prop.ID))
	if err != nil {
		glog.Error(err)
		return
	}

	propModel.PropName = prop.PropName
	propModel.PropValue = prop.PropValue
	propModel.TaskID = int(prop.TaskID)

	err = dao.UpdateProp(l.svcCtx.DB, &propModel)
	if err != nil {
		glog.Error(err)
		return
	}

	return
}
