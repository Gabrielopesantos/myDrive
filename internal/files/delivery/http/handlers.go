package http

import (
	"bytes"
	"io"
    "os"
	"net/http"

	"github.com/gabrielopesantos/myDrive-api/config"
	"github.com/gabrielopesantos/myDrive-api/internal/files"
	"github.com/gabrielopesantos/myDrive-api/internal/models"
	httpErrors "github.com/gabrielopesantos/myDrive-api/pkg/http_errors"
	"github.com/gabrielopesantos/myDrive-api/pkg/logger"
	"github.com/gabrielopesantos/myDrive-api/pkg/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/minio/minio-go/v7"
	"github.com/opentracing/opentracing-go"
)

// fileHandlers
type fileHandlers struct {
	cfg         *config.Config
	fileService files.Service
	logger      logger.Logger
}

func NewFileHandlers(cfg *config.Config, fileService files.Service, logger logger.Logger) files.Handlers {
	return &fileHandlers{
		cfg:         cfg,
		fileService: fileService,
		logger:      logger,
	}
}

func (h *fileHandlers) GetUserFiles() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, _ := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "fileHandlers.GetFileById")
		defer span.Finish()

		return c.JSON(http.StatusOK, "")
	}
}

func (h *fileHandlers) GetFileById() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "fileHandlers.GetFileById")
		defer span.Finish()

		fileId := c.Param("file_id")

		parsedFileId, err := uuid.Parse(fileId)
		if err != nil {
			// Invalid file id
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		file, err := h.fileService.GetFileById(ctx, parsedFileId)

		object, err := h.fileService.RetrieveObjectFromBucket(ctx, file)
		if err != nil {
			return c.JSON(httpErrors.ErrorResponse(err))
		}

        filePath, err := saveObjectTmp(file, object)
		if err != nil {
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.File(filePath)
	}
}

func (h *fileHandlers) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "fileHandlers.Insert")
		defer span.Finish()

		bucket := "files"
		userID := c.Get("uid").(uuid.UUID)
		fileDescription := c.FormValue("description")

		file, err := utils.ReadFile(c, "file")
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			c.JSON(httpErrors.ErrorResponse(err))
		}

		fileHeaders, err := file.Open()
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			c.JSON(httpErrors.ErrorResponse(err))
		}
		defer fileHeaders.Close()

		binaryFile := bytes.NewBuffer(nil)
		if _, err = io.Copy(binaryFile, fileHeaders); err != nil {
			utils.LogResponseError(c, h.logger, err)
			c.JSON(httpErrors.ErrorResponse(err))
		}

		reader := bytes.NewReader(binaryFile.Bytes())

		contentType := http.DetectContentType(binaryFile.Bytes())

		fileID := uuid.New()

		uploadedFile := &models.File{
			FileId:      fileID,
			FileOwnerId: userID,
			UploadInput: &models.UploadInput{
				File:        reader,
				Name:        file.Filename,
				Size:        file.Size,
				ContentType: contentType,
				BucketName:  bucket,
			},
			Description: fileDescription,
		}

		uploadInfo, err := h.fileService.Insert(ctx, uploadedFile)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, uploadInfo)
	}
}

func saveObjectTmp(file *models.File, object *minio.Object) (string, error) {

    localFile, err := os.Create("/tmp/" + file.Name)
    if err != nil {
        return "", err
    }
    if _, err = io.Copy(localFile, object); err != nil {
        return "", err
    }

    return "/tmp/" + file.Name, nil
}
