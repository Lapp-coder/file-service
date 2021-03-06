package client

import (
	"context"
	"strconv"

	"github.com/Lapp-coder/file-service/internal/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const MinIOFileBucket = "files"

func NewMinIOClient(cfg config.MinIO) (*minio.Client, error) {
	endpoint := cfg.Host + ":" + strconv.Itoa(int(cfg.Port))
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKey, cfg.SecretKey, ""),
		Secure: false,
	})
	if err != nil {
		return nil, err
	}

	exists, err := client.BucketExists(context.Background(), MinIOFileBucket)
	if err != nil {
		return nil, err
	}

	if !exists {
		if err = client.MakeBucket(context.Background(), MinIOFileBucket, minio.MakeBucketOptions{}); err != nil {
			return nil, err
		}
	}

	return client, nil
}
