package composites

import (
	"github.com/Lapp-coder/file-service/internal/adapters/api"
	file3 "github.com/Lapp-coder/file-service/internal/adapters/api/v1/file"
	file2 "github.com/Lapp-coder/file-service/internal/adapters/db/file"
	"github.com/Lapp-coder/file-service/internal/domain/file"
	"github.com/jackc/pgx"
	"github.com/minio/minio-go/v7"
)

type FileComposite struct {
	Handler    api.Handler
	Service    file.Service
	Repository file.Repository
}

func NewFileComposite(minioClient *minio.Client, pgConn *pgx.Conn) *FileComposite {
	repository := file2.NewRepository(minioClient, pgConn)
	service := file.NewService(repository)
	handler := file3.NewHandler(service)

	return &FileComposite{
		Handler:    handler,
		Service:    service,
		Repository: repository,
	}
}
