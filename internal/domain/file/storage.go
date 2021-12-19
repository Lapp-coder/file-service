package file

type Storage interface {
	SaveFile(File) error
	GetFileByUUID(string) (File, error)
	GetFileStatisticByUUID(string) (Statistic, error)
}
