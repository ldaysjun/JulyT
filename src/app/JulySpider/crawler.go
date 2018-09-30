package JulySpider

import (
	//"github.com/google/uuid"
	//"strconv"
	"app/julyNet"
	"sync"
)

var crawler *Crawler
var lock *sync.Mutex = &sync.Mutex {}


type Crawler struct {
	//资源矩阵
	matrix *Matrix
	Process map[string]*Spider
	crawlerPullHandle func(spider *Spider)	//拉取spider处理函数
	crawlerPushHandle func()				//spider入队处理函数
	CrawlerPushRequestHandle func(request *julyNet.CrawlRequest)
}


//单例方式创建
func NewCrawler() *Crawler{
	if crawler==nil {
		lock.Lock()
		defer lock.Unlock()

		crawler= new(Crawler)
		crawler.Process = make(map[string]*Spider)
		crawler.matrix = NewMatrix()
	}
	return crawler
}

//spider入队、如果需要异步入队，需要加锁
func (crawler *Crawler)PushSpider(spider *Spider)  {
	crawler.matrix.pushSpider(spider)

	if spider.SonSpider {
		if crawler.CrawlerPushRequestHandle!=nil {
			crawler.CrawlerPushRequestHandle(spider.Request)
		}
	}

	if crawler.crawlerPushHandle != nil{
		crawler.crawlerPushHandle()
	}
}

//提取spider
func (crawler *Crawler)PullSpider() {
	spider := crawler.matrix.pullSpider()

	if crawler.crawlerPullHandle != nil && spider!=nil{
		crawler.crawlerPullHandle(spider)
	}
}


//设置handle
func (crawler *Crawler)SetCrawlerHandle(crawlerPullHandle func(spider *Spider),crawlerPushHandle func())  {
	if crawlerPullHandle!=nil {
		crawler.crawlerPullHandle = crawlerPullHandle
	}
	if crawlerPushHandle!=nil {
		crawler.crawlerPushHandle = crawlerPushHandle
	}
}


//用于直连
func (crawler *Crawler)PushRequestToEngine(request *julyNet.CrawlRequest){
	if crawler.CrawlerPushRequestHandle!=nil{
		crawler.CrawlerPushRequestHandle(request)
	}
}






