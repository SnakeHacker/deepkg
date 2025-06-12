package extract_task_result

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/SnakeHacker/deepkg/admin/internal/model/gorm_model"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/golang/glog"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRelationshipLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateRelationshipLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRelationshipLogic {
	return &CreateRelationshipLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateRelationshipLogic) CreateRelationship(req *types.CreateRelationshipReq) (err error) {
	relationship := req.Relationship
	relationshipModel := gorm_model.Relationship{
		SourceEntityID:   int(relationship.SourceEntityID),
		TargetEntityID:   int(relationship.TargetEntityID),
		RelationshipName: relationship.RelationshipName,
		TaskID:           int(relationship.TaskID),
	}

	err = dao.CreateRelationship(l.svcCtx.DB, &relationshipModel)
	if err != nil {
		glog.Error(err)
		return
	}

	return
}
