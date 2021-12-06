package v1

import (
	"bytes"
	"fmt"
	"os"

	"github.com/Lapp-coder/file-service/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (h Handler) uploadFile(ctx *fiber.Ctx) error {
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		return respondJSON(ctx, fiber.StatusBadRequest, err)
	}

	f, err := fileHeader.Open()
	if err != nil {
		logrus.Error(err)
		return respondJSON(ctx, fiber.StatusInternalServerError, errFailedToOpenFile)
	}

	var buf bytes.Buffer
	if _, err = buf.ReadFrom(f); err != nil {
		logrus.Error(err)
		return respondJSON(ctx, fiber.StatusInternalServerError, errFailedToReadFileContent)
	}

	file := model.File{
		UUID:    uuid.NewString(),
		Content: buf.Bytes(),
		FileMetadata: model.FileMetadata{
			Name: fileHeader.Filename,
			Size: fileHeader.Size,
		},
	}

	if err = h.service.File.SaveFile(file); err != nil {
		return respondJSON(ctx, fiber.StatusInternalServerError, err)
	}

	url := fmt.Sprintf("%s/api/v1/files/%s", ctx.BaseURL(), file.UUID)
	return respondJSON(ctx, fiber.StatusCreated, fiber.Map{"url": url})
}

func (h Handler) getFileByUUID(ctx *fiber.Ctx) error {
	fileUUID := ctx.Params("uuid")
	if _, err := uuid.Parse(fileUUID); err != nil {
		logrus.Error(err)
		return respondJSON(ctx, fiber.StatusBadRequest, errIncorrectFileUUID)
	}

	file, err := h.service.File.GetFileByUUID(fileUUID)
	if err != nil {
		return respondJSON(ctx, fiber.StatusBadRequest, err)
	}

	filePath := "/tmp/" + file.UUID
	tmpFile, err := os.Create(filePath)
	if err != nil {
		logrus.Error(err)
		return respondJSON(ctx, fiber.StatusInternalServerError, errFailedToCreateFileForSend)
	}
	defer os.Remove(filePath)

	if _, err = tmpFile.Write(file.Content); err != nil {
		logrus.Error(err)
		return respondJSON(ctx, fiber.StatusInternalServerError, errFailedToCreateFileForSend)
	}

	if err = tmpFile.Close(); err != nil {
		logrus.Error(err)
		return respondJSON(ctx, fiber.StatusInternalServerError, errFailedToCreateFileForSend)
	}

	return ctx.SendFile(filePath)
}

func (h Handler) getFileStatisticByUUID(ctx *fiber.Ctx) error {
	fileUUID := ctx.Params("uuid")
	if _, err := uuid.Parse(fileUUID); err != nil {
		logrus.Error(err)
		return respondJSON(ctx, fiber.StatusBadRequest, errIncorrectFileUUID)
	}

	fileStatistic, err := h.service.File.GetFileStatisticByUUID(fileUUID)
	if err != nil {
		return respondJSON(ctx, fiber.StatusInternalServerError, err)
	}

	return respondJSON(ctx, fiber.StatusOK, fiber.Map{"statistics": fileStatistic})
}
