package julyScheduler

import (
	"app/julyNet"
	"sync"
	"sync/atomic"
)

const (
	HighPriority = 1
	MediumPriority = 2
	LowPriority = 3
)

type Matrix struct {
	requests   map[int][]*julyNet.CrawlRequest
	priorities []int
	resCount   int32
	lock       sync.Mutex
	dupeFilter *DupeFilter
}

func NewMatrix() *Matrix {

	matrix := new(Matrix)
	matrix.requests = make(map[int][]*julyNet.CrawlRequest)
	matrix.priorities = []int{HighPriority,MediumPriority,LowPriority}
	matrix.resCount = 0
	matrix.dupeFilter = NewDupeFilter(nil)
	return matrix

}


//添加请求
func (matrix *Matrix)addRequest(request *julyNet.CrawlRequest)  {
	matrix.lock.Lock()
	defer matrix.lock.Unlock()

	if request.Priority == 0 {
		request.Priority = 1
	}

	//过滤检验，避免重复下载，设置NotFilter可以不检验，
	if !request.NotFilter {
		isRepeat:=matrix.dupeFilter.filter(request)
		if isRepeat {
			return
		}
	}

	priority := request.Priority
	if _,found:=matrix.requests[priority];!found {
		matrix.requests[priority] = []*julyNet.CrawlRequest{}
	}

	//添加请求到队列
	matrix.requests[priority] = append(matrix.requests[priority], request)
	atomic.AddInt32(&matrix.resCount,1)
}


//取出请求
func (matrix *Matrix)pullRequest() *julyNet.CrawlRequest{
	matrix.lock.Lock()
	defer matrix.lock.Unlock()
	var request *julyNet.CrawlRequest

	
	//按照优先级出队
	for i:=0;i<len(matrix.priorities);i++  {
		priority := matrix.priorities[i]
		if len(matrix.requests[priority])>0 {

			//取出队首元素
			request = matrix.requests[priority][0]
			matrix.requests[priority] = matrix.requests[priority][1:]
			atomic.AddInt32(&matrix.resCount,-1)
			break
		}
	}
	return request
}