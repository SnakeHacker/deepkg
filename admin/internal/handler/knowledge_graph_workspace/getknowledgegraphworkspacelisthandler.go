package knowledge_graph_workspace

import (
	"net/http"

	"github.com/SnakeHacker/deepkg/admin/internal/logic/knowledge_graph_workspace"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetKnowledgeGraphWorkspaceListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetKnowledgeGraphWorkspaceListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := knowledge_graph_workspace.NewGetKnowledgeGraphWorkspaceListLogic(r.Context(), svcCtx)
		resp, err := l.GetKnowledgeGraphWorkspaceList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
