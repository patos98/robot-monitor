package filesource

import "os"

type FileSourceFn func() ([]byte, error)

func Local(path string) FileSourceFn {
	return FileSourceFn(func() ([]byte, error) {
		return os.ReadFile(path)
	})
}

func (fileSourceFn FileSourceFn) GetContent() ([]byte, error) {
	return fileSourceFn()
}
