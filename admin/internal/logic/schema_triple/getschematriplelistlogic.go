package schema_triple

import (
	"context"
	"errors"

	"github.com/SnakeHacker/deepkg/admin/common"
	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/SnakeHacker/deepkg/admin/internal/model/gorm_model"
	"github.com/golang/glog"

	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSchemaTripleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSchemaTripleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSchemaTripleListLogic {
	return &GetSchemaTripleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSchemaTripleListLogic) GetSchemaTripleList(req *types.GetSchemaTripleListReq) (resp *types.GetSchemaTripleListResp, err error) {
	tripleModels, total, err := dao.SelectSchemaTriples(l.svcCtx.DB, int(req.WorkSpaceID), req.PageNumber, req.PageSize)
	if err != nil {
		glog.Error(err)
		return
	}

	creatorIDs := []int64{}
	for _, tripleModel := range tripleModels {
		creatorIDs = append(creatorIDs, int64(tripleModel.CreatorID))
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

	ontologyIDs := []int64{}
	for _, tripleModel := range tripleModels {
		ontologyIDs = append(ontologyIDs, int64(tripleModel.SourceOntologyID), int64(tripleModel.TargetOntologyID))
	}

	ontologyModels, err := dao.SelectSchemaOntologiesByIDs(l.svcCtx.DB, ontologyIDs)
	if err != nil {
		glog.Error(err)
		return
	}

	ontologyMap := make(map[int64]gorm_model.SchemaOntology)
	for _, ontologyModel := range ontologyModels {
		ontologyMap[int64(ontologyModel.ID)] = *ontologyModel
	}

	triples := []types.SchemaTriple{}
	for _, tripleModel := range tripleModels {
		creator, ok := userMap[int64(tripleModel.CreatorID)]
		if !ok {
			err = errors.New("user not found")
			glog.Error(err)
			return
		}

		sourceOntology, ok := ontologyMap[int64(tripleModel.SourceOntologyID)]
		if !ok {
			err = errors.New("source ontology not found")
			glog.Error(err)
			return
		}

		targetOntology, ok := ontologyMap[int64(tripleModel.TargetOntologyID)]
		if !ok {
			err = errors.New("target ontology not found")
			glog.Error(err)
			return
		}

		triple := types.SchemaTriple{
			ID:                 int64(tripleModel.ID),
			SourceOntologyID:   int64(tripleModel.SourceOntologyID),
			SourceOntologyName: sourceOntology.OntologyName,
			TargetOntologyID:   int64(tripleModel.TargetOntologyID),
			TargetOntologyName: targetOntology.OntologyName,
			Relationship:       tripleModel.Relationship,
			WorkSpaceID:        int64(tripleModel.WorkSpaceID),
			CreatorID:          int64(tripleModel.CreatorID),
			CreatorName:        creator.Username,
			CreatedAt:          tripleModel.CreatedAt.Format(common.TIME_FORMAT),
		}
		triples = append(triples, triple)
	}

	resp = &types.GetSchemaTripleListResp{
		Total:         total,
		SchemaTriples: triples,
		PageSize:      req.PageSize,
		PageNumber:    req.PageNumber,
	}

	return
}
