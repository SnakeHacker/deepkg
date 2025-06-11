package schema_triple

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSchemaTripleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSchemaTripleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSchemaTripleListLogic {
	return &GetSchemaTripleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSchemaTripleListLogic) GetSchemaTripleList(req *types.GetSchemaTripleListReq) (resp *types.GetSchemaTripleListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
