package session

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"strings"

	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/golang/glog"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPublicKeyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPublicKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPublicKeyLogic {
	return &GetPublicKeyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPublicKeyLogic) GetPublicKey() (resp *types.GetPublickeyResp, err error) {
	x509PublicKey, err := x509.MarshalPKIXPublicKey(&l.svcCtx.PrivateKey.PublicKey)
	if err != nil {
		glog.Error(err)
		return
	}

	block := &pem.Block{
		Bytes: x509PublicKey,
	}

	pemPulicKey := pem.EncodeToMemory(block)

	// Trans []byte to string
	publicKeyStr := strings.Replace(string(pemPulicKey), "\n", "", -1)

	// Remove the useless part
	publicKey := strings.Trim(publicKeyStr, "-----BEGIN -----")
	publicKey = strings.Trim(publicKey, "-----END -----")

	resp = &types.GetPublickeyResp{
		PublicKey: publicKey,
	}

	return
}
