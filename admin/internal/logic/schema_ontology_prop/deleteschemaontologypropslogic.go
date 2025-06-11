package schema_ontology_prop

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/golang/glog"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSchemaOntologyPropsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSchemaOntologyPropsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSchemaOntologyPropsLogic {
	return &DeleteSchemaOntologyPropsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSchemaOntologyPropsLogic) DeleteSchemaOntologyProps(req *types.DeleteSchemaOntologyPropsReq) (err error) {
	err = dao.DeleteSchemaOntologyPropsByIDs(l.svcCtx.DB, req.IDs)
	if err != nil {
		glog.Error(err)
		return
	}

	return
}
