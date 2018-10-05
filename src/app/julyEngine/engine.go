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
	engine.taskPool     = julyTaskPool.NewTaskPool(100,50,false)
	engine.requestQueue = julyScheduler.NewQueue(engine.queuePullHandle,engine.queueAfterPushHandle)
	engine.crawler      = JulySpider.NewCrawler()
	engine.crawler.SetCrawlerHandle(engine.crawlerPullHandle,engine.crawlerPushHandle)
	engine.crawler.CrawlerPushRequestHandle = engine.pushRequestToQueue

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

}

//添加数据到Queue
func (engine *Engine)pushRequestToQueue(request *julyNet.CrawlRequest)  {
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
	engine.pushRequestToQueue(spider.Request)
}
//spider入队处理
func (engine *Engine)crawlerPushHandle() {
	fmt.Println("入队")

	engine.taskPool.SubmitTask(func() error {
		engine.crawler.PullSpider()
		return nil
	})
}

/*Download相关处理函数*/
//下载完成处理
func (engine *Engine)downFinishHandle(rsp *http.Response,uuid string){
	engine.lock.Lock()
	body,err:= ioutil.ReadAll(rsp.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}

	inputReader := strings.NewReader(string(body))
	node,err:=Xpath.ParseHTML(inputReader)
	if err!=nil {
		log.Println(err.Error())
		//return
	}
	spiders := engine.crawler.Process
	spider := spiders[uuid]
	engine.lock.Unlock()

	if spider != nil{
		spider.ParseHandleTest(node)
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



