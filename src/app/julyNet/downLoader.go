package julyNet

import (
	"net/http"
)


type Downloader struct {
	julyHttp Downer
	DownFinishHandle func(rsp *http.Response,uuid string)
}

func NewDownLoad() *Downloader{
	d := new(Downloader)
	d.julyHttp = NewJulyHttp()
	return d
}

func (d *Downloader)DownLoad(req *CrawlRequest) (rsp *http.Response, err error) {
	rsp, err = d.julyHttp.DownLoad(req)
	if d.DownFinishHandle !=nil {
		d.DownFinishHandle(rsp,req.UUID)
	}
	return rsp,err
}

type Downer interface {
	DownLoad(request *CrawlRequest) (rsp *http.Response, err error)
}
