package schema_ontology_prop

import (
	"net/http"

	"github.com/SnakeHacker/deepkg/admin/common/response"
	"github.com/SnakeHacker/deepkg/admin/internal/logic/schema_ontology_prop"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateSchemaOntologyPropHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateSchemaOntologyPropReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := schema_ontology_prop.NewCreateSchemaOntologyPropLogic(r.Context(), svcCtx)
		err := l.CreateSchemaOntologyProp(&req)
		response.Response(w, nil, err)

	}
}
