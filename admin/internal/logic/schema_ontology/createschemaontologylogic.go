package schema_ontology

import (
	"context"
	"fmt"

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

	creator, err := l.svcCtx.GetUserFromCache(req.Authorization)
	if err != nil {
		glog.Error(err)
		return
	}

	ontologyModel := gorm_model.SchemaOntology{
		OntologyName: ontology.OntologyName,
		OntologyDesc: ontology.OntologyDesc,
		WorkSpaceID:  int(ontology.WorkSpaceID),
		CreatorID:    int(creator.ID),
	}

	err = dao.CreateSchemaOntology(l.svcCtx.DB, &ontologyModel)
	if err != nil {
		glog.Error(err)
		return
	}

	workspaceModel, err := dao.SelectKnowledgeGraphWorkspaceByID(l.svcCtx.DB, ontology.WorkSpaceID)
	if err != nil {
		glog.Error("查询工作空间失败：", err)
		return
	}

	stmt := fmt.Sprintf("USE %s; CREATE TAG IF NOT EXISTS `%s`(name STRING NOT NULL DEFAULT '%s' COMMENT '名称') COMMENT = '%s';", workspaceModel.WorkSpaceName, ontology.OntologyName, ontology.OntologyName, ontology.OntologyDesc)
	glog.Info("创建标签：", stmt)
	_, err = l.svcCtx.Nebula.Execute(stmt)
	if err != nil {
		glog.Error("创建标签失败：", err)
		return
	}

	stmt = fmt.Sprintf("CREATE TAG INDEX IF NOT EXISTS `%s_index_name` ON `%s`(name(10));", ontology.OntologyName, ontology.OntologyName)
	glog.Infof("创建name索引：%s", stmt)
	_, err = l.svcCtx.Nebula.Execute(stmt)
	if err != nil {
		glog.Error("创建name索引失败：", err)
		return
	}

	return
}
