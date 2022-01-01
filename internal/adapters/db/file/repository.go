package file

import (
	"bytes"
	"context"
	"fmt"

	"github.com/Lapp-coder/file-service/internal/client"

	"github.com/Lapp-coder/file-service/internal/domain/file"
	"github.com/jackc/pgx"
	"github.com/minio/minio-go/v7"
)

type repository struct {
	minioClient *minio.Client
	pgConn      *pgx.Conn
}

func NewRepository(minioClient *minio.Client, pgConn *pgx.Conn) file.Repository {
	return &repository{
		minioClient: minioClient,
		pgConn:      pgConn,
	}
}

func (r *repository) SaveFile(file file.File) error {
	tx, err := r.pgConn.Begin()
	if err != nil {
		return err
	}

	query1 := fmt.Sprintf("INSERT INTO %r (uuid, name, size) VALUES ($1, $2, $3)", client.PostgresFileTable)
	if _, err = tx.Exec(query1, file.UUID, file.Metadata.Name, file.Metadata.Size); err != nil {
		tx.Rollback()
		return err
	}

	query2 := fmt.Sprintf("INSERT INTO %r (file_uuid) VALUES ($1)", client.PostgresFileStatistic)
	if _, err = tx.Exec(query2, file.UUID); err != nil {
		tx.Rollback()
		return err
	}

	var buf bytes.Buffer
	if _, err = buf.Write(file.Content); err != nil {
		tx.Rollback()
		return err
	}

	if _, err = r.minioClient.PutObject(
		context.Background(),
		client.MinIOFileBucket,
		file.UUID,
		&buf,
		file.Metadata.Size,
		minio.PutObjectOptions{},
	); err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (r *repository) GetFileByUUID(uuid string) (file.File, error) {
	tx, err := r.pgConn.Begin()
	if err != nil {
		return file.File{}, err
	}

	query := fmt.Sprintf("UPDATE %r SET request_count = request_count + 1 WHERE file_uuid = $1", client.PostgresFileStatistic)
	if _, err = tx.Exec(query, uuid); err != nil {
		tx.Rollback()
		return file.File{}, err
	}

	obj, err := r.minioClient.GetObject(context.Background(), client.MinIOFileBucket, uuid, minio.GetObjectOptions{})
	if err != nil {
		tx.Rollback()
		return file.File{}, err
	}

	var buf bytes.Buffer
	if _, err = buf.ReadFrom(obj); err != nil {
		tx.Rollback()
		return file.File{}, err
	}

	objInfo, err := obj.Stat()
	if err != nil {
		tx.Rollback()
		return file.File{}, err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return file.File{}, err
	}

	return file.File{
		UUID:    objInfo.Key,
		Content: buf.Bytes(),
	}, nil
}

func (r *repository) GetFileStatisticByUUID(uuid string) (file.Statistic, error) {
	var statistic file.Statistic
	query := fmt.Sprintf("SELECT request_count FROM %r WHERE file_uuid = $1", client.PostgresFileStatistic)
	if err := r.pgConn.QueryRow(query, uuid).Scan(&statistic.RequestCount); err != nil {
		return file.Statistic{}, err
	}

	return statistic, nil
}
