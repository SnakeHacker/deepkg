package document_dir

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/SnakeHacker/deepkg/admin/internal/model/gorm_model"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/golang/glog"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDocumentDirLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateDocumentDirLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDocumentDirLogic {
	return &CreateDocumentDirLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateDocumentDirLogic) CreateDocumentDir(req *types.CreateDocumentDirReq) (err error) {
	dd := req.DocumentDir
	err = dao.CreateDocumentDir(l.svcCtx.DB, &gorm_model.DocumentDir{
		DirName:   dd.DirName,
		ParentID:  dd.ParentID,
		SortIndex: dd.SortIndex,
		Remark:    dd.Remark,
	})

	if err != nil {
		glog.Error(err)
		return
	}

	return
}
