package extract_task

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/SnakeHacker/deepkg/admin/internal/model/gorm_model"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/golang/glog"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateExtractTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateExtractTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateExtractTaskLogic {
	return &CreateExtractTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateExtractTaskLogic) CreateExtractTask(req *types.CreateExtractTaskReq) (err error) {
	et := req.ExtractTask

	etModel := gorm_model.ExtractTask{
		TaskName:    et.TaskName,
		WorkSpaceID: et.WorkSpaceID,
		TaskStatus:  EXTRACT_TASK_STATUS_WAITING,
		Published:   false,
		Remark:      et.Remark,
		// TODO
		CreatorID: 0,
	}
	err = dao.CreateExtractTask(l.svcCtx.DB, &etModel)
	if err != nil {
		glog.Error(err)
		return
	}

	return
}
