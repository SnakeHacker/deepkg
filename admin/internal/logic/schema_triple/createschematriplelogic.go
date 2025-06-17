package schema_triple

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/SnakeHacker/deepkg/admin/internal/model/gorm_model"
	"github.com/golang/glog"

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

func (l *CreateSchemaTripleLogic) CreateSchemaTriple(req *types.CreateSchemaTripleReq) (err error) {
	triple := req.SchemaTriple

	tripleModel := gorm_model.SchemaTriple{
		WorkSpaceID:      int(triple.WorkSpaceID),
		SourceOntologyID: int(triple.SourceOntologyID),
		TargetOntologyID: int(triple.TargetOntologyID),
		Relationship:     triple.Relationship,
		// TODO
		CreatorID: 1,
	}

	err = dao.CreateSchemaTriple(l.svcCtx.DB, &tripleModel)
	if err != nil {
		glog.Error(err)
		return
	}

	return
}
