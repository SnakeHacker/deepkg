package extract_task

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishExtractTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishExtractTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishExtractTaskLogic {
	return &PublishExtractTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishExtractTaskLogic) PublishExtractTask(req *types.PublishExtractTaskReq) error {
	// todo: add your logic here and delete this line

	return nil
}
