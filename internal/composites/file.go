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
	Storage file.Storage
	Service file.Service
	Handler api.Handler
}

func NewFileComposite(minioClient *minio.Client, pgConn *pgx.Conn) *FileComposite {
	storage := file2.NewStorage(minioClient, pgConn)
	service := file.NewService(storage)
	handler := file3.NewHandler(service)

	return &FileComposite{
		Storage: storage,
		Service: service,
		Handler: handler,
	}
}
