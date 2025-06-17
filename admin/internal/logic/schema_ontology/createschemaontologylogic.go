package schema_ontology

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/SnakeHacker/deepkg/admin/internal/model/gorm_model"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/golang/glog"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSchemaOntologyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateSchemaOntologyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSchemaOntologyLogic {
	return &CreateSchemaOntologyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateSchemaOntologyLogic) CreateSchemaOntology(req *types.CreateSchemaOntologyReq) (err error) {
	ontology := req.SchemaOntology

	ontologyModel := gorm_model.SchemaOntology{
		OntologyName: ontology.OntologyName,
		OntologyDesc: ontology.OntologyDesc,
		WorkSpaceID:  int(ontology.WorkSpaceID),
		// TODO
		CreatorID: 1,
	}

	err = dao.CreateSchemaOntology(l.svcCtx.DB, &ontologyModel)
	if err != nil {
		glog.Error(err)
		return
	}

	return
}
