package minio

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/golang/glog"
	minioV7 "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioConf struct {
	AccessKey string
	SecretKey string
	PublicURL string
	Host      string
	Port      int
	Secure    bool
}

type Client struct {
	*minioV7.Client
}

func (m *MinioConf) Validate() (err error) {
	if m.AccessKey == "" {
		err = errors.New("accesskey is empty")
		glog.Error(err)
		return
	}

	if m.SecretKey == "" {
		err = errors.New("accesskey is empty")
		glog.Error(err)
		return
	}

	if m.Host == "" {
		err = errors.New("minio host is empty")
		glog.Error(err)
		return
	}

	if m.Port <= 1024 || m.Port >= 65535 {
		err = errors.New("invalid minio port")
		glog.Error(err)
		return
	}

	return
}

func NewMinioClient(conf MinioConf) (client *Client, err error) {
	c, err := minioV7.New(
		fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		&minioV7.Options{
			Creds:  credentials.NewStaticV4(conf.AccessKey, conf.SecretKey, ""),
			Secure: conf.Secure,
		})

	if err != nil {
		glog.Error(err)
		return
	}

	client = &Client{Client: c}

	return
}

func MakeBucketPolicy(bucket, access string) (policy string) {
	// Reference:
	// 1. https://blog.csdn.net/qq_35991226/article/details/108889535
	// 2. https://docs.min.io/docs/golang-client-api-reference.html#PresignedPostPolicy
	principal := `"Principal":{"AWS":["*"]}`
	if access != "" {
		principal = fmt.Sprintf(`"Principal":{"AWS": "arn:aws:iam::cloudfront:user/CloudFront Origin Access Identity %s"}`, access)
	}

	policy = fmt.Sprintf(`{"Version":"2012-10-17",
		"Statement":[
			{
				"Effect":"Allow",
				%s,
				"Action":["s3:ListBucketMultipartUploads","s3:GetBucketLocation","s3:ListBucket"],
				"Resource":["arn:aws:s3:::%s"]
		   },
		   {
				"Effect":"Allow",
				%s,
				"Action":["s3:ListMultipartUploadParts","s3:PutObject","s3:AbortMultipartUpload","s3:DeleteObject","s3:GetObject"],
				"Resource":["arn:aws:s3:::%s/*"]
		   },
		   {
				"Effect":"Allow",
		   		"Principal":{"AWS":["*"]},
		   		"Action":["s3:GetObject"],
		   		"Resource":["arn:aws:s3:::%s/*"]
		   }]
	   }`,
		principal, bucket, principal, bucket, bucket)
	return
}

func (client *Client) CreateBucketIfNotExisted(bucket string) (err error) {
	location := "cn-north-1"
	exists, err := client.BucketExists(context.Background(), bucket)
	if err != nil {
		glog.Error(err)
		return
	}

	if !exists {
		err = client.MakeBucket(context.Background(), bucket,
			minioV7.MakeBucketOptions{
				Region:        location,
				ObjectLocking: false,
			})
		if err != nil {
			glog.Error(err)
			return
		}

		policy := MakeBucketPolicy(bucket, "")
		err = client.SetBucketPolicy(context.Background(), bucket, policy)
		if err != nil {
			glog.Error(err)
			return
		}

		glog.Infof("Successfully created bucket [%s]", bucket)
	}
	return
}

func (client *Client) CreateAccessBucketIfNotExisted(bucket string, access string) (err error) {
	location := "cn-north-1"
	exists, err := client.BucketExists(context.Background(), bucket)
	if err != nil {
		glog.Error(err)
		return
	}

	if !exists {
		err = client.MakeBucket(context.Background(), bucket,
			minioV7.MakeBucketOptions{
				Region:        location,
				ObjectLocking: false,
			})
		if err != nil {
			glog.Error(err)
			return
		}

		policy := MakeBucketPolicy(bucket, access)
		err = client.SetBucketPolicy(context.Background(), bucket, policy)
		if err != nil {
			glog.Error(err)
			return
		}

		glog.Infof("Successfully created bucket [%s]", bucket)
	}

	return
}

func MinioUpload(client *Client, bucket string,
	fp string, objName string, contentType string) (err error) {

	err = client.CreateBucketIfNotExisted(bucket)
	if err != nil {
		glog.Error(err)
		return
	}

	n, err := client.FPutObject(context.Background(), bucket, objName, fp, minioV7.PutObjectOptions{ContentType: contentType})
	if err != nil {
		glog.Error(err)
		return
	}

	glog.Infof("Successfully uploaded %s of info %v", objName, n)

	return
}

func (client *Client) MinioDownload(bucket string, fp string, objName string) (err error) {
	err = client.FGetObject(context.Background(), bucket, objName, fp, minioV7.GetObjectOptions{})
	if err != nil {
		glog.Error(err)
		return
	}

	return
}

func (client *Client) MinioUploadObject(bucket, objName string,
	reader io.Reader, objSize int64, contentType string) (err error) {

	err = client.CreateBucketIfNotExisted(bucket)
	if err != nil {
		glog.Warningf("%v bucket: %v", err, bucket)
	}

	_, err = client.PutObject(context.Background(),
		bucket, objName, reader, objSize, minioV7.PutObjectOptions{ContentType: contentType})
	if err != nil {
		glog.Error(err)
		return
	}

	return
}
