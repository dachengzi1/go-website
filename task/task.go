package task

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

func addAllTag() bool {
	time.Sleep(time.Second * 20)
	fmt.Println("add all tag", time.Now())
	return true
}
func cleanAllTag() bool {
	time.Sleep(time.Second * 20)
	fmt.Println("clean all tag", time.Now())
	return true
}
func init() {
	fmt.Println("cron start")
	c := cron.New()
	//c.AddFunc("*/10 * * * * *", func() {
	//	addAllTag()
	//})
	//c.AddFunc("*/5 * * * * *", func() {
	//	cleanAllTag()
	//})
	//在当前执行的程序中启动 Cron 调度程序。其实这里的主体是 goroutine + for + select + timer 的调度控制哦
	c.Start()

	//t1 := time.NewTimer(time.Second * 10)
	//4、time.NewTimer + for + select + t1.Reset
	//	如果你是初学者，大概会有疑问，这是干嘛用的？
	//	（1）time.NewTimer
	//	会创建一个新的定时器，持续你设定的时间 d 后发送一个 channel 消息
	//	（2）for + select
	//		阻塞 select 等待 channel
	//		（3）t1.Reset
	//		会重置定时器，让它重新开始计时
	//		（注意，本文适用于 “t.C已经取走，可直接使用 Reset”）

	//for {
	//	select {
	//	case <-t1.C:
	//		t1.Reset(time.Second * 10)
	//	}
	//}
}
