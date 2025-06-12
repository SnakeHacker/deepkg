package extract_task_result

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/golang/glog"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRelationshipLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRelationshipLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRelationshipLogic {
	return &UpdateRelationshipLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRelationshipLogic) UpdateRelationship(req *types.UpdateRelationshipReq) (err error) {
	relationship := req.Relationship
	relationshipModel, err := dao.SelectRelationshipByID(l.svcCtx.DB, int(relationship.ID))
	if err != nil {
		glog.Error(err)
		return
	}

	relationshipModel.RelationshipName = relationship.RelationshipName
	relationshipModel.SourceEntityID = int(relationship.SourceEntityID)
	relationshipModel.TargetEntityID = int(relationship.TargetEntityID)
	relationshipModel.TaskID = int(relationship.TaskID)

	err = dao.UpdateRelationship(l.svcCtx.DB, &relationshipModel)
	if err != nil {
		glog.Error(err)
		return
	}

	return
}
