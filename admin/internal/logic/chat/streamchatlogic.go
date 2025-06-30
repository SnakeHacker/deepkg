package chat

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/common/ai/llm"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StreamChatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStreamChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StreamChatLogic {
	return &StreamChatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StreamChatLogic) StreamChat(req *types.StreamChatReq, ch chan<- string) (err error) {
	sysPrompt := ``
	msgs := []llm.Message{}

	msgs = append(msgs, llm.Message{
		Role:    "system",
		Content: sysPrompt,
	}, llm.Message{
		Role:    "user",
		Content: req.Content,
	})

	for _, msg := range req.History {
		msgs = append(msgs, llm.Message{
			Role:    msg.Role,
			Content: msg.Content,
		})
	}

	return l.svcCtx.LLM.InferStream(l.ctx, req.Content, llm.History{
		Messages: msgs,
	}, ch)
}
