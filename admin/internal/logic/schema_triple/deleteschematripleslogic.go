package schema_triple

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/golang/glog"

	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSchemaTriplesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSchemaTriplesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSchemaTriplesLogic {
	return &DeleteSchemaTriplesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSchemaTriplesLogic) DeleteSchemaTriples(req *types.DeleteSchemaTriplesReq) (err error) {
	err = dao.DeleteSchemaTriplesByIDs(l.svcCtx.DB, req.IDs)
	if err != nil {
		glog.Error(err)
		return
	}
	return
}
