package extract_task

import (
	"context"
	"errors"

	"github.com/SnakeHacker/deepkg/admin/common"
	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/SnakeHacker/deepkg/admin/internal/model/gorm_model"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/golang/glog"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetExtractTaskListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetExtractTaskListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetExtractTaskListLogic {
	return &GetExtractTaskListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetExtractTaskListLogic) GetExtractTaskList(req *types.GetExtractTaskListReq) (resp *types.GetExtractTaskListResp, err error) {
	taskModels, total, err := dao.SelectExtractTasks(l.svcCtx.DB, int(req.WorkSpaceID), req.PageNumber, req.PageSize)
	if err != nil {
		glog.Error(err)
		return
	}

	creatorIDs := []int64{}
	for _, taskModel := range taskModels {
		creatorIDs = append(creatorIDs, int64(taskModel.CreatorID))
	}

	userModels, err := dao.SelectUserModelsByIDs(l.svcCtx.DB, creatorIDs)
	if err != nil {
		glog.Error(err)
		return
	}
	userMap := map[int64]gorm_model.User{}
	for _, userModel := range userModels {
		userMap[int64(userModel.ID)] = *userModel
	}

	tasks := []types.ExtractTask{}
	for _, taskModel := range taskModels {

		creator, ok := userMap[int64(taskModel.CreatorID)]
		if !ok {
			err = errors.New("user not found")
			glog.Error(err)
			return
		}

		tasks = append(tasks, types.ExtractTask{
			ID:          int64(taskModel.ID),
			TaskName:    taskModel.TaskName,
			Remark:      taskModel.Remark,
			TaskStatus:  taskModel.TaskStatus,
			WorkSpaceID: int64(taskModel.WorkSpaceID),
			Published:   taskModel.Published,

			CreatorID:   int64(creator.ID),
			CreatorName: creator.Username,
			CreatedAt:   taskModel.CreatedAt.Format(common.TIME_FORMAT),
		})
	}

	resp = &types.GetExtractTaskListResp{
		Total:        total,
		PageSize:     req.PageSize,
		PageNumber:   req.PageNumber,
		ExtractTasks: tasks,
	}

	return
}
