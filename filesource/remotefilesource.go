package filesource

import (
	"io"
	"net/http"
)

type RemoteFileSource struct {
	url string
}

func Remote(url string) RemoteFileSource {
	return RemoteFileSource{
		url: url,
	}
}

func (remoteFileSource RemoteFileSource) GetContent() (content []byte, err error) {
	resp, err := http.Get(remoteFileSource.url)
	if err != nil {
		return
	}

	return io.ReadAll(resp.Body)
}
