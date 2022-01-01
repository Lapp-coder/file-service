package file

import (
	"github.com/sirupsen/logrus"
)

type Service interface {
	SaveFile(File) error
	GetFileByUUID(string) (File, error)
	GetFileStatisticByUUID(string) (Statistic, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) SaveFile(file File) error {
	if err := s.repository.SaveFile(file); err != nil {
		logrus.Error(err)
		return ErrFailedToSaveFile
	}

	return nil
}

func (s *service) GetFileByUUID(uuid string) (File, error) {
	file, err := s.repository.GetFileByUUID(uuid)
	if err != nil {
		logrus.Error(err)
		return File{}, ErrFailedToGetFileByUUID
	}

	return file, nil
}

func (s *service) GetFileStatisticByUUID(uuid string) (Statistic, error) {
	fileStatistic, err := s.repository.GetFileStatisticByUUID(uuid)
	if err != nil {
		logrus.Error(err)
		return Statistic{}, ErrFailedToGetFileStatisticByUUID
	}

	return fileStatistic, nil
}
