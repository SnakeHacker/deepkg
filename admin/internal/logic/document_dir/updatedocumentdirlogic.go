package document_dir

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/golang/glog"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDocumentDirLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateDocumentDirLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDocumentDirLogic {
	return &UpdateDocumentDirLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateDocumentDirLogic) UpdateDocumentDir(req *types.UpdateDocumentDirReq) (err error) {
	dir := req.DocumentDir
	dirModel, err := dao.SelectDocumentDirByID(l.svcCtx.DB, dir.ID)
	if err != nil {
		glog.Error(err)
		return
	}

	dirModel.DirName = dir.DirName
	dirModel.ParentID = dir.ParentID
	dirModel.SortIndex = dir.SortIndex
	dirModel.Remark = dir.Remark

	err = dao.UpdateDocumentDir(l.svcCtx.DB, &dirModel)
	if err != nil {
		glog.Error(err)
		return
	}

	return
}
