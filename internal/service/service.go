package service

import (
	"github.com/Lapp-coder/file-service/internal/model"
	"github.com/Lapp-coder/file-service/internal/repository"
)

type File interface {
	SaveFile(model.File) error
	GetFileByUUID(string) (model.File, error)
	GetFileStatisticByUUID(string) (model.FileStatistic, error)
}

type Service struct {
	File
}

func New(repos repository.Repository) Service {
	return Service{
		File: NewFileService(repos.File),
	}
}
