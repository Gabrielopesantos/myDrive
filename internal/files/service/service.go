package service

import (
	"context"
	"fmt"
	"github.com/gabrielopesantos/myDrive-api/config"
	"github.com/gabrielopesantos/myDrive-api/internal/files"
	"github.com/gabrielopesantos/myDrive-api/internal/models"
	httpErrors "github.com/gabrielopesantos/myDrive-api/pkg/http_errors"
	"github.com/gabrielopesantos/myDrive-api/pkg/logger"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
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

func (s *fileService) GetFileById(ctx context.Context, fileID uuid.UUID) (*models.File, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "userService.GetUsers")
	defer span.Finish()

	// Verify is there is a file with that ID (Maybe first check redis)

	// Else check db and save in redis
	fileData, err := s.fileRepo.CheckFileExistence(ctx, fileID)
	if err != nil {
		return nil, errors.Wrap(err, "Check")
	}

	fileData.UploadInput.File, err = s.fileMinioRepo.GetObject(ctx, fileData.BucketURL)
	if err != nil {
		return nil, errors.Wrap(err, "Check")
	}

	// Add to redis

	return fileData, nil
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

func (s *fileService) RetrieveObjectFromBucket(ctx context.Context, file *models.File) (*minio.Object, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "userService.GetUsers")
	defer span.Finish()

    object, err := s.fileMinioRepo.GetObject(ctx, file.BucketURL)
    if err != nil {
        return nil, httpErrors.NewInternalServerError(errors.Wrap(err, "fileSErvice.RetrieveObjectFromBucket.GetObject"))
    }

    return object, nil
}

func (s *fileService) generateMinioURL(bucket, key string) string {
	return fmt.Sprintf("%s/minio/%s/%s", s.cfg.Minio.Endpoint, bucket, key)
}
