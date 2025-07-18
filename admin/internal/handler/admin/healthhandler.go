package admin

import (
	"net/http"

	"github.com/SnakeHacker/deepkg/admin/common/response"
	"github.com/SnakeHacker/deepkg/admin/internal/logic/admin"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
)

func HealthHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := admin.NewHealthLogic(r.Context(), svcCtx)
		resp, err := l.Health()
		response.Response(w, resp, err)

	}
}
