package document

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

type GetDocumentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDocumentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDocumentListLogic {
	return &GetDocumentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDocumentListLogic) GetDocumentList(req *types.GetDocumentListReq) (resp *types.GetDocumentListResp, err error) {
	docModels, total, err := dao.SelectDocuments(l.svcCtx.DB, req.DirID, req.PageNumber, req.PageSize)
	if err != nil {
		glog.Error(err)
		return
	}

	creatorIDs := []int64{}
	for _, docModel := range docModels {
		creatorIDs = append(creatorIDs, int64(docModel.CreatorID))
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

	docs := []types.Document{}
	for _, docModel := range docModels {

		creator, ok := userMap[int64(docModel.CreatorID)]
		if !ok {
			err = errors.New("user not found")
			glog.Error(err)
			return
		}

		docs = append(docs, types.Document{
			ID:      int64(docModel.ID),
			DocName: docModel.DocName,
			DocDesc: docModel.DocDesc,
			DocPath: docModel.DocPath,
			DirID:   int64(docModel.DirID),

			CreatorID:   int64(creator.ID),
			CreatorName: creator.Username,
			CreatedAt:   docModel.CreatedAt.Format(common.TIME_FORMAT),
		})
	}

	resp = &types.GetDocumentListResp{
		Documents:  docs,
		PageNumber: req.PageNumber,
		PageSize:   req.PageSize,
		Total:      total,
	}

	return
}
