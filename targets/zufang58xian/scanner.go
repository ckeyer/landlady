package zufang58xian

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	moduleName = "58zufang_xian"
	homePage   = "http://xa.58.com/zufang/"
	pagerPre   = "http://xa.58.com/zufang/pn"
)

type Zufang58xian struct {
	batch string
}

func newZufang58xian() *Zufang58xian {
	return &Zufang58xian{
		batch: time.Now().Format("20060102T150405"),
	}
}

// ModuleName 模块名称
func (z Zufang58xian) ModuleName() string {
	return moduleName
}

// BatchName 批次
func (z Zufang58xian) BatchName() string {
	return z.batch
}

// PageCount 统计列表页共有多少页
func (z *Zufang58xian) PageCount(cli *http.Client) (int, error) {
	resp, err := cli.Get(homePage)
	if err != nil {
		return 0, err
	}

	doc, err := goquery.NewDocumentFromResponse(resp)
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
	resp, err := cli.Get(pageURL(1))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromResponse(resp)
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

func pageURL(n int) string {
	if n <= 1 {
		return homePage
	}
	return fmt.Sprintf("%s%v/", pagerPre, n)
}
