package extract_task

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetExtractTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetExtractTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetExtractTaskLogic {
	return &GetExtractTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetExtractTaskLogic) GetExtractTask(req *types.GetExtractTaskReq) (resp *types.GetExtractTaskResp, err error) {
	// todo: add your logic here and delete this line

	return
}
