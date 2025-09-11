package knowledge_graph_workspace

import (
	"context"
	"fmt"

	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/SnakeHacker/deepkg/admin/internal/model/gorm_model"
	"github.com/golang/glog"

	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateKnowledgeGraphWorkspaceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateKnowledgeGraphWorkspaceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateKnowledgeGraphWorkspaceLogic {
	return &CreateKnowledgeGraphWorkspaceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateKnowledgeGraphWorkspaceLogic) CreateKnowledgeGraphWorkspace(req *types.CreateKnowledgeGraphWorkspaceReq) (err error) {
	workspace := req.KnowledgeGraphWorkspace

	creator, err := l.svcCtx.GetUserFromCache(req.Authorization)
	if err != nil {
		glog.Error(err)
		return
	}

	workspaceModel := gorm_model.KnowledgeGraphWorkspace{
		WorkSpaceName: workspace.KnowledgeGraphWorkspaceName,
		CreatorID:     int(creator.ID),
	}

	err = dao.CreateKnowledgeGraphWorkspace(l.svcCtx.DB, &workspaceModel)
	if err != nil {
		glog.Error(err)
		return
	}

	_, err = l.svcCtx.Nebula.Execute(fmt.Sprintf("CREATE SPACE %s (vid_type=INT64, partition_num=1, replica_factor=1);", workspace.KnowledgeGraphWorkspaceName))
	glog.Infof("创建图空间：%s", workspace.KnowledgeGraphWorkspaceName)
	if err != nil {
		glog.Error("创建 Nebula 图空间失败:", err)
		return
	}

	return
}
