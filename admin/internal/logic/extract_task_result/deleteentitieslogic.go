package extract_task_result

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/golang/glog"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteEntitiesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteEntitiesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteEntitiesLogic {
	return &DeleteEntitiesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteEntitiesLogic) DeleteEntities(req *types.DeleteEntitiesReq) (err error) {
	tx := l.svcCtx.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err = dao.DeleteRelationshipsByEntityIDs(tx, req.IDs)
	if err != nil {
		glog.Error(err)
		tx.Rollback()
		return
	}

	err = dao.DeletePropsByEntityIDs(tx, req.IDs)
	if err != nil {
		glog.Error(err)
		tx.Rollback()
		return
	}

	err = dao.DeleteEntitysByIDs(tx, req.IDs)
	if err != nil {
		glog.Error(err)
		tx.Rollback()
		return
	}

	err = tx.Commit().Error
	if err != nil {
		glog.Error(err)
		tx.Rollback()
		return
	}

	return
}
