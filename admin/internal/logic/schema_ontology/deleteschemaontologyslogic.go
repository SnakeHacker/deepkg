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

type DeleteSchemaOntologysLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSchemaOntologysLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSchemaOntologysLogic {
	return &DeleteSchemaOntologysLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSchemaOntologysLogic) DeleteSchemaOntologys(req *types.DeleteSchemaOntologysReq) (err error) {
	ontologyModels, err := dao.SelectSchemaOntologiesByIDs(l.svcCtx.DB, req.IDs)
	if err != nil {
		glog.Error(err)
		return
	}

	for _, ontology := range ontologyModels {
		workspaceModel, err := dao.SelectKnowledgeGraphWorkspaceByID(l.svcCtx.DB, int64(ontology.WorkSpaceID))
		if err != nil {
			glog.Error("查询工作空间失败：", err)
			return err
		}

		stmt := fmt.Sprintf("USE %s;", workspaceModel.WorkSpaceName)
		glog.Infof("选择图空间: %s", stmt)
		_, err = l.svcCtx.Nebula.Execute(stmt)
		if err != nil {
			glog.Error("选择图空间失败:", err)
			return err
		}

		propModels, err := dao.SelectSchemaOntologyPropsByOntologyIDs(l.svcCtx.DB, []int64{int64(ontology.ID)})
		if err != nil {
			glog.Error("查询本体属性失败：", err)
			return err
		}
		for _, prop := range propModels {
			stmt = fmt.Sprintf("DROP TAG INDEX IF EXISTS `%s_index_%s`;", ontology.OntologyName, prop.PropName)
			glog.Infof("删除%s属性索引: %s", prop.PropName, stmt)
			_, err = l.svcCtx.Nebula.Execute(stmt)
			if err != nil {
				glog.Errorf("删除%s属性索引失败:%s", prop.PropName, err)
				return err
			}
		}

		stmt = fmt.Sprintf("USE %s; DROP TAG IF EXISTS `%s`;", workspaceModel.WorkSpaceName, ontology.OntologyName)
		glog.Infof("删除标签: %s", stmt)
		_, err = l.svcCtx.Nebula.Execute(stmt)
		if err != nil {
			glog.Error("删除Nebula标签失败:", err)
			return err
		}
	}

	err = dao.DeleteSchemaOntologysByIDs(l.svcCtx.DB, req.IDs)
	if err != nil {
		glog.Error(err)
		return
	}

	return
}
