package file

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type Service interface {
	SaveFile(File) error
	GetFileByUUID(uuid.UUID) (File, error)
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

func (s *service) GetFileByUUID(uuid uuid.UUID) (File, error) {
	file, err := s.repository.GetFileByUUID(uuid.String())
	if err != nil {
		logrus.Error(err)
		return File{}, ErrFailedToGetFileByUUID
	}

	return file, nil
}
