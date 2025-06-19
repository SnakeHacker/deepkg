package schema_ontology

import (
	"context"
	"errors"

	"github.com/SnakeHacker/deepkg/admin/common"
	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/SnakeHacker/deepkg/admin/internal/model/gorm_model"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/golang/glog"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSchemaOntologyListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSchemaOntologyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSchemaOntologyListLogic {
	return &GetSchemaOntologyListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSchemaOntologyListLogic) GetSchemaOntologyList(req *types.GetSchemaOntologyListReq) (resp *types.GetSchemaOntologyListResp, err error) {
	ontologyModels, total, err := dao.SelectSchemaOntologys(l.svcCtx.DB, int(req.WorkSpaceID), req.PageNumber, req.PageSize)
	if err != nil {
		glog.Error(err)
		return
	}

	creatorIDs := []int64{}
	for _, ontologyModel := range ontologyModels {
		creatorIDs = append(creatorIDs, int64(ontologyModel.CreatorID))
	}

	userModels, err := dao.SelectUserModelsByIDs(l.svcCtx.DB, creatorIDs)
	if err != nil {
		glog.Error(err)
		return
	}
	userMap := map[int64]gorm_model.User{}
	for _, userModel := range userModels {
		userMap[int64(userModel.ID)] = *userModel
	}

	ontologys := []types.SchemaOntology{}
	for _, ontologyModel := range ontologyModels {

		creator, ok := userMap[int64(ontologyModel.CreatorID)]
		if !ok {
			err = errors.New("user not found")
			glog.Error(err)
			return
		}

		ontologys = append(ontologys, types.SchemaOntology{
			ID:           int64(ontologyModel.ID),
			OntologyName: ontologyModel.OntologyName,
			OntologyDesc: ontologyModel.OntologyDesc,
			WorkSpaceID:  int64(ontologyModel.WorkSpaceID),
			CreatorID:    int64(creator.ID),
			CreatorName:  creator.Username,
			CreatedAt:    ontologyModel.CreatedAt.Format(common.TIME_FORMAT),
		})
	}

	resp = &types.GetSchemaOntologyListResp{
		Total:           total,
		PageSize:        req.PageSize,
		PageNumber:      req.PageNumber,
		SchemaOntologys: ontologys,
	}

	return
}
