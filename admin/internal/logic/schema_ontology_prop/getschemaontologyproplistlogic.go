package schema_ontology_prop

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

type GetSchemaOntologyPropListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSchemaOntologyPropListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSchemaOntologyPropListLogic {
	return &GetSchemaOntologyPropListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSchemaOntologyPropListLogic) GetSchemaOntologyPropList(req *types.GetSchemaOntologyPropListReq) (resp *types.GetSchemaOntologyPropListResp, err error) {
	propModels, total, err := dao.SelectSchemaOntologyProps(l.svcCtx.DB, int(req.OntologyID), req.PageNumber, req.PageSize)
	if err != nil {
		glog.Error(err)
		return
	}

	creatorIDs := []int64{}
	for _, propModel := range propModels {
		creatorIDs = append(creatorIDs, int64(propModel.CreatorID))
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

	props := []types.SchemaOntologyProp{}
	for _, propModel := range propModels {

		creator, ok := userMap[int64(propModel.CreatorID)]
		if !ok {
			err = errors.New("user not found")
			glog.Error(err)
			return
		}

		props = append(props, types.SchemaOntologyProp{
			ID:          int64(propModel.ID),
			PropName:    propModel.PropName,
			PropDesc:    propModel.PropDesc,
			OntologyID:  int64(propModel.OntologyID),
			CreatorID:   int64(creator.ID),
			CreatorName: creator.Username,
			CreatedAt:   propModel.CreatedAt.Format(common.TIME_FORMAT),
		})
	}

	resp = &types.GetSchemaOntologyPropListResp{
		Total:               total,
		PageSize:            req.PageSize,
		PageNumber:          req.PageNumber,
		SchemaOntologyProps: props,
	}

	return
}
