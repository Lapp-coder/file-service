package file

import "errors"

var (
	errFailedToOpenFile          = errors.New("failed to open file from request")
	errFailedToReadFileContent   = errors.New("failed to read file content")
	errIncorrectFileUUID         = errors.New("incorrect file uuid")
	errFailedToCreateFileForSend = errors.New("failed to create file for send")
)
