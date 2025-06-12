package extract_task

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateExtractTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateExtractTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateExtractTaskLogic {
	return &UpdateExtractTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateExtractTaskLogic) UpdateExtractTask(req *types.UpdateExtractTaskReq) error {
	// todo: add your logic here and delete this line

	return nil
}
