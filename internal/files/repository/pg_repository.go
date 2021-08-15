package repository

import (
	"context"
	"github.com/gabrielopesantos/myDrive-api/internal/files"
	"github.com/gabrielopesantos/myDrive-api/internal/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"strings"
)

type fileRepo struct {
	db *sqlx.DB
}

func NewFileRepository(db *sqlx.DB) files.Repository {
	return &fileRepo{
		db: db,
	}
}

func (r *fileRepo) CheckFileExistence(ctx context.Context, fileID uuid.UUID) (*models.File, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "fileRepo.RecordFileInsertion")
	defer span.Finish()

	f := &models.File{}
	if err := r.db.QueryRowxContext(ctx, getFileByIdQuery, fileID).StructScan(f); err != nil {
		return nil, errors.Wrap(err, "userRepo.CheckFileExistence.StructScan")
	}

	return f, nil
}

func (r *fileRepo) RecordFileInsertion(ctx context.Context, file *models.File) (*models.File, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "fileRepo.RecordFileInsertion")
	defer span.Finish()

	f := &models.File{}
	if err := r.db.QueryRowxContext(ctx, createFileInsertion, file.FileId, file.FileOwnerId, file.BucketURL,
		file.Name, file.ContentType, file.Size, file.Description, r.tagsToString(file.Tags)).StructScan(f); err != nil {
		return nil, errors.Wrap(err, "fileRepo.RecordFileInsertion.StructScan")
	}

	return f, nil
}

func (r *fileRepo) tagsToString(tags []string) string {
	return strings.Join(tags, ",")
}
