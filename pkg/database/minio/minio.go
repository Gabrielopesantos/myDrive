package minio

import (
	"context"
	"fmt"
	"github.com/gabrielopesantos/myDrive-api/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/pkg/errors"
)

func NewMinioClient(cfg *config.Config) (*minio.Client, error) {
	minioClient, err := minio.New(cfg.Minio.Endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(cfg.Minio.RootUser, cfg.Minio.RootPassword, ""),
		Secure: cfg.Minio.UseSSL})
	if err != nil {
		return nil, err
	}

	return minioClient, nil
}

// CreateBucket checks if a bucket exists and creates it if it does not
func CreateBucket(client *minio.Client, bucketName string) error {

	exists, err := client.BucketExists(context.TODO(), bucketName)
	if err != nil {
		return errors.Wrap(err, "Failed to check if bucket exists")
	}

	if !exists {
		err = client.MakeBucket(context.TODO(), bucketName, minio.MakeBucketOptions{})

		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("Failed to create bucket with name %s", bucketName))
		}
	}

	return nil
}
