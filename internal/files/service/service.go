package service

import (
	"context"
	"fmt"
	"github.com/gabrielopesantos/myDrive-api/config"
	"github.com/gabrielopesantos/myDrive-api/internal/files"
	"github.com/gabrielopesantos/myDrive-api/internal/models"
	httpErrors "github.com/gabrielopesantos/myDrive-api/pkg/http_errors"
	"github.com/gabrielopesantos/myDrive-api/pkg/logger"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

type fileService struct {
	cfg           *config.Config
	fileRepo      files.Repository
	fileMinioRepo files.MinioRepository
	logger        logger.Logger
}

func NewFileService(cfg *config.Config, fileRepo files.Repository, fileMinioRepo files.MinioRepository, logger logger.Logger) files.Service {
	return &fileService{
		cfg:           cfg,
		fileRepo:      fileRepo,
		fileMinioRepo: fileMinioRepo,
		logger:        logger,
	}
}

func (s *fileService) Insert(ctx context.Context, file *models.File) (*models.File, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "userService.GetUsers")
	defer span.Finish()

	uploadInfo, err := s.fileMinioRepo.PutObject(ctx, *file)
	if err != nil {
		return nil, httpErrors.NewInternalServerError(errors.Wrap(err, "fileService.Insert.PutObject"))
	}

	file.BucketURL = s.generateMinioURL(file.BucketName, uploadInfo.Key)

	fileInsertion, err := s.fileRepo.RecordFileInsertion(ctx, file)
	if err != nil {
		return nil, httpErrors.NewInternalServerError(errors.Wrap(err, "fileService.Insert.RecordFileInsertion"))
	}

	//fileInsertion.File = nil // Does this clean?

	return fileInsertion, nil
}

func (s *fileService) generateMinioURL(bucket, key string) string {
	return fmt.Sprintf("%s/minio/%s/%s", s.cfg.Minio.Endpoint, bucket, key)
}
