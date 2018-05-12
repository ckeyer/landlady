package targets

import (
	"bytes"
	"io"
	"net/http"
)

type HTTPTarget struct{}

// DownloadPage
func (ht HTTPTarget) DownloadPage(cli *http.Client, req *http.Request) (*bytes.Buffer, error) {
	resp, err := cli.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		return nil, err
	}

	return buf, nil
}
