package schema_ontology

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/golang/glog"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSchemaOntologyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSchemaOntologyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSchemaOntologyLogic {
	return &UpdateSchemaOntologyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSchemaOntologyLogic) UpdateSchemaOntology(req *types.UpdateSchemaOntologyReq) (err error) {
	ontology := req.SchemaOntology
	ontologyModel, err := dao.SelectSchemaOntologyByID(l.svcCtx.DB, ontology.ID)
	if err != nil {
		glog.Error(err)
		return err
	}

	ontologyModel.OntologyName = ontology.OntologyName
	ontologyModel.OntologyDesc = ontology.OntologyDesc
	ontologyModel.WorkSpaceID = int(ontology.WorkSpaceID)
	ontologyModel.CreatorID = int(ontology.CreatorID)

	err = dao.UpdateSchemaOntology(l.svcCtx.DB, &ontologyModel)
	if err != nil {
		glog.Error(err)
		return err
	}

	return
}
