package document

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/golang/glog"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDocumentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDocumentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDocumentLogic {
	return &GetDocumentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDocumentLogic) GetDocument(req *types.GetDocumentReq) (resp *types.GetDocumentResp, err error) {
	doc, err := dao.SelectDocumentByID(l.svcCtx.DB, req.ID)
	if err != nil {
		glog.Error(err)
		return
	}

	resp = &types.GetDocumentResp{
		Document: doc,
	}

	return
}
