package zufang58xian

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/ckeyer/logrus"
	pb "github.com/funxdata/landlady/proto"
	"github.com/funxdata/landlady/targets"
)

const (
	moduleName = "58zufang_xian"
	homePage   = "http://xa.58.com/zufang/"
	pagerPre   = "http://xa.58.com/zufang/pn"
)

type Zufang58xian struct {
	targets.HTTPTarget

	logger *logrus.Logger
}

func New() *Zufang58xian {
	logger := logrus.New(logrus.Fields{"module": moduleName})
	logger.SetLevel(logrus.DebugLevel)

	return &Zufang58xian{
		logger: logger,
	}
}

// ModuleName 模块名称
func (z Zufang58xian) ModuleName() string {
	return moduleName
}

// PageCount 统计列表页共有多少页
func (z *Zufang58xian) PageCount(cli *http.Client) (int, error) {
	resp, err := cli.Get(homePage)
	if err != nil {
		return 0, err
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return 0, err
	}

	count := 0
	doc.Find(".main .content .pager a").Each(func(i int, s *goquery.Selection) {
		c := s.Find("span").Text()
		n, _ := strconv.Atoi(c)
		if n > count {
			count = n
		}
	})

	if count == 0 {
		return 0, fmt.Errorf("get 0 page for %s.", moduleName)
	}

	return count, nil
}

func (z *Zufang58xian) ScanURLs(cli *http.Client, pageIndex int) ([]string, error) {
	resp, err := cli.Get(z.pageURL(pageIndex))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	urls := []string{}
	doc.Find(".main .content .listUl li").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Find(".des a").Attr("href")
		if _, err := url.Parse(href); err != nil {
			return
		}
		urls = append(urls, href)
	})

	return urls, nil
}

// ResolveRequest
func (z Zufang58xian) Handle(cli *http.Client, req *http.Request) (*pb.House, error) {
	uinfo, body, err := z.DownloadPage(cli, req)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return nil, err
	}

	_ = doc
	info := &pb.TaskMetadata{
		Url:      uinfo,
		Module:   z.ModuleName(),
		HandleAt: time.Now(),
	}

	return &pb.House{Metadata: info}, nil
}

func (z Zufang58xian) shortURL(u *url.URL) string {
	return fmt.Sprintf("%s://%s%s", u.Scheme, u.Host, u.EscapedPath())
}

func (z Zufang58xian) pageURL(n int) string {
	if n <= 1 {
		return homePage
	}
	return fmt.Sprintf("%s%v/", pagerPre, n)
}
