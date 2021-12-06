package model

type File struct {
	UUID    string
	Content []byte
	FileMetadata
	FileStatistic
}

type FileMetadata struct {
	Name string
	Size int64
}

type FileStatistic struct {
	RequestCount int `json:"request_count"`
}
