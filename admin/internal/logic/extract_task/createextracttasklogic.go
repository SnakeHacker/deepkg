package extract_task

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateExtractTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateExtractTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateExtractTaskLogic {
	return &CreateExtractTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateExtractTaskLogic) CreateExtractTask(req *types.CreateExtractTaskReq) error {
	// todo: add your logic here and delete this line

	return nil
}
