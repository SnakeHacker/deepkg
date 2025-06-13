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
	tx := l.svcCtx.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

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
	err = dao.CreateExtractTask(tx, &etModel)
	if err != nil {
		glog.Error(err)
		tx.Rollback()
		return
	}

	for _, docID := range et.DocIDs {
		etd := gorm_model.ExtractTaskDocument{
			TaskID: int(etModel.ID),
			DocID:  int(docID),
		}
		err = dao.CreateExtractTaskDocument(tx, &etd)
		if err != nil {
			glog.Error(err)
			tx.Rollback()
			return
		}
	}

	for _, tripleID := range et.TripleIDs {
		ett := gorm_model.ExtractTaskTriple{
			TaskID:   int(etModel.ID),
			TripleID: int(tripleID),
		}

		err = dao.CreateExtractTaskTriple(tx, &ett)
		if err != nil {
			glog.Error(err)
			tx.Rollback()
			return
		}
	}

	err = tx.Commit().Error
	if err != nil {
		glog.Error(err)
		tx.Rollback()
		return
	}

	// TODO(mickey): 后台异步执行抽取任务

	return
}
