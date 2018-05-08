package scan

import "net/http"

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
