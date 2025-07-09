package schema_ontology_prop

import (
	"context"
	"fmt"

	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/golang/glog"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSchemaOntologyPropsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSchemaOntologyPropsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSchemaOntologyPropsLogic {
	return &DeleteSchemaOntologyPropsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSchemaOntologyPropsLogic) DeleteSchemaOntologyProps(req *types.DeleteSchemaOntologyPropsReq) (err error) {
	propModels, err := dao.SelectSchemaOntologyPropsByIDs(l.svcCtx.DB, req.IDs)
	if err != nil {
		glog.Error(err)
		return
	}

	for _, propModel := range propModels {
		ontologyModel, e := dao.SelectSchemaOntologyByID(l.svcCtx.DB, int64(propModel.OntologyID))
		if e != nil {
			glog.Error(e)
			return e
		}

		workspaceModel, err := dao.SelectKnowledgeGraphWorkspaceByID(l.svcCtx.DB, int64(ontologyModel.WorkSpaceID))
		if err != nil {
			glog.Error("查询工作空间失败：", err)
			return err
		}

		stmt := fmt.Sprintf("USE %s; DROP TAG INDEX IF EXISTS `%s_index_%s`;", workspaceModel.WorkSpaceName, ontologyModel.OntologyName, propModel.PropName)
		glog.Infof("删除%s属性索引: %s", propModel.PropName, stmt)
		_, err = l.svcCtx.Nebula.Execute(stmt)
		if err != nil {
			glog.Error("删除属性索引失败:", err)
			return err
		}

		stmt = fmt.Sprintf("ALTER TAG `%s` DROP (`%s`);", ontologyModel.OntologyName, propModel.PropName)
		glog.Info("删除标签属性：", stmt)
		_, err = l.svcCtx.Nebula.Execute(stmt)
		if err != nil {
			glog.Error("删除标签属性失败：", err)
			return err
		}
	}

	err = dao.DeleteSchemaOntologyPropsByIDs(l.svcCtx.DB, req.IDs)
	if err != nil {
		glog.Error(err)
		return
	}

	return
}
