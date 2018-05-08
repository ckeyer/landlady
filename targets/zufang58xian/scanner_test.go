package zufang58xian

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

// TestPageCount
func TestPageCount(t *testing.T) {
	cli := http.DefaultClient
	zufang := newZufang58xian()

	n, err := zufang.PageCount(cli)
	if err != nil {
		t.Error(err)
		return
	}
	if n != 70 {
		t.Errorf("get page count: %v", n)
	}
}

// TestScan
func TestScan(t *testing.T) {
	cli := http.DefaultClient
	zufang := newZufang58xian()

	urls, err := zufang.ScanURLs(cli, 1)
	if err != nil {
		t.Error(err)
		return
	}
	if len(urls) == 0 {
		t.Error("0 urls")
	}
	for i, v := range urls {
		t.Logf("%v: %v", i+1, v)
	}

	// t.Error("...")
}

// TestJump
func TestJump(t *testing.T) {
	return
	t.Error("...")
	cli := http.DefaultClient
	_ = "http://xa.58.com/zufang/34021339544745x.shtml"
	jurl := `http://jxjump.58.com/service?target=FCADV8oV3os7xtAhI2suhvPnTEK_30M80Za1mFz42mTmlhpHZpk1zEffDjpdRkNz3Q5xoKYl4Bi0ja0Sib7CGDpBCM0sqdsF1EjPRyBP1ZKMCGY1x3hFqeDuk1c7B5bZPVuVAeEvZYetOdm5HKMZL2LxkyIukfCbRGVaWhwAwIAsnVFXVOe12UF_gQYsiu1SCX0Xj4nhW2er_V8Lqir8uSjtm3EgnTnuV4ut2oMzJts5psgXUtKF7IOhn7w&local=483&pubid=32705457&apptype=0&psid=169432669199989300131892197&entinfo=34021339544745_0&cookie=||http%3A%2F%2Fwww.baidu.com%2Flink%3Furl%3DxtfRkxkUxCCcbgyrwhv6uj_05yvkwqkCCdnxAVMTUlaYRt-NH7joLHQJYOdl3pgL%26ck%3D10016.3.9.233.148.234.141.1432%26shh%3Dwww.baidu.com%26sht%3Dbaidu%26wd%3D%26eqid%3Db3fb000300033c18000000025af1caab|c5/nn1rxyrFTka/rGgpBAg==&fzbref=0&key=&params=jxzfbestpc^desc&from=1-list-1`
	ju, _ := url.Parse(jurl)
	for k, v := range ju.Query() {
		t.Logf("%s: %+v", k, v)
	}

	resp, err := cli.Get(jurl)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("status; %+v", resp)
	return
	bs, _ := ioutil.ReadAll(resp.Body)
	t.Logf("body: %s", bs)
	t.Error("...")
}
