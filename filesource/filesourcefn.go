package filesource

type FileSourceFn func() ([]byte, error)

func (fileSourceFn FileSourceFn) GetContent() ([]byte, error) {
	return fileSourceFn()
}
