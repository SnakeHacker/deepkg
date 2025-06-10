package user

import (
	"context"
	"time"

	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/golang/glog"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserListLogic) GetUserList(req *types.GetUserListReq) (resp *types.GetUserListResp, err error) {
	usersModel, total, err := dao.SelectUsers(l.svcCtx.DB, req.PageNumber, req.PageSize)
	if err != nil {
		glog.Error(err)
		return
	}

	users := []types.User{}
	for _, user := range usersModel {
		createdAt, _ := time.Parse(time.RFC3339, user.CreatedAt)
		updatedAt, _ := time.Parse(time.RFC3339, user.UpdatedAt)

		users = append(users, types.User{
			ID:        int64(user.ID),
			UserCode:  user.UserCode,
			OrgID:     int64(user.OrgID),
			OrgName:   user.OrgName,
			Account:   user.Account,
			Username:  user.Username,
			Phone:     user.Phone,
			Mail:      user.Mail,
			Enable:    user.Enable,
			Role:      user.Role,
			Avatar:    user.Avatar,
			CreatedAt: createdAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: updatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	resp = &types.GetUserListResp{
		PageSize:   req.PageSize,
		PageNumber: req.PageNumber,
		Users:      users,
		Total:      total,
	}

	return
}
