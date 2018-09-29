package JulySpider

import (
	"app/julyNet"
	"github.com/google/uuid"
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
		uuidData,_ :=uuid.NewUUID()
		spider.Request.UUID = uuidData.String()

		crawler.PushSpider(spider)
	}
}





