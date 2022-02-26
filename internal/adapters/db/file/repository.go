package file

import (
	"bytes"
	"context"

	"github.com/Lapp-coder/file-service/internal/client"

	"github.com/Lapp-coder/file-service/internal/domain/file"
	"github.com/minio/minio-go/v7"
)

type repository struct {
	minioClient *minio.Client
}

func NewRepository(minioClient *minio.Client) file.Repository {
	return &repository{
		minioClient: minioClient,
	}
}

func (r *repository) SaveFile(file file.File) error {
	var buf bytes.Buffer
	if _, err := buf.Write(file.Content); err != nil {
		return err
	}

	if _, err := r.minioClient.PutObject(
		context.Background(),
		client.MinIOFileBucket,
		file.UUID,
		&buf,
		file.Metadata.Size,
		minio.PutObjectOptions{},
	); err != nil {
		return err
	}

	return nil
}

func (r *repository) GetFileByUUID(uuid string) (file.File, error) {
	obj, err := r.minioClient.GetObject(context.Background(), client.MinIOFileBucket, uuid, minio.GetObjectOptions{})
	if err != nil {
		return file.File{}, err
	}

	var buf bytes.Buffer
	if _, err = buf.ReadFrom(obj); err != nil {
		return file.File{}, err
	}

	objInfo, err := obj.Stat()
	if err != nil {
		return file.File{}, err
	}

	return file.File{
		UUID:    objInfo.Key,
		Content: buf.Bytes(),
	}, nil
}
