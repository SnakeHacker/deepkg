package schema_triple

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSchemaTripleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateSchemaTripleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSchemaTripleLogic {
	return &CreateSchemaTripleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateSchemaTripleLogic) CreateSchemaTriple(req *types.CreateSchemaTripleReq) error {
	// todo: add your logic here and delete this line

	return nil
}
