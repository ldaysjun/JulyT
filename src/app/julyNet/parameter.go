package julyNet

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Param struct {
	url           *url.URL
	proxy         *url.URL
	method        string
	body          io.Reader
	header        http.Header
	enableCookie  bool
	dialTimeout   time.Duration
	connTimeout   time.Duration
	retryTimes    int
	retryPause    time.Duration
	redirectTimes string
	client        *http.Client
}

func CreateParam(req *CrawlRequest) (param *Param, err error) {
	param = new(Param)
	param.url, err = StrToUrl(req.Url)

	if err != nil {
		return nil, err
	}

	param.header = req.Header
	if param.header == nil {
		param.header = make(http.Header)
	}

	if req.Proxy != "" {
		param.proxy, err = url.Parse(req.Proxy)
		if err != nil {
			return nil, err
		}
	}

	param.enableCookie = req.UseCookie

	if param.dialTimeout = req.DialTimeout; param.dialTimeout < 0 {
		param.dialTimeout = 0
	}
	param.connTimeout = req.ConnTimeout
	param.retryTimes = req.RetryTimes
	param.redirectTimes = req.RedirectTimes
	param.retryPause = req.RetryPause

	method := strings.ToUpper(req.Method)
	switch method {
	case "GET":
		param.method = method
	case "POST":
		param.method = method
		param.header.Add("Content-Type", "application/x-www-form-urlencoded")
		strings.NewReader(req.PostData)
		param.body = strings.NewReader(req.PostData)

	default:
		param.method = "GET"
	}

	//默认不使用Keep-Alive
	param.header.Set("Connection", "Close")

	return param, nil
}
