package file

type Repository interface {
	SaveFile(File) error
	GetFileByUUID(string) (File, error)
}
