package document

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/golang/glog"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDocumentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDocumentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDocumentListLogic {
	return &GetDocumentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDocumentListLogic) GetDocumentList(req *types.GetDocumentListReq) (resp *types.GetDocumentListResp, err error) {
	docs, total, err := dao.SelectDocuments(l.svcCtx.DB, req.DirID, req.PageNumber, req.PageSize)
	if err != nil {
		glog.Error(err)
		return
	}

	resp = &types.GetDocumentListResp{
		Documents:  docs,
		PageNumber: req.PageNumber,
		PageSize:   req.PageSize,
		Total:      total,
	}

	return
}
