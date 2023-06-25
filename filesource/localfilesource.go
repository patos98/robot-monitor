package filesource

import "os"

type LocalFileSource struct {
	path string
}

func Local(path string) LocalFileSource {
	return LocalFileSource{path: path}
}

func (lfs LocalFileSource) GetContent() ([]byte, error) {
	return os.ReadFile(lfs.path)
}
