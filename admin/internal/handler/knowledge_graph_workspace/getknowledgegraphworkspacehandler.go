package knowledge_graph_workspace

import (
	"net/http"

	"github.com/SnakeHacker/deepkg/admin/common/response"
	"github.com/SnakeHacker/deepkg/admin/internal/logic/knowledge_graph_workspace"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetKnowledgeGraphWorkspaceHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetKnowledgeGraphWorkspaceReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := knowledge_graph_workspace.NewGetKnowledgeGraphWorkspaceLogic(r.Context(), svcCtx)
		resp, err := l.GetKnowledgeGraphWorkspace(&req)
		response.Response(w, resp, err)

	}
}
