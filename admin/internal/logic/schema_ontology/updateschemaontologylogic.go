package schema_ontology

import (
	"context"
	"fmt"

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

	creator, err := l.svcCtx.GetUserFromCache(req.Authorization)
	if err != nil {
		glog.Error(err)
		return
	}

	ontologyModel.OntologyName = ontology.OntologyName
	ontologyModel.OntologyDesc = ontology.OntologyDesc
	ontologyModel.WorkSpaceID = int(ontology.WorkSpaceID)
	ontologyModel.CreatorID = int(creator.ID)

	err = dao.UpdateSchemaOntology(l.svcCtx.DB, &ontologyModel)
	if err != nil {
		glog.Error(err)
		return err
	}

	workspaceModel, err := dao.SelectKnowledgeGraphWorkspaceByID(l.svcCtx.DB, ontology.WorkSpaceID)
	if err != nil {
		glog.Error("查询工作空间失败：", err)
		return
	}

	stmt := fmt.Sprintf("USE %s; ALTER TAG %s COMMENT = '%s';", workspaceModel.WorkSpaceName, ontology.OntologyName, ontology.OntologyDesc)
	glog.Info("修改标签：", stmt)
	_, err = l.svcCtx.Nebula.Execute(stmt)
	if err != nil {
		glog.Error("修改标签失败：", err)
		return
	}

	return
}
