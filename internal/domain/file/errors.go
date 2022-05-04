package file

import "errors"

var (
	ErrFailedToSaveFile               = errors.New("failed to save file")
	ErrFailedToGetFileByUUID          = errors.New("failed to get file by uuid")
)
