package repository

import (
	"github.com/Lapp-coder/file-service/internal/model"
	"github.com/jackc/pgx"
	"github.com/minio/minio-go/v7"
)

type File interface {
	SaveFile(model.File) error
	GetFileByUUID(string) (model.File, error)
	GetFileStatisticByUUID(string) (model.FileStatistic, error)
}

type Repository struct {
	File
}

func New(client *minio.Client, conn *pgx.Conn) Repository {
	return Repository{
		File: NewFileRepository(client, conn),
	}
}
