package julyEngine

import (
	"app/JulySpider"
	"app/julyNet"
	"app/julyScheduler"
	"app/julyTaskPool"
	"fmt"
	"io/ioutil"
	"net/http"
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


}

func NewEngine() *Engine {

	engine := new(Engine)

	//初始化各个组件
	engine.taskPool     = julyTaskPool.NewTaskPool(11,50,false)
	engine.requestQueue = julyScheduler.NewQueue(engine.queuePullHandle,engine.queueAfterPushHandle)
	engine.crawler      = JulySpider.NewCrawler()
	engine.crawler.SetCrawlerHandle(engine.crawlerPullHandle,engine.crawlerPushHandle)
	engine.downLoad = julyNet.NewDownLoad()
	engine.downLoad.DownFinishHandle = engine.downFinishHandle

	return engine
}

func (engine *Engine)Run(){
	//engine.listenQueue()
	engine.listenCrawler()
}


func (engine *Engine)listenQueue() {
	fmt.Println("监听queue")
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

//监听Crawler，如果有数据立刻处理
func (engine *Engine)listenCrawler(){
	fmt.Println("监听Crawler")
	crawler := engine.crawler

	engine.taskPool.SubmitTask(func() error {
		for  {
			spiders := crawler.GetSpiders()
			if len(spiders)>0 {
				crawler.PullSpider()
			}
		}
		return nil
	})
}

//添加数据到Queue
func (engine *Engine)pushRequestToQueue(request *julyNet.CrawlRequest)  {
	fmt.Println("添加数据",request.Url)
	engine.taskPool.SubmitTask(func() error {
		//time.Sleep(2*time.Second)
		engine.requestQueue.PushRequest(request)
		return nil
	})
}

//下载HTML
func (engine *Engine)downloadHTML(req *julyNet.CrawlRequest)  {
	engine.downLoad.DownLoad(req)
}

//=======================所有任务回调handle=============================================

/*Queue相关处理函数*/
//队列入队后相关操作
func (engine *Engine)queueAfterPushHandle() {
	fmt.Println("入队后处理")
	engine.requestQueue.PullRequest()
}

//队列拉取处理
func (engine *Engine)queuePullHandle(request *julyNet.CrawlRequest)  {
	fmt.Println("当前Url:",request.Url)
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
}

/*Download相关处理函数*/
//下载完成处理
func (engine *Engine)downFinishHandle(rsp *http.Response,uuid string){

	fmt.Println("uuid:",uuid)
	b,_ := ioutil.ReadAll(rsp.Body)
	spiders := engine.crawler.Process
	spider := spiders[uuid]
	spider.Parse.Parse(string(b))



	//node,err:=Xpath.ParseHTML(inputReader)
	//fmt.Println("结果")
	//
	//if err != nil {
	//	fmt.Println("xmlpath parse file failed!!!")
	//	return
	//}
	//
	//path := Xpath.MustCompile("//*[@id=\"main\"]/article[1]/header/h1/a")
	//fmt.Println(path.String(node))

}






