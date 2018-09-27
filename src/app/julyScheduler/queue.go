package julyScheduler

import (
	"app/julyNet"
)

const(
	//暂停队列
	PassQueue = 0
	//启动队列
	StartQueue = 1
)


type PullRequestHandle func(request *julyNet.CrawlRequest) //拉取请求处理
type PushRequestHandle func()                              //添加请求处理

type Queue struct {

	//运行状态
	status int
	//资源矩阵
	matrix *Matrix
	//资源矩阵大小
	matrixSize int

	pullHandle PullRequestHandle
	pushHandle PushRequestHandle
}

func NewQueue(pullHandle PullRequestHandle,pushHandle PushRequestHandle)  *Queue{

	queue := new(Queue)
	queue.matrix = NewMatrix()
	queue.pullHandle = pullHandle
	queue.pushHandle = pushHandle
	queue.status = StartQueue

	return queue
}

func (queue *Queue)PushRequest(request *julyNet.CrawlRequest)  {
	queue.matrix.addRequest(request)

	if queue.pushHandle != nil {
		queue.pushHandle()
	}
}


func (queue *Queue)PullRequest() (request *julyNet.CrawlRequest){

	request =queue.matrix.pullRequest()

	if request!=nil && queue.pullHandle!=nil{
		queue.pullHandle(request)
	}

	return request
}


func (queue *Queue)MatrixSize() int32{
	return queue.matrix.resCount
}





