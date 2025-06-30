package chat

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SnakeHacker/deepkg/admin/internal/logic/chat"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func StreamChatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StreamChatReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		// l := chat.NewStreamChatLogic(r.Context(), svcCtx)
		// resp, err := l.StreamChat(&req)
		// response.Response(w, resp, err)

		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
			return
		}

		ch := make(chan string)
		l := chat.NewStreamChatLogic(r.Context(), svcCtx)

		go func() {
			// 修改为传入完整的请求对象
			err := l.StreamChat(&req, ch)
			if err != nil {
				ch <- "error: " + err.Error()
			}
			close(ch)
		}()

		writer := bufio.NewWriter(w)
		for msg := range ch {
			// 按照 SSE 格式发送数据
			data := map[string]string{
				"result": msg,
			}
			jsonData, err := json.Marshal(data)
			if err != nil {
				continue
			}

			_, err = writer.WriteString(fmt.Sprintf("data: %s\n\n", jsonData))
			if err != nil {
				return
			}
			writer.Flush()
			flusher.Flush()
		}

	}
}
