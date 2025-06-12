package extract_task

import (
	"net/http"

	"github.com/SnakeHacker/deepkg/admin/internal/logic/extract_task"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
)

func GetExtractTaskListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetExtractTaskListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := extract_task.NewGetExtractTaskListLogic(r.Context(), svcCtx)
		resp, err := l.GetExtractTaskList(&req)
		response.Response(w, resp, err)

	}
}
