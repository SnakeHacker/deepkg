package knowledge_graph_workspace

import (
	"net/http"

	"github.com/SnakeHacker/deepkg/admin/internal/logic/knowledge_graph_workspace"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateKnowledgeGraphWorkspaceHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateKnowledgeGraphWorkspaceReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := knowledge_graph_workspace.NewUpdateKnowledgeGraphWorkspaceLogic(r.Context(), svcCtx)
		err := l.UpdateKnowledgeGraphWorkspace(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
