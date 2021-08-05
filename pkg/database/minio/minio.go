package minio

import (
	"github.com/gabrielopesantos/myDrive-api/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
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