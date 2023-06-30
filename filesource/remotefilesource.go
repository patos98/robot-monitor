package filesource

import (
	"io"
	"net/http"
)

func Remote(url string) FileSourceFn {
	return FileSourceFn(func() (content []byte, err error) {
		resp, err := http.Get(url)
		if err != nil {
			return
		}

		return io.ReadAll(resp.Body)
	})
}
