package JulySpider

import (
	"app/julyNet"
	"app/julyUtils/julyUuid"
)

type Spider struct {
	Parse Parser
	SpiderName string
	Request *julyNet.CrawlRequest
}

//像crawler注册
func (spider *Spider)Registered()  {

	crawler := NewCrawler()
	if spider.Request != nil {
		spider.Request.UUID = julyUuid.GenerateUuid()
		crawler.PushSpider(spider)
	}
}





