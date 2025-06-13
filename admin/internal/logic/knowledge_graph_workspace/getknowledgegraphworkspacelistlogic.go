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

type GetKnowledgeGraphWorkspaceListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetKnowledgeGraphWorkspaceListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetKnowledgeGraphWorkspaceListLogic {
	return &GetKnowledgeGraphWorkspaceListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetKnowledgeGraphWorkspaceListLogic) GetKnowledgeGraphWorkspaceList(req *types.GetKnowledgeGraphWorkspaceListReq) (resp *types.GetKnowledgeGraphWorkspaceListResp, err error) {
	workspaceModels, total, err := dao.SelectKnowledgeGraphWorkspaces(l.svcCtx.DB, int(req.CreatorID), req.PageNumber, req.PageSize)
	if err != nil {
		glog.Error(err)
		return
	}

	userModel, err := dao.SelectUserModelByID(l.svcCtx.DB, int64(int(req.CreatorID)))
	if err != nil {
		glog.Error(err)
		return
	}

	workspaces := []types.KnowledgeGraphWorkspace{}
	for _, workspaceModel := range workspaceModels {
		workspaces = append(workspaces, types.KnowledgeGraphWorkspace{
			ID:                          int64(workspaceModel.ID),
			KnowledgeGraphWorkspaceName: workspaceModel.WorkSpaceName,
			CreatorID:                   int64(workspaceModel.CreatorID),
			CreatorName:                 userModel.Username,
			CreatedAt:                   workspaceModel.CreatedAt.Format(common.TIME_FORMAT),
		})
	}

	resp = &types.GetKnowledgeGraphWorkspaceListResp{
		Total:                    total,
		PageNumber:               req.PageNumber,
		PageSize:                 req.PageSize,
		KnowledgeGraphWorkspaces: workspaces,
	}

	return
}
