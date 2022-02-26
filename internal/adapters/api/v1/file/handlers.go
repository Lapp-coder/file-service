package file

import (
	"bytes"
	"fmt"
	"os"

	"github.com/Lapp-coder/file-service/internal/adapters/api"

	"github.com/Lapp-coder/file-service/internal/domain/file"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const (
	filesURI = "/files"
	fileURI  = "/files/:uuid"
)

type handler struct {
	service file.Service
}

func NewHandler(service file.Service) api.Handler {
	return &handler{service: service}
}

func (h *handler) Init(router fiber.Router) {
	router.Post(filesURI, h.uploadFile)
	router.Get(fileURI, h.getFileByUUID)
}

func (h *handler) uploadFile(ctx *fiber.Ctx) error {
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		return api.RespondJSON(ctx, fiber.StatusBadRequest, err)
	}

	f, err := fileHeader.Open()
	if err != nil {
		logrus.Error(err)
		return api.RespondJSON(ctx, fiber.StatusInternalServerError, errFailedToOpenFile)
	}

	var buf bytes.Buffer
	if _, err = buf.ReadFrom(f); err != nil {
		logrus.Error(err)
		return api.RespondJSON(ctx, fiber.StatusInternalServerError, errFailedToReadFileContent)
	}

	file := file.File{
		UUID:    uuid.NewString(),
		Content: buf.Bytes(),
		Metadata: file.Metadata{
			Name: fileHeader.Filename,
			Size: fileHeader.Size,
		},
	}

	if err = h.service.SaveFile(file); err != nil {
		return api.RespondJSON(ctx, fiber.StatusInternalServerError, err)
	}

	url := fmt.Sprintf("%s/api/v1/files/%s", ctx.BaseURL(), file.UUID)
	return api.RespondJSON(ctx, fiber.StatusCreated, fiber.Map{"url": url})
}

func (h *handler) getFileByUUID(ctx *fiber.Ctx) error {
	fileUUID, err := uuid.Parse(ctx.Params("uuid"))
	if err != nil {
		logrus.Error(err)
		return api.RespondJSON(ctx, fiber.StatusBadRequest, errIncorrectFileUUID)
	}

	f, err := h.service.GetFileByUUID(fileUUID)
	if err != nil {
		return api.RespondJSON(ctx, fiber.StatusBadRequest, err)
	}

	filePath := "/tmp/" + f.UUID
	tmpFile, err := os.Create(filePath)
	if err != nil {
		logrus.Error(err)
		return api.RespondJSON(ctx, fiber.StatusInternalServerError, errFailedToCreateFileForSend)
	}
	defer os.Remove(filePath)

	if _, err = tmpFile.Write(f.Content); err != nil {
		logrus.Error(err)
		return api.RespondJSON(ctx, fiber.StatusInternalServerError, errFailedToCreateFileForSend)
	}

	if err = tmpFile.Close(); err != nil {
		logrus.Error(err)
		return api.RespondJSON(ctx, fiber.StatusInternalServerError, errFailedToCreateFileForSend)
	}

	return ctx.SendFile(filePath)
}
