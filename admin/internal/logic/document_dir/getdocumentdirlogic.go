package document_dir

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/golang/glog"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDocumentDirLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDocumentDirLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDocumentDirLogic {
	return &GetDocumentDirLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDocumentDirLogic) GetDocumentDir(req *types.GetDocumentDirReq) (resp *types.GetDocumentDirResp, err error) {
	dModel, err := dao.SelectDocumentDirByID(l.svcCtx.DB, req.ID)
	if err != nil {
		glog.Error(err)
		return
	}

	childrenDirModels, _, err := dao.SelectDocumentDirsByParentIDs(l.svcCtx.DB, []int64{req.ID})
	if err != nil {
		glog.Error(err)
		return
	}

	dir := types.DocumentDir{
		ID:        int64(dModel.ID),
		DirName:   dModel.DirName,
		ParentID:  dModel.ParentID,
		Children:  []types.DocumentDir{},
		SortIndex: dModel.SortIndex,
		Remark:    dModel.Remark,
	}

	for _, child := range childrenDirModels {
		dir.Children = append(dir.Children, types.DocumentDir{
			ID:        int64(child.ID),
			DirName:   child.DirName,
			ParentID:  child.ParentID,
			Children:  []types.DocumentDir{},
			SortIndex: child.SortIndex,
			Remark:    child.Remark,
		})
	}

	resp = &types.GetDocumentDirResp{
		DocumentDir: dir,
	}

	return
}
