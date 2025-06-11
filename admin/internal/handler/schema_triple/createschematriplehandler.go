package schema_triple

import (
	"net/http"

	"github.com/SnakeHacker/deepkg/admin/common/response"
	"github.com/SnakeHacker/deepkg/admin/internal/logic/schema_triple"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateSchemaTripleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateSchemaTripleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := schema_triple.NewCreateSchemaTripleLogic(r.Context(), svcCtx)
		err := l.CreateSchemaTriple(&req)
		response.Response(w, nil, err)

	}
}
