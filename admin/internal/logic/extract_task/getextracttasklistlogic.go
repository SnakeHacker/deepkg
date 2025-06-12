package extract_task

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetExtractTaskListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetExtractTaskListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetExtractTaskListLogic {
	return &GetExtractTaskListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetExtractTaskListLogic) GetExtractTaskList(req *types.GetExtractTaskListReq) (resp *types.GetExtractTaskListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
