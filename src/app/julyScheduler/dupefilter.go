package julyScheduler

import (
	"app/julyNet"
	"app/julyUtils/julySet"
	"crypto/sha1"
	"encoding/hex"
)

var (
	//指纹集合
	fingerSet *julySet.Set
)

type DupeFilter struct {
	Filter func(request *julyNet.CrawlRequest)
}

func NewDupeFilter(filter func(request*julyNet.CrawlRequest)) *DupeFilter{

	fingerSet = julySet.New()
	dupeFilter := new(DupeFilter)

	if filter!=nil {
		dupeFilter.Filter = filter
	}

	return dupeFilter
}

func (f *DupeFilter)RequestFilter(request *julyNet.CrawlRequest)  {
	if f.Filter == nil {
		f.filter(request)
	}else {
		f.Filter(request)
	}
}

func (f *DupeFilter)filter(request *julyNet.CrawlRequest) bool {

	//生成指纹，过滤相同URL
	h := sha1.New()
	h.Write([]byte(request.Url))
	bs:= h.Sum(nil)

	finger := hex.EncodeToString(bs)
	if !fingerSet.Contains(finger){
		fingerSet.Add(finger)
		return false
	}
	return true
}









