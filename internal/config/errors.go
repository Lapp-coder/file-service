package config

import "errors"

var (
	ErrMinIOAccessKeyIsEmpty   = errors.New("minio access key is empty")
	ErrMinIOSecretKeyIsEmpty   = errors.New("minio secret key is empty")
	ErrPostgresPasswordIsEmpty = errors.New("postgres password is empty")
)
