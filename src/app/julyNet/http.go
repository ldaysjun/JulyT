package julyNet

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"net/http/cookiejar"
	"time"
)

type JulyHttp struct {
	CookieJar *cookiejar.Jar
}

// 创建一个JulyHttp下载器
func NewJulyHttp() Downer {
	s := new(JulyHttp)
	return s
}


func (self *JulyHttp) DownLoad(request *CrawlRequest) (rsp *http.Response, err error) {
	param, err := CreateParam(request)
	if err != nil {
		return nil, err
	}

	param.client = self.createClient(param)
	rsp, err = self.httpRequest(param)

	if err != nil {
		fmt.Println(err)
	}
	return rsp,err
}

func (self *JulyHttp) createClient(param *Param) *http.Client {

	client := &http.Client{}

	transparent := &http.Transport{
		Dial: func(network, addr string) (net.Conn, error) {
			conn, err := net.DialTimeout(network, addr, param.dialTimeout)
			if err != nil {
				return nil, err
			}
			return conn, nil
		},
	}

	if param.url.Scheme == "https" {
		transparent.TLSClientConfig = &tls.Config{RootCAs: nil, InsecureSkipVerify: true}
		transparent.DisableCompression = true
	}

	if param.proxy != nil {
		transparent.Proxy = http.ProxyURL(param.proxy)
	}

	if param.enableCookie {
		client.Jar = self.CookieJar
	}

	client.Transport = transparent
	return client
}

func (self *JulyHttp) httpRequest(param *Param) (rsp *http.Response, err error) {

	req, err := http.NewRequest(param.method, param.url.String(), param.body)
	if err != nil {
		return nil, err
	}

	req.Header = param.header
	if param.retryTimes <= 0 {
		rsp, err = param.client.Do(req)
	} else {
		for i := 0; i < param.retryTimes; i++ {
			rsp, err = param.client.Do(req)
			if err != nil {
				time.Sleep(param.retryPause)
				continue
			}
			break
		}

	}

	return rsp, err
}
