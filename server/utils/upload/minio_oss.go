package upload

import (
	"context"
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/zap"
	"mime/multipart"
)

type MinioOss struct {}

func (*MinioOss) UploadFile(file *multipart.FileHeader) (string, string, error) {
	f, openError := file.Open()
	if openError != nil {
		global.GVA_LOG.Error("function file.Open() Failed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() Failed, err:" + openError.Error())
	}

	defer f.Close() // 创建文件 defer 关闭

	ctx := context.Background()
	client := NewMinioOss()
	bucketName := global.GVA_CONFIG.MinioOSS.BucketName
	location := "us-east-1"
	err := client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := client.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			global.GVA_LOG.Info("We already own", zap.Any("name", bucketName))
		} else {
			global.GVA_LOG.Error("create error", zap.Any("err", err))
		}
	} else {
		global.GVA_LOG.Info("Successfully created", zap.Any("name", bucketName))
	}

	client.SetBucketPolicy(ctx, bucketName, "{\"Version\":\"2012-10-17\",\"Statement\":[{\"Effect\":\"Allow\",\"Principal\":{\"AWS\":[\"*\"]},\"Action\":[\"s3:ListBucket\",\"s3:ListBucketMultipartUploads\",\"s3:GetBucketLocation\"],\"Resource\":[\"arn:aws:s3:::car\"]},{\"Effect\":\"Allow\",\"Principal\":{\"AWS\":[\"*\"]},\"Action\":[\"s3:DeleteObject\",\"s3:GetObject\",\"s3:ListMultipartUploadParts\",\"s3:PutObject\",\"s3:AbortMultipartUpload\"],\"Resource\":[\"arn:aws:s3:::car/*\"]}]}")
	objectName := file.Filename
	// Upload the zip file with FPutObject
	uploadInfo, err := client.PutObject(context.Background(), bucketName, objectName, f, file.Size, minio.PutObjectOptions{ContentType:"application/octet-stream"})
	if err != nil {
		global.GVA_LOG.Error("upload error", zap.Any("err", err))
		return "", "", err
	}
	global.GVA_LOG.Info("Successfully Uploaded", zap.Any("info", uploadInfo))

	return "http://"+global.GVA_CONFIG.MinioOSS.Endpoint+"/"+ bucketName + "/"+objectName, objectName, nil
}

func (*MinioOss) DeleteFile(key string) error {
	ctx := context.Background()
	client := NewMinioOss()

	opts := minio.RemoveObjectOptions {
		GovernanceBypass: true,
	}
	err := client.RemoveObject(ctx, global.GVA_CONFIG.MinioOSS.BucketName, key, opts)
	if err != nil {
		return err
	}
	return nil
}

func NewMinioOss () *minio.Client{
	endpoint := global.GVA_CONFIG.MinioOSS.Endpoint
	accessKeyID := global.GVA_CONFIG.MinioOSS.AccessKeyId
	secretAccessKey := global.GVA_CONFIG.MinioOSS.AccessKeySecret
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil
	}
	return minioClient
}