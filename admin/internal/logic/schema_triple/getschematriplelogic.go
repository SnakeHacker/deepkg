package schema_triple

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/common"
	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/golang/glog"

	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSchemaTripleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSchemaTripleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSchemaTripleLogic {
	return &GetSchemaTripleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSchemaTripleLogic) GetSchemaTriple(req *types.GetSchemaTripleReq) (resp *types.GetSchemaTripleResp, err error) {
	tripleModel, err := dao.SelectSchemaTripleByID(l.svcCtx.DB, req.ID)
	if err != nil {
		glog.Error(err)
		return
	}

	userModel, err := dao.SelectUserByID(l.svcCtx.DB, int64(tripleModel.CreatorID))
	if err != nil {
		glog.Error(err)
		return
	}

	resp = &types.GetSchemaTripleResp{
		SchemaTriple: types.SchemaTriple{
			ID:               int64(tripleModel.ID),
			SourceOntologyID: int64(tripleModel.SourceOntologyID),
			TargetOntologyID: int64(tripleModel.TargetOntologyID),
			Relationship:     tripleModel.Relationship,
			WorkSpaceID:      int64(tripleModel.WorkSpaceID),
			CreatorID:        userModel.ID,
			CreatorName:      userModel.Username,
			CreatedAt:        tripleModel.CreatedAt.Format(common.TIME_FORMAT),
		},
	}

	return
}
