package schema_triple

import (
	"context"
	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/golang/glog"

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

func (l *UpdateSchemaTripleLogic) UpdateSchemaTriple(req *types.UpdateSchemaTripleReq) (err error) {
	triple := req.SchemaTriple
	tripleModel, err := dao.SelectSchemaTripleByID(l.svcCtx.DB, triple.ID)
	if err != nil {
		glog.Error(err)
		return
	}

	tripleModel.SourceOntologyID = int(triple.SourceOntologyID)
	tripleModel.TargetOntologyID = int(triple.TargetOntologyID)
	tripleModel.Relationship = triple.Relationship
	tripleModel.WorkSpaceID = int(triple.WorkSpaceID)
	tripleModel.CreatorID = int(triple.CreatorID)

	err = dao.UpdateSchemaTriple(l.svcCtx.DB, &tripleModel)
	if err != nil {
		glog.Error(err)
		return err
	}

	return nil
}
