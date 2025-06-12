package extract_task

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateExtractTaskStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateExtractTaskStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateExtractTaskStatusLogic {
	return &UpdateExtractTaskStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateExtractTaskStatusLogic) UpdateExtractTaskStatus(req *types.UpdateExtractTaskStatusReq) error {
	// todo: add your logic here and delete this line

	return nil
}
