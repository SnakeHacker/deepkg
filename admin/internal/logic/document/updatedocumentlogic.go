package document

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/golang/glog"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDocumentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateDocumentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDocumentLogic {
	return &UpdateDocumentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateDocumentLogic) UpdateDocument(req *types.UpdateDocumentReq) (err error) {
	doc := req.Document
	docModel, err := dao.SelectDocumentModelByID(l.svcCtx.DB, req.Document.ID)
	if err != nil {
		glog.Error(err)
		return
	}

	docModel.DocName = doc.DocName
	docModel.DocDesc = doc.DocDesc
	docModel.DocPath = doc.DocPath
	docModel.DirID = int(doc.DirID)

	err = dao.UpdateDocument(l.svcCtx.DB, docModel)
	if err != nil {
		glog.Error(err)
		return
	}

	return
}
