package julyNet

import (
	"net/http"
)


type Downloader struct {
	julyHttp Downer
	downFinishHandle func(rsp *http.Response)
}

func NewDownLoad() *Downloader{
	d := new(Downloader)
	d.julyHttp = NewJulyHttp()
	return d
}

func (d *Downloader)DownLoad(req *CrawlRequest) (rsp *http.Response, err error) {
	rsp, err = d.julyHttp.DownLoad(req)
	if d.downFinishHandle !=nil {
		d.downFinishHandle(rsp)
	}
	return
}

type Downer interface {
	DownLoad(request *CrawlRequest) (rsp *http.Response, err error)
}
