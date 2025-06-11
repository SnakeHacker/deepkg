package schema_ontology

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/golang/glog"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSchemaOntologysLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSchemaOntologysLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSchemaOntologysLogic {
	return &DeleteSchemaOntologysLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSchemaOntologysLogic) DeleteSchemaOntologys(req *types.DeleteSchemaOntologysReq) (err error) {
	err = dao.DeleteSchemaOntologysByIDs(l.svcCtx.DB, req.IDs)
	if err != nil {
		glog.Error(err)
		return
	}

	return
}
