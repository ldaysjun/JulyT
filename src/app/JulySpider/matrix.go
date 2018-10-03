package JulySpider

import (
	"sync"
	"sync/atomic"
)

type Matrix struct {
	spiders    []*Spider
	priorities []int
	resCount   int32
	process    map[string]*Spider
	lock       sync.Mutex
}


func NewMatrix() *Matrix {

	matrix := new(Matrix)
	matrix.spiders = make([]*Spider,0)
	matrix.resCount = 0
	return matrix
}


//添加spider
func (matrix *Matrix)pushSpider(spider *Spider){
	matrix.lock.Lock()
	defer matrix.lock.Unlock()

	if !spider.SonSpider {
		//添加请求到队列
		matrix.spiders = append(matrix.spiders, spider)
		atomic.AddInt32(&matrix.resCount,1)
	}else {
		//将数据加到待处理
		if _,found:=crawler.Process[spider.Request.UUID];!found {
			crawler.Process[spider.Request.UUID] = spider
		}
	}
}

func (matrix *Matrix)pullSpider() *Spider{
	matrix.lock.Lock()
	defer matrix.lock.Unlock()

	if len(matrix.spiders)<=0 {
		return nil
	}

	//取出队首元素
	spider := matrix.spiders[0]
	matrix.spiders = matrix.spiders[1:]

	//将数据加到待处理
	if _,found:=crawler.Process[spider.Request.UUID];!found {
		crawler.Process[spider.Request.UUID] = spider
	}

	atomic.AddInt32(&matrix.resCount,-1)
	return spider
}

func (matrix *Matrix)getSpiders() []*Spider{


	return matrix.spiders
}


