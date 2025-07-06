package schema_triple

import (
	"context"
	"fmt"

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

	creator, err := l.svcCtx.GetUserFromCache(req.Authorization)
	if err != nil {
		glog.Error(err)
		return
	}

	tripleModel := gorm_model.SchemaTriple{
		WorkSpaceID:      int(triple.WorkSpaceID),
		SourceOntologyID: int(triple.SourceOntologyID),
		TargetOntologyID: int(triple.TargetOntologyID),
		Relationship:     triple.Relationship,
		CreatorID:        int(creator.ID),
	}

	err = dao.CreateSchemaTriple(l.svcCtx.DB, &tripleModel)
	if err != nil {
		glog.Error(err)
		return
	}

	sourceOntologyModel, err := dao.SelectSchemaOntologyByID(l.svcCtx.DB, triple.SourceOntologyID)
	if err != nil {
		glog.Error(err)
		return
	}

	targetOntologyModel, err := dao.SelectSchemaOntologyByID(l.svcCtx.DB, triple.TargetOntologyID)
	if err != nil {
		glog.Error(err)
		return
	}

	workspaceModel, err := dao.SelectKnowledgeGraphWorkspaceByID(l.svcCtx.DB, triple.WorkSpaceID)
	if err != nil {
		glog.Error("查询工作空间失败：", err)
		return
	}

	tripleStr := fmt.Sprintf("%s -> %s -> %s", sourceOntologyModel.OntologyName, triple.Relationship, targetOntologyModel.OntologyName)
	stmt := fmt.Sprintf("USE %s; CREATE EDGE IF NOT EXISTS %s() COMMENT = '%s';", workspaceModel.WorkSpaceName, triple.Relationship, tripleStr)
	glog.Info("创建边类型：", stmt)
	_, err = l.svcCtx.Nebula.Execute(stmt)
	if err != nil {
		glog.Error("创建边类型失败：", err)
		return
	}

	return
}
