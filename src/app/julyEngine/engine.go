package julyEngine

import (
	"app/JulySpider"
	"app/JulySpider/Xpath"
	"app/julyNet"
	"app/julyScheduler"
	"app/julyTaskPool"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

const(
	//PauseEngine = 0
)

type Engine struct {

	taskPool     *julyTaskPool.TaskPool   	//任务池
	requestQueue *julyScheduler.Queue	    //请求队列
	crawler      *JulySpider.Crawler        //爬虫服务
	downLoad     *julyNet.Downloader
	status       int	                    //Engine状态
	lock         sync.Mutex
}

func NewEngine() *Engine {

	engine := new(Engine)

	//初始化各个组件
	engine.taskPool     = julyTaskPool.NewTaskPool(2,50,false)
	engine.requestQueue = julyScheduler.NewQueue(engine.queuePullHandle,engine.queueAfterPushHandle)
	engine.crawler      = JulySpider.NewCrawler()
	engine.crawler.SetCrawlerHandle(engine.crawlerPullHandle,engine.crawlerPushHandle)
	engine.downLoad = julyNet.NewDownLoad()
	engine.downLoad.DownFinishHandle = engine.downFinishHandle

	return engine
}

func Run() *Engine{
	log.Println("启动JulyT")
	engine:=NewEngine()
	return engine
}

func (engine *Engine)listenQueue() {
	queue := engine.requestQueue
	engine.taskPool.SubmitTask(func() error {
		for  {
			if queue.MatrixSize()>0 {
				queue.PullRequest()
			}
		}
		return nil
	})
}

func (engine *Engine)CloseEngine()  {
	engine.taskPool.CloseTaskPool()
}

//监听Crawler，如果有数据立刻处理
func (engine *Engine)listenCrawler(){
	//fmt.Println("监听Crawler")
	//crawler := engine.crawler
	//
	//engine.taskPool.SubmitTask(func() error {
	//	for  {
	//		spiders := crawler.GetSpiders()
	//		if len(spiders)>0 {
	//			crawler.PullSpider()
	//		}
	//	}
	//	return nil
	//})
}

//添加数据到Queue
func (engine *Engine)pushRequestToQueue(request *julyNet.CrawlRequest)  {
	//fmt.Println("添加数据",request.Url)
	//engine.taskPool.SubmitTask(func() error {
	//	//time.Sleep(2*time.Second)
	//
	//	return nil
	//})
	engine.requestQueue.PushRequest(request)
}

//下载HTML
func (engine *Engine)downloadHTML(req *julyNet.CrawlRequest)  {
	engine.downLoad.DownLoad(req)
}

//=======================所有任务回调handle=============================================

/*Queue相关处理函数*/
//队列入队后相关操作
func (engine *Engine)queueAfterPushHandle() {
	engine.requestQueue.PullRequest()
}

//队列拉取处理
func (engine *Engine)queuePullHandle(request *julyNet.CrawlRequest)  {
	engine.downloadHTML(request)
}


/*Crawler相关处理函数*/
//提取spider处理
func (engine *Engine)crawlerPullHandle(spider *JulySpider.Spider)  {
	//fmt.Println("crawlerPullHandle 当前id:",GoID(),"|","uuid:",spider.Request.UUID)
	engine.pushRequestToQueue(spider.Request)
}
//spider入队处理
func (engine *Engine)crawlerPushHandle() {
	fmt.Println("入队")

	engine.taskPool.SubmitTask(func() error {
		//spiders := engine.crawler.GetSpiders()
		//if len(spiders)>0 {
		//
		//}
		engine.crawler.PullSpider()
		return nil
	})
}

/*Download相关处理函数*/
//下载完成处理
func (engine *Engine)downFinishHandle(rsp *http.Response,uuid string){
	body,err:= ioutil.ReadAll(rsp.Body)
	inputReader := strings.NewReader(string(body))
	node,err:=Xpath.ParseHTML(inputReader)


	if err != nil {
		fmt.Println("xmlpath parse file failed!!!")
		return
	}
	if err!=nil {
		log.Println(err.Error())
	}
	spiders := engine.crawler.Process
	spider := spiders[uuid]
	if spider != nil && spider.Parse!=nil{
		spider.Parse.Parse(node,spider)
	}
	if spider !=nil && spider.NextStep!=nil {
		spider.NextStep.Next(node,spider)
	}
}




func GoID() int {
	//var buf [64]byte
	//n := runtime.Stack(buf[:], false)
	//idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	//id, err := strconv.Atoi(idField)
	//if err != nil {
	//	panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	//}
	return 0
}



