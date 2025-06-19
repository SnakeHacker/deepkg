package schema_ontology_prop

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/SnakeHacker/deepkg/admin/internal/model/gorm_model"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/golang/glog"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSchemaOntologyPropLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateSchemaOntologyPropLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSchemaOntologyPropLogic {
	return &CreateSchemaOntologyPropLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateSchemaOntologyPropLogic) CreateSchemaOntologyProp(req *types.CreateSchemaOntologyPropReq) (err error) {
	prop := req.SchemaOntologyProp

	ontologyModel, err := dao.SelectSchemaOntologyByID(l.svcCtx.DB, prop.OntologyID)
	if err != nil {
		glog.Error(err)
		return
	}
	creator, err := l.svcCtx.GetUserFromCache(req.Authorization)
	if err != nil {
		glog.Error(err)
		return
	}

	ontologyPropModel := gorm_model.SchemaOntologyProp{
		PropName:    prop.PropName,
		PropDesc:    prop.PropDesc,
		WorkSpaceID: int(ontologyModel.WorkSpaceID),
		OntologyID:  int(prop.OntologyID),
		CreatorID:   int(creator.ID),
	}

	err = dao.CreateSchemaOntologyProp(l.svcCtx.DB, &ontologyPropModel)
	if err != nil {
		glog.Error(err)
		return
	}

	return
}
