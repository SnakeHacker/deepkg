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

type UpdateSchemaOntologyPropLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSchemaOntologyPropLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSchemaOntologyPropLogic {
	return &UpdateSchemaOntologyPropLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSchemaOntologyPropLogic) UpdateSchemaOntologyProp(req *types.UpdateSchemaOntologyPropReq) (err error) {
	prop := req.SchemaOntologyProp
	propModel, err := dao.SelectSchemaOntologyPropByID(l.svcCtx.DB, prop.ID)
	if err != nil {
		glog.Error(err)
		return err
	}

	creator, err := l.svcCtx.GetUserFromCache(req.Authorization)
	if err != nil {
		glog.Error(err)
		return
	}

	propModel.PropName = prop.PropName
	propModel.PropDesc = prop.PropDesc
	propModel.CreatorID = int(creator.ID)

	err = dao.UpdateSchemaOntologyProp(l.svcCtx.DB, &propModel)
	if err != nil {
		glog.Error(err)
		return err
	}

	ontologyModel, err := dao.SelectSchemaOntologyByID(l.svcCtx.DB, int64(propModel.OntologyID))
	if err != nil {
		glog.Error(err)
		return err
	}

	workspaceModel, err := dao.SelectKnowledgeGraphWorkspaceByID(l.svcCtx.DB, int64(ontologyModel.WorkSpaceID))
	if err != nil {
		glog.Error("查询工作空间失败：", err)
		return
	}

	stmt := fmt.Sprintf("USE %s; ALTER TAG %s CHANGE (%s STRING COMMENT '%s');", workspaceModel.WorkSpaceName, ontologyModel.OntologyName, prop.PropName, prop.PropDesc)
	glog.Info("修改标签属性：", stmt)
	_, err = l.svcCtx.Nebula.Execute(stmt)
	if err != nil {
		glog.Error("修改标签属性失败：", err)
		return
	}

	return
}
