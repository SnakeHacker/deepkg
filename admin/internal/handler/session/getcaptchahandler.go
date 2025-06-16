package session

import (
	"net/http"

	"github.com/SnakeHacker/deepkg/admin/common/response"
	"github.com/SnakeHacker/deepkg/admin/internal/logic/session"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
)

func GetCaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := session.NewGetCaptchaLogic(r.Context(), svcCtx)
		resp, err := l.GetCaptcha()
		response.Response(w, resp, err)

	}
}
