package repository

import (
	"bytes"
	"context"
	"fmt"

	"github.com/jackc/pgx"

	"github.com/Lapp-coder/file-service/internal/model"
	"github.com/minio/minio-go/v7"
)

type FileRepository struct {
	minioClient *minio.Client
	pgConn      *pgx.Conn
}

func NewFileRepository(client *minio.Client, conn *pgx.Conn) FileRepository {
	return FileRepository{minioClient: client, pgConn: conn}
}

func (r FileRepository) SaveFile(file model.File) error {
	tx, err := r.pgConn.Begin()
	if err != nil {
		return err
	}

	query1 := fmt.Sprintf("INSERT INTO %s (uuid, name, size) VALUES ($1, $2, $3)", fileTable)
	if _, err = tx.Exec(query1, file.UUID, file.Name, file.Size); err != nil {
		tx.Rollback()
		return err
	}

	query2 := fmt.Sprintf("INSERT INTO %s (file_uuid) VALUES ($1)", fileStatisticTable)
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
		fileBucket,
		file.UUID,
		&buf,
		file.Size,
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

func (r FileRepository) GetFileByUUID(uuid string) (model.File, error) {
	tx, err := r.pgConn.Begin()
	if err != nil {
		return model.File{}, err
	}

	query := fmt.Sprintf("UPDATE %s SET request_count = request_count + 1 WHERE file_uuid = $1", fileStatisticTable)
	if _, err = tx.Exec(query, uuid); err != nil {
		tx.Rollback()
		return model.File{}, err
	}

	obj, err := r.minioClient.GetObject(context.Background(), fileBucket, uuid, minio.GetObjectOptions{})
	if err != nil {
		tx.Rollback()
		return model.File{}, err
	}

	var buf bytes.Buffer
	if _, err = buf.ReadFrom(obj); err != nil {
		tx.Rollback()
		return model.File{}, err
	}

	objInfo, err := obj.Stat()
	if err != nil {
		tx.Rollback()
		return model.File{}, err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return model.File{}, err
	}

	return model.File{
		UUID:    objInfo.Key,
		Content: buf.Bytes(),
	}, nil
}

func (r FileRepository) GetFileStatisticByUUID(fileUUID string) (model.FileStatistic, error) {
	var fileStatistic model.FileStatistic
	query := fmt.Sprintf("SELECT request_count FROM %s WHERE file_uuid = $1", fileStatisticTable)
	if err := r.pgConn.QueryRow(query, fileUUID).Scan(&fileStatistic.RequestCount); err != nil {
		return model.FileStatistic{}, err
	}

	return fileStatistic, nil
}
