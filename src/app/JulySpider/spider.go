package JulySpider

import (
	"app/JulySpider/Xpath"
	"app/julyNet"
	"github.com/google/uuid"
)

type Spider struct {
	Parse Analysis    //处理入口
	NextStep Analysis //下一步处理
	SpiderName string //Spider名字
	SonSpider bool
	Request *julyNet.CrawlRequest

	ParseHandle func(node *Xpath.Node,spider *Spider)
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
	req.NotFilter = true
	nextStepSpider.Request = req
	nextStepSpider.ParseHandle = nextStep
	nextStepSpider.SonSpider = true
	nextStepSpider.Registered()
}

func (spider *Spider)ParseHandleTest(node *Xpath.Node)  {
	//fmt.Println("到了这里")
		if spider.ParseHandle != nil {
			//fmt.Println("ParseHandle")
			spider.ParseHandle(node,spider)
		}
}





