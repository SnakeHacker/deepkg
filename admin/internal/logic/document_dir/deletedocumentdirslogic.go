package document_dir

import (
	"context"
	"errors"

	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/golang/glog"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteDocumentDirsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteDocumentDirsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDocumentDirsLogic {
	return &DeleteDocumentDirsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteDocumentDirsLogic) DeleteDocumentDirs(req *types.DeleteDocumentDirsReq) (err error) {

	_, total, err := dao.SelectDocumentDirsByParentIDs(l.svcCtx.DB, req.IDs)
	if err != nil {
		glog.Error(err)
		return
	}

	if total > 0 {
		err = errors.New("directory is not empty")
		glog.Error(err)
		return
	}

	err = dao.DeleteDocumentDirsByIDs(l.svcCtx.DB, req.IDs)
	if err != nil {
		glog.Error(err)
		return
	}

	return
}
