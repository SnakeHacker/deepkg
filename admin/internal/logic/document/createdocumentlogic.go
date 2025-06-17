package document

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/SnakeHacker/deepkg/admin/internal/model/gorm_model"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/golang/glog"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDocumentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateDocumentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDocumentLogic {
	return &CreateDocumentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateDocumentLogic) CreateDocument(req *types.CreateDocumentReq) (err error) {

	doc := req.Document
	docModel := gorm_model.Document{
		DocName: doc.DocName,
		DocDesc: doc.DocDesc,
		DocPath: doc.DocPath,
		DirID:   int(doc.DirID),
		// TODO
		CreatorID: 1,
	}

	err = dao.CreateDocument(l.svcCtx.DB, &docModel)
	if err != nil {
		glog.Error(err)
		return
	}

	return
}
