package document

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/golang/glog"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteDocumentsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteDocumentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDocumentsLogic {
	return &DeleteDocumentsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteDocumentsLogic) DeleteDocuments(req *types.DeleteDocumentsReq) (err error) {
	err = dao.DeleteDocumentsByIDs(l.svcCtx.DB, req.IDs)
	if err != nil {
		glog.Error(err)
		return
	}

	return
}
