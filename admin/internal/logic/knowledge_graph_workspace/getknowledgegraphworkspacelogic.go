package knowledge_graph_workspace

import (
	"context"
	"github.com/SnakeHacker/deepkg/admin/common"
	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/golang/glog"

	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetKnowledgeGraphWorkspaceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetKnowledgeGraphWorkspaceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetKnowledgeGraphWorkspaceLogic {
	return &GetKnowledgeGraphWorkspaceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetKnowledgeGraphWorkspaceLogic) GetKnowledgeGraphWorkspace(req *types.GetKnowledgeGraphWorkspaceReq) (resp *types.GetKnowledgeGraphWorkspaceResp, err error) {
	workspaceModel, err := dao.SelectKnowledgeGraphWorkspaceByID(l.svcCtx.DB, req.ID)
	if err != nil {
		glog.Error(err)
		return
	}

	userModel, err := dao.SelectUserByID(l.svcCtx.DB, int64(workspaceModel.CreatorID))
	if err != nil {
		glog.Error(err)
		return
	}

	resp = &types.GetKnowledgeGraphWorkspaceResp{
		KnowledgeGraphWorkspace: types.KnowledgeGraphWorkspace{
			ID:                          int64(workspaceModel.ID),
			KnowledgeGraphWorkspaceName: workspaceModel.WorkSpaceName,
			CreatorID:                   int64(workspaceModel.CreatorID),
			CreatorName:                 userModel.Username,
			CreatedAt:                   workspaceModel.CreatedAt.Format(common.TIME_FORMAT),
		},
	}

	return
}
