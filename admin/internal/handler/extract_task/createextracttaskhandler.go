package extract_task

import (
	"net/http"

	"github.com/SnakeHacker/deepkg/admin/internal/logic/extract_task"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
)

func CreateExtractTaskHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateExtractTaskReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := extract_task.NewCreateExtractTaskLogic(r.Context(), svcCtx)
		err := l.CreateExtractTask(&req)
		response.Response(w, nil, err)

	}
}
