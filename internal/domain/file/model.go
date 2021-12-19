package file

type File struct {
	UUID      string
	Content   []byte
	Metadata  Metadata
	Statistic Statistic
}

type Metadata struct {
	Name string
	Size int64
}

type Statistic struct {
	RequestCount int `json:"request_count"`
}
