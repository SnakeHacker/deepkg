package response

import (
	"net/http"

	"github.com/SnakeHacker/deepkg/admin/common/errorx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type Body struct {
	Code    int         `json:"returnCode"`
	Msg     string      `json:"returnDesc"`
	Data    interface{} `json:"data,omitempty"`
	Success bool        `json:"success"`
}

func Response(w http.ResponseWriter, resp interface{}, err error) {
	var body Body
	if err != nil {
		body.Code = -1
		body.Msg = err.Error()

		if resp == errorx.UNAUTHORIZED_CODE {
			body.Code = errorx.UNAUTHORIZED_CODE
			httpx.WriteJson(w, http.StatusUnauthorized, body)
			return
		}

		body.Code = errorx.BADREQUEST_CODE
		httpx.WriteJson(w, http.StatusOK, body)
		return
	}

	body.Code = errorx.SUCCESS_CODE
	body.Msg = "OK"
	body.Data = resp
	body.Success = true
	w.Header().Set("content-type", "application/json; charset=utf-8")

	httpx.OkJson(w, body)
}
