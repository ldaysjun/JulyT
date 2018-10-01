package JulySpider

import (
	"app/JulySpider/Xpath"
	"app/julyNet"
	"fmt"
	"github.com/google/uuid"
)

type Spider struct {
	Parse Analysis    //处理入口
	NextStep Analysis //下一步处理
	SpiderName string //Spider名字
	SonSpider bool
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

func (spider *Spider)RunNextStep(url string,nextStep func(node *Xpath.Node,spider *Spider))  {
	nextStepSpider := new(Spider)
	req := new(julyNet.CrawlRequest)
	req.Url = url
	nextStepSpider.Request = req
	nextStepSpider.NextStep = Analysis(nextStep)
	fmt.Println("Alamofire")
	fmt.Println(nextStepSpider.NextStep)
	nextStepSpider.SonSpider = true
	nextStepSpider.Registered()
}





