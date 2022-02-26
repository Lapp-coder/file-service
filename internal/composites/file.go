package composites

import (
	"github.com/Lapp-coder/file-service/internal/adapters/api"
	fileHandler "github.com/Lapp-coder/file-service/internal/adapters/api/v1/file"
	fileRepository "github.com/Lapp-coder/file-service/internal/adapters/db/file"
	fileService "github.com/Lapp-coder/file-service/internal/domain/file"
	"github.com/minio/minio-go/v7"
)

type FileComposite struct {
	Handler    api.Handler
	Service    fileService.Service
	Repository fileService.Repository
}

func NewFileComposite(minioClient *minio.Client) *FileComposite {
	repository := fileRepository.NewRepository(minioClient)
	service := fileService.NewService(repository)
	handler := fileHandler.NewHandler(service)

	return &FileComposite{
		Handler:    handler,
		Service:    service,
		Repository: repository,
	}
}
