package filesource

import "os"

func Local(path string) FileSourceFn {
	return FileSourceFn(func() ([]byte, error) {
		return os.ReadFile(path)
	})
}
