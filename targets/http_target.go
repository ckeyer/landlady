package targets

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"

	pb "github.com/funxdata/landlady/proto"
)

type HTTPTarget struct{}

// DownloadPage
func (ht HTTPTarget) DownloadPage(cli *http.Client, req *http.Request) (*pb.URLInfo, *bytes.Buffer, error) {
	resp, err := cli.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		return nil, nil, err
	}

	ui := &pb.URLInfo{
		OriginURL: req.URL.String(),
		RealURL:   resp.Request.URL.String(),
		ClearURL:  ClearURL(resp.Request.URL),
	}

	return ui, buf, nil
}

func ClearURL(u *url.URL) string {
	return fmt.Sprintf("%s://%s%s", u.Scheme, u.Host, u.EscapedPath())
}
