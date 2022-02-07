package taskPoll

import (
	"errors"
	"log"
	"sync/atomic"
)

type Task struct {
	Handler func(v ...interface{})
	Params  []interface{}
}
type Pool struct {
	capacity       uint64
	runningWorkers uint64
	state          uint64
	taskC          chan *Task
	closeC         chan bool
	PanicHandler   func(interface{})
}

var ErrInvalidPoolCap = errors.New("invalid pool cap")
var ErrPoolAlreadyClosed = errors.New("pool already closed")

const (
	RUNNING = 1
	STOPED  = 0
)

func NewPool(capacity uint64) (*Pool, error) {
	if capacity < 0 {
		return nil, ErrInvalidPoolCap
	}
	return &Pool{
		capacity: capacity,
		state:    RUNNING,
		taskC:    make(chan *Task, capacity),
		closeC:   make(chan bool),
	}, nil
}

func (p *Pool) run() {
	p.incRunning()
	go func() {
		defer func() {
			p.delRunning()
			if r := recover(); r != nil { // 恢复 panic
				if p.PanicHandler != nil { // 如果设置了 PanicHandler, 调用
					p.PanicHandler(r)
				} else { // 默认处理
					log.Printf("Worker panic: %s\n", r)
				}
			}
		}()

		for {
			select { // 阻塞等待任务、结束信号到来
			case task, ok := <-p.taskC: // 从 channel 中消费任务
				if !ok { // 如果 channel 被关闭, 结束 worker 运行
					return
				}
				// 执行任务
				task.Handler(task.Params...)
			case <-p.closeC: // 如果收到关闭信号, 结束 worker 运行
				return
			}
		}
	}()
}

func (p *Pool) Put(task *Task) error {
	if p.state == STOPED { // 如果任务池处于关闭状态, 再 put 任务会返回 ErrPoolAlreadyClosed 错误
		return ErrPoolAlreadyClosed
	}
	if p.GetRunningWorkers() < p.GetCap() { // 如果任务池满, 则不再创建 worker
		// 创建启动一个 worker
		p.run()
	}
	// 将任务推入队列, 等待消费
	p.taskC <- task
	return nil
}

func (p *Pool) Close() {
	p.state = STOPED // 设置 state 为已停止

	for len(p.taskC) > 0 { // 阻塞等待所有任务被 worker 消费
	}

	p.closeC <- true // 发送销毁 worker 信号
	close(p.taskC)   // 关闭任务队列
}

func (p *Pool) incRunning() {
	atomic.AddUint64(&p.runningWorkers, 1)
}

func (p *Pool) delRunning() {
	atomic.AddUint64(&p.runningWorkers, ^uint64(0))
}

func (p *Pool) GetRunningWorkers() uint64 {
	return atomic.LoadUint64(&p.runningWorkers)
}

func (p *Pool) GetCap() uint64 {
	return atomic.LoadUint64(&p.capacity)
}
