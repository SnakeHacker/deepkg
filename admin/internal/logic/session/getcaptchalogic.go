package session

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/common/captcha"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/golang/glog"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCaptchaLogic {
	return &GetCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCaptchaLogic) GetCaptcha() (resp *types.GetCaptchaResp, err error) {
	captchaID, captchaImage, err := captcha.GenerateCaptcha(l.svcCtx.Captcha)
	if err != nil {
		glog.Error(err)
		return
	}

	resp = &types.GetCaptchaResp{
		CaptchaID:     captchaID,
		CaptchaBase64: captchaImage,
	}

	return
}
