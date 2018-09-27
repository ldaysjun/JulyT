package JulySpider

import (
	"app/julyNet"
	"fmt"
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
		//uuidData,_ :=uuid.NewUUID()
		//spider.Request.UUID = uuidData.String()
		fmt.Println("注册uuid：",spider.Request.UUID)
		fmt.Println("spidername:",spider.SpiderName)
		crawler.PushSpider(spider)
	}
}





