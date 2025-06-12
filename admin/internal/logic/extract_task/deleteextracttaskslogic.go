package extract_task

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteExtractTasksLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteExtractTasksLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteExtractTasksLogic {
	return &DeleteExtractTasksLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteExtractTasksLogic) DeleteExtractTasks(req *types.DeleteExtractTasksReq) error {
	// todo: add your logic here and delete this line

	return nil
}
