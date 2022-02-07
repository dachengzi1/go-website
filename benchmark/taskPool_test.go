package benchmark

import (
	//"context"
	"sync/atomic"
	"testing"
	"goschool/taskPoll"
)

var sum int64

func demoTask(v ...interface{}) {
	for i := 0; i < 100; i++ {
		atomic.AddInt64(&sum, 1)
	}
}

var runTimes = 1000000
// 原生 goroutine
func BenchmarkGoroutineSetTimes(b *testing.B) {

	for i := 0; i < runTimes; i++ {
		go demoTask()
	}
}
// 使用协程池
func BenchmarkPutSetTimes(b *testing.B) {
	pool, err := taskPoll.NewPool(20)
	if err != nil {
		b.Error(err)
	}

	//ctx := context.Background()
	task := &taskPoll.Task{
		Handler: demoTask,
	}

	for i := 0; i < runTimes; i++ {
		pool.Put(task)
	}
}
