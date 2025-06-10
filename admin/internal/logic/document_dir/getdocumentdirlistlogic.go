package document_dir

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/golang/glog"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDocumentDirListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDocumentDirListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDocumentDirListLogic {
	return &GetDocumentDirListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDocumentDirListLogic) GetDocumentDirList(req *types.GetDocumentDirListReq) (resp *types.GetDocumentDirListResp, err error) {
	dirModels, _, err := dao.SelectDocumentDirs(l.svcCtx.DB)
	if err != nil {
		glog.Error(err)
		return
	}

	dirMap := make(map[int64]*types.DocumentDir)
	for _, model := range dirModels {
		dir := types.DocumentDir{
			ID:        int64(model.ID),
			DirName:   model.DirName,
			ParentID:  model.ParentID,
			Children:  []types.DocumentDir{},
			SortIndex: model.SortIndex,
			Remark:    model.Remark,
		}
		dirMap[int64(model.ID)] = &dir
	}

	// 构建树形结构
	var rootDirs []types.DocumentDir
	for _, model := range dirModels {
		if model.ParentID == 0 {
			// 根目录
			rootDirs = append(rootDirs, *dirMap[int64(model.ID)])
		} else {
			if parent, ok := dirMap[model.ParentID]; ok {
				parent.Children = append(parent.Children, *dirMap[int64(model.ID)])
			}
		}
	}

	resp = &types.GetDocumentDirListResp{
		DocumentDirs: rootDirs,
	}

	return
}
