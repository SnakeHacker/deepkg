package extract_task

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/SnakeHacker/deepkg/admin/internal/job"
	"github.com/SnakeHacker/deepkg/admin/internal/model/gorm_model"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/golang/glog"

	"github.com/zeromicro/go-zero/core/logx"
)

type RunExtractTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRunExtractTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RunExtractTaskLogic {
	return &RunExtractTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RunExtractTaskLogic) RunExtractTask(req *types.RunExtractTaskReq) (err error) {
	taskModel, err := dao.SelectExtractTaskByID(l.svcCtx.DB, int(req.ID))
	if err != nil {
		glog.Error(err)
		return
	}
	taskModel.TaskStatus = EXTRACT_TASK_STATUS_RUNNING
	glog.Infof("Run extract task: 【%v】%v", taskModel.ID, taskModel.TaskName)

	err = dao.UpdateExtractTask(l.svcCtx.DB, &taskModel)
	if err != nil {
		glog.Error(err)
		return
	}

	go func(tm gorm_model.ExtractTask) {
		err = job.DoExtractTask(l.svcCtx, int(req.ID))
		if err != nil {
			glog.Error(err)
			tm.TaskStatus = EXTRACT_TASK_STATUS_FAILED
			err = dao.UpdateExtractTask(l.svcCtx.DB, &tm)
			if err != nil {
				glog.Error(err)
				return
			}

			return
		}

		tm.TaskStatus = EXTRACT_TASK_STATUS_SUCCESSED
		err = dao.UpdateExtractTask(l.svcCtx.DB, &tm)
		if err != nil {
			glog.Error(err)
			return
		}

		return

	}(taskModel)

	return
}
