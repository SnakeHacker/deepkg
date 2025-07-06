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

type DeleteSchemaTriplesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSchemaTriplesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSchemaTriplesLogic {
	return &DeleteSchemaTriplesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSchemaTriplesLogic) DeleteSchemaTriples(req *types.DeleteSchemaTriplesReq) (err error) {
	tripleModels, err := dao.SelectSchemaTriplesByIDs(l.svcCtx.DB, req.IDs)
	if err != nil {
		glog.Error(err)
		return
	}

	for _, tripleModel := range tripleModels {
		workspaceModel, err := dao.SelectKnowledgeGraphWorkspaceByID(l.svcCtx.DB, int64(tripleModel.WorkSpaceID))
		if err != nil {
			glog.Error("查询工作空间失败：", err)
			return err
		}

		stmt := fmt.Sprintf("USE %s; DROP EDGE IF EXISTS %s", workspaceModel.WorkSpaceName, tripleModel.Relationship)
		glog.Info("删除边类型：", stmt)
		_, err = l.svcCtx.Nebula.Execute(stmt)
		if err != nil {
			glog.Error("删除边类型失败：", err)
			return err
		}
	}

	err = dao.DeleteSchemaTriplesByIDs(l.svcCtx.DB, req.IDs)
	if err != nil {
		glog.Error(err)
		return
	}
	return
}
