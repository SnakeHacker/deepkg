package schema_triple

import (
	"context"
	"fmt"

	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/golang/glog"

	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSchemaTripleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSchemaTripleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSchemaTripleLogic {
	return &UpdateSchemaTripleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSchemaTripleLogic) UpdateSchemaTriple(req *types.UpdateSchemaTripleReq) (err error) {
	triple := req.SchemaTriple
	tripleModel, err := dao.SelectSchemaTripleByID(l.svcCtx.DB, triple.ID)
	if err != nil {
		glog.Error(err)
		return
	}

	tripleModel.SourceOntologyID = int(triple.SourceOntologyID)
	tripleModel.TargetOntologyID = int(triple.TargetOntologyID)
	tripleModel.Relationship = triple.Relationship
	tripleModel.WorkSpaceID = int(triple.WorkSpaceID)

	err = dao.UpdateSchemaTriple(l.svcCtx.DB, &tripleModel)
	if err != nil {
		glog.Error(err)
		return err
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
	stmt := fmt.Sprintf("USE %s; ALTER EDGE %s COMMENT = '%s';", workspaceModel.WorkSpaceName, triple.Relationship, tripleStr)
	glog.Info("修改边类型：", stmt)
	_, err = l.svcCtx.Nebula.Execute(stmt)
	if err != nil {
		glog.Error("修改边类型失败：", err)
		return
	}

	return
}
