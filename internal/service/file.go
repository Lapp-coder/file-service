package service

import (
	"github.com/Lapp-coder/file-service/internal/model"
	"github.com/Lapp-coder/file-service/internal/repository"
	"github.com/sirupsen/logrus"
)

type FileService struct {
	repos repository.File
}

func NewFileService(repos repository.File) FileService {
	return FileService{repos: repos}
}

func (s FileService) SaveFile(file model.File) error {
	if err := s.repos.SaveFile(file); err != nil {
		logrus.Error(err)
		return ErrFailedToSaveFile
	}

	return nil
}

func (s FileService) GetFileByUUID(uuid string) (model.File, error) {
	file, err := s.repos.GetFileByUUID(uuid)
	if err != nil {
		logrus.Error(err)
		return model.File{}, ErrFailedToGetFileByUUID
	}

	return file, nil
}

func (s FileService) GetFileStatisticByUUID(uuid string) (model.FileStatistic, error) {
	fileStatistic, err := s.repos.GetFileStatisticByUUID(uuid)
	if err != nil {
		logrus.Error(err)
		return model.FileStatistic{}, ErrFailedToGetFileStatisticByUUID
	}

	return fileStatistic, nil
}
