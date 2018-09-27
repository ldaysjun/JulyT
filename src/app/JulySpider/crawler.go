package JulySpider

import (
	"fmt"
	"sync"
)

var crawler *Crawler
var lock *sync.Mutex = &sync.Mutex {}


type Crawler struct {
	spiders []*Spider						//待处理爬虫实例
	Process map[string]*Spider

	crawlerPullHandle func(spider *Spider)	//拉取spider处理函数
	crawlerPushHandle func()				//spider入队处理函数
	lock sync.Mutex							//保证线程安全
}


//单例方式创建
func NewCrawler() *Crawler{
	if crawler==nil {
		lock.Lock()
		defer lock.Unlock()

		crawler= new(Crawler)
		crawler.spiders = make([]*Spider,0)
		crawler.Process = make(map[string]*Spider)
	}
	return crawler
}

//spider入队
func (crawler *Crawler)PushSpider(spider *Spider)  {
	fmt.Println("注册成功：",spider.Request.Url)

	if spider == nil {
		return
	}

	if crawler.crawlerPushHandle != nil{
		crawler.crawlerPushHandle()
	}

	crawler.spiders = append(crawler.spiders, spider)
}

//提取spider
func (crawler *Crawler)PullSpider() {
	crawler.lock.Lock()
	defer crawler.lock.Unlock()

	n := len(crawler.spiders)
	if n<= 0 {
		return
	}

	spider := crawler.spiders[0]
	crawler.spiders = crawler.spiders[1:]

	//将spider添加到处理队列
	if _,found:=crawler.Process[spider.Request.Url];!found {
		crawler.Process[spider.Request.UUID] = spider
	}

	if crawler.crawlerPullHandle != nil {
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

func (crawler *Crawler)GetSpiders()[]*Spider{
	return crawler.spiders
}





