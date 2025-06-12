package extract_task

import (
	"net/http"

	"github.com/SnakeHacker/deepkg/admin/internal/logic/extract_task"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
)

func PublishExtractTaskHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PublishExtractTaskReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := extract_task.NewPublishExtractTaskLogic(r.Context(), svcCtx)
		err := l.PublishExtractTask(&req)
		response.Response(w, nil, err)

	}
}
