package extract_task

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/golang/glog"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateExtractTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateExtractTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateExtractTaskLogic {
	return &UpdateExtractTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateExtractTaskLogic) UpdateExtractTask(req *types.UpdateExtractTaskReq) (err error) {
	et := req.ExtractTask
	etModel, err := dao.SelectExtractTaskByID(l.svcCtx.DB, int(et.ID))
	if err != nil {
		glog.Error(err)
		return err
	}

	etModel.TaskName = et.TaskName
	etModel.Remark = et.Remark

	err = dao.UpdateExtractTask(l.svcCtx.DB, &etModel)
	if err != nil {
		glog.Error(err)
		return
	}

	return
}
