package schema_triple

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSchemaTripleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSchemaTripleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSchemaTripleLogic {
	return &GetSchemaTripleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSchemaTripleLogic) GetSchemaTriple(req *types.GetSchemaTripleReq) (resp *types.GetSchemaTripleResp, err error) {
	// todo: add your logic here and delete this line

	return
}
