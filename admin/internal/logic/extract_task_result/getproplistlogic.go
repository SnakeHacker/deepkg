package extract_task_result

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPropListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPropListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPropListLogic {
	return &GetPropListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPropListLogic) GetPropList(req *types.GetPropListReq) (resp *types.GetPropListResp, err error) {
	propModels, total, err := dao.SelectProps(l.svcCtx.DB, int(req.EntityID), req.PageNumber, req.PageSize)

	props := []types.EntityProp{}
	for _, pm := range propModels {
		prop := types.EntityProp{
			ID:        int64(pm.ID),
			PropName:  pm.PropName,
			PropValue: pm.PropValue,
		}
		props = append(props, prop)
	}

	resp = &types.GetPropListResp{
		Total:      total,
		Props:      props,
		PageSize:   req.PageSize,
		PageNumber: req.PageNumber,
	}

	return
}
