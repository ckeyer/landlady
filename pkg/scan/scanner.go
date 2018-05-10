package scan

import (
	"net/http"

	pb "github.com/funxdata/landlady/proto"
)

type Metadata interface {
	ModuleName() string
	BatchName() string
}

type PageCounter interface {
	Metadata
	PageCount(*http.Client) (int, error)
}

type URLScanner interface {
	Metadata
	ScanURLs(cli *http.Client, pageIndex int) ([]string, error)
}

type Handler interface {
	Handle(cli *http.Client, req *http.Request) (*pb.House, error)
}
