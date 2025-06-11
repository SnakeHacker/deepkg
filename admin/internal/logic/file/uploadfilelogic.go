package file

import (
	"bytes"
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"path"

	"github.com/SnakeHacker/deepkg/admin/common"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/golang/glog"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadFileLogic {
	return &UploadFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadFileLogic) UploadFile(req *types.UploadFileReq, r *http.Request) (resp *types.UploadFileResp, err error) {
	err = r.ParseMultipartForm(common.MULTIPART_MEMORY)
	if err != nil {
		glog.Error(err)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		glog.Error(err)
		return
	}
	defer file.Close()

	var buf bytes.Buffer
	tee := io.TeeReader(file, &buf)

	glog.Infof("receivce file: %v, %v", header.Filename, header.Header.Get("Content-Type"))

	byteContainer, err := io.ReadAll(tee)
	if err != nil {
		glog.Error(err)
		return
	}

	objID := fmt.Sprintf("%x%s", md5.Sum(byteContainer), path.Ext(header.Filename))

	bucket := common.FILE_BUCKET

	err = l.svcCtx.Minio.CreateBucketIfNotExisted(bucket)
	if err != nil {
		glog.Error(err)
		return
	}

	err = l.svcCtx.Minio.MinioUploadObject(bucket, objID, bytes.NewReader(byteContainer), header.Size, header.Header.Get("Content-Type"))
	if err != nil {
		glog.Error(err)
		return
	}

	return &types.UploadFileResp{
		Host:     fmt.Sprintf("%s/%s", l.svcCtx.Config.Minio.PublicURL, bucket),
		ObjectID: objID,
	}, nil
}
