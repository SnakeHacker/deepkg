package knowledge_graph_workspace

import (
	"context"
	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/golang/glog"

	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteKnowledgeGraphWorkspacesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteKnowledgeGraphWorkspacesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteKnowledgeGraphWorkspacesLogic {
	return &DeleteKnowledgeGraphWorkspacesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteKnowledgeGraphWorkspacesLogic) DeleteKnowledgeGraphWorkspaces(req *types.DeleteKnowledgeGraphWorkspacesReq) (err error) {
	err = dao.DeleteKnowledgeGraphWorkspacesByIDs(l.svcCtx.DB, req.IDs)
	if err != nil {
		glog.Error(err)
		return
	}

	return
}
