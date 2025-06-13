package knowledge_graph_workspace

import (
	"context"
	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/golang/glog"

	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateKnowledgeGraphWorkspaceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateKnowledgeGraphWorkspaceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateKnowledgeGraphWorkspaceLogic {
	return &UpdateKnowledgeGraphWorkspaceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateKnowledgeGraphWorkspaceLogic) UpdateKnowledgeGraphWorkspace(req *types.UpdateKnowledgeGraphWorkspaceReq) (err error) {
	workspace := req.KnowledgeGraphWorkspace
	workspaceModel, err := dao.SelectKnowledgeGraphWorkspaceByID(l.svcCtx.DB, workspace.ID)
	if err != nil {
		glog.Error(err)
		return
	}

	workspaceModel.WorkSpaceName = workspace.KnowledgeGraphWorkspaceName
	workspaceModel.CreatorID = int(workspace.CreatorID)

	err = dao.UpdateKnowledgeGraphWorkspace(l.svcCtx.DB, &workspaceModel)
	if err != nil {
		glog.Error(err)
		return
	}

	return
}
