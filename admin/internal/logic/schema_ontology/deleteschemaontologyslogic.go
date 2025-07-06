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

		stmt := fmt.Sprintf("USE %s; DROP TAG IF EXISTS %s;", workspaceModel.WorkSpaceName, ontology.OntologyName)
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
