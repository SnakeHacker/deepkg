package schema_triple

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSchemaTripleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSchemaTripleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSchemaTripleLogic {
	return &UpdateSchemaTripleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSchemaTripleLogic) UpdateSchemaTriple(req *types.UpdateSchemaTripleReq) error {
	// todo: add your logic here and delete this line

	return nil
}
