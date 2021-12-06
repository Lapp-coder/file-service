package repository

import (
	"github.com/jackc/pgx"

	"github.com/Lapp-coder/file-service/internal/config"
)

const (
	fileTable          = "file"
	fileStatisticTable = "file_statistic"
)

func NewPostgresConn(cfg config.Postgres) (*pgx.Conn, error) {
	connCfg := pgx.ConnConfig{
		Host:     cfg.Host,
		Port:     cfg.Port,
		User:     cfg.User,
		Password: cfg.Password,
		Database: cfg.DBName,
	}
	conn, err := pgx.Connect(connCfg)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
