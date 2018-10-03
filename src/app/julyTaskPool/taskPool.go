package julyTaskPool

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type closeSignal struct{}
type funcTask func() error

var (
	ErrorTaskPoolClose = errors.New("this taskPool has been closed")
	ErrorGCNodeIsNil   = errors.New("this Node is nil")
)

type TaskNode struct {
	pool   *TaskPool
	task   chan funcTask
	isIdel bool
	//最近最后使用时间
	recentUsageTime time.Time
}

func (t *TaskNode) run() {
	go func() {

		for task := range t.task {
			if task == nil {
				//fmt.Println("break")
				break
			}

			if t.isIdel {
				//fmt.Println("使用复用节点")
			}
			//执行任务
			task()
			//回收任务节点
			t.pool.taskNodeGC(t)
		}
	}()
}

//=======================================================

//任务池
type TaskPool struct {

	//任务池大小
	poolSize int32
	//过期时间
	expiredDuration time.Duration
	//空闲任务节点
	taskNodes []*TaskNode
	//运行中节点数
	running int32
	//TaskPool关闭通知
	closeNotice chan closeSignal
	//是否使用缓冲队列
	isUseCache bool
	//TaskPool 缓冲队列
	cacheTasks []funcTask
	//lock
	lock sync.Mutex
	//once
	once sync.Once
}

func NewTaskPool(poolSize int32, expiredDuration int, isUseCache bool) *TaskPool {

	if poolSize <= 0 {
		poolSize = 10
	}

	if expiredDuration <= 0 {
		expiredDuration = 20
	}

	pool := new(TaskPool)
	pool.poolSize = poolSize
	pool.expiredDuration = time.Duration(expiredDuration) * time.Second
	pool.closeNotice = make(chan closeSignal, 1)
	pool.cacheTasks = make([]funcTask, 0)
	pool.isUseCache = isUseCache

	//启动定时回收
	go pool.idleNodeGC()
	if isUseCache {
		go pool.listenCache()
	}

	return pool
}

//提交任务，由内部调度分配
func (p *TaskPool) SubmitTask(task funcTask) error {
	if len(p.closeNotice) > 0 {
		return ErrorTaskPoolClose
	}

	if p.isUseCache {
		//fmt.Println("使用cache")
		p.cacheTask(task)
	} else {
		taskNode := p.getTaskNode()
		taskNode.task <- task
	}

	return nil
}

func (p *TaskPool) cacheTask(task funcTask) {
	p.lock.Lock()
	defer p.lock.Unlock()

	p.cacheTasks = append(p.cacheTasks, task)
}

func (p *TaskPool) listenCache() {
	for {
		if len(p.cacheTasks) <= 0 {
			continue
		}
		task := p.cacheTasks[0]
		p.cacheTasks[0] = nil
		p.cacheTasks = p.cacheTasks[1:]

		taskNode := p.getTaskNode()
		taskNode.task <- task
	}
}

//===========================内部私有====================================
func (p *TaskPool) getTaskNode() *TaskNode {
	var t *TaskNode

	if len(p.taskNodes) > 0 {
		//取出空闲队列最后可用节点
		//fmt.Println("存在空闲队列：%d", len(p.taskNodes))
		t = p.getNodeFromTaskNodes()
	} else {
		//当前任务池还有空间
		if p.Running() < p.PoolSize() {
			t = new(TaskNode)
			t.pool = p
			t.task = make(chan funcTask, 1)
			t.run()
			p.runningIncrease()
		} else {
			//阻塞等待
			//fmt.Println("进入等待")
			for {
				p.lock.Lock()
				if len(p.taskNodes) <= 0 {
					//fmt.Println("等待任务节点")
					p.lock.Unlock()
					continue
				}
				//fmt.Println("获取任务节点")
				t = p.getNodeFromTaskNodes()
				p.lock.Unlock()

				break
			}
		}
	}

	return t
}

//从空闲任务列表中获取任务节点,并且移除
func (p *TaskPool) getNodeFromTaskNodes() *TaskNode {

	tempNodes := p.taskNodes
	n := len(p.taskNodes)
	if n <= 0 {
		return nil
	}

	t := tempNodes[n-1]
	tempNodes[n-1] = nil
	p.taskNodes = tempNodes[:n-1]

	return t
}

//返回空闲队列
func (p *TaskPool) taskNodeGC(node *TaskNode) error {

	if node == nil {
		return ErrorGCNodeIsNil
	}

	p.lock.Lock()
	defer p.lock.Unlock()

	fmt.Println("回收节点")

	node.recentUsageTime = time.Now()
	node.isIdel = true
	p.taskNodes = append(p.taskNodes, node)

	return nil
}

//空闲节点回收
func (p *TaskPool) idleNodeGC() {

	ticker := time.NewTicker(p.expiredDuration)
	tempNodes := p.taskNodes
	for _= range ticker.C {

		p.lock.Lock()
		defer p.lock.Unlock()

		nowTime := time.Now()
		if len(p.closeNotice) <= 0 {
			return
		}
		n := 0
		for i, t := range tempNodes {
			if nowTime.Sub(t.recentUsageTime) <= p.expiredDuration {
				break
			}
			//fmt.Println("开始回收")
			n = i
			t.task <- nil
			tempNodes[i] = nil
		}
		if n+1 >= len(tempNodes) {
			p.taskNodes = tempNodes[:0]
		} else {
			p.taskNodes = tempNodes[n+1:]
		}
	}
}

func (p *TaskPool) CloseTaskPool() {
	p.once.Do(func() {
		p.closeNotice <- closeSignal{}
		p.lock.Lock()
		defer p.lock.Unlock()
		tempTaskNodes := p.taskNodes
		for index, taskNode := range tempTaskNodes {
			taskNode.task = nil
			tempTaskNodes[index] = nil
		}
		p.taskNodes = nil
	})
}

//对运行数做原子操作+1
func (p *TaskPool) runningIncrease() {
	atomic.AddInt32(&p.running, 1)
}

//对运行数做原子操作-1
func (p *TaskPool) runningDec() {
	atomic.AddInt32(&p.running, -1)
}

//原子操,获取运行数量
func (p *TaskPool) Running() int {
	return int(atomic.LoadInt32(&p.running))
}

//原子操,获取任务池大小
func (p *TaskPool) PoolSize() int {
	return int(atomic.LoadInt32(&p.poolSize))
}
