package schema_triple

import (
	"net/http"

	"github.com/SnakeHacker/deepkg/admin/common/response"
	"github.com/SnakeHacker/deepkg/admin/internal/logic/schema_triple"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteSchemaTriplesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteSchemaTriplesReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := schema_triple.NewDeleteSchemaTriplesLogic(r.Context(), svcCtx)
		err := l.DeleteSchemaTriples(&req)
		response.Response(w, nil, err)

	}
}
