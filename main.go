package main

import (
	"github.com/gin-gonic/gin"
	"goschool/cache"
	_ "goschool/config"
	"goschool/db"
	"goschool/middleware"
	"goschool/router"
	_ "goschool/task"
)

func main() {

	r := gin.Default()

	db.InitDb()       //链接数据库
	cache.RedisConn() //链接redis

	r.Use(middleware.Auth())
	router.InitRouter(r)

	r.Run(":8091")

	////创建任务池
	//pool, err := taskPoll.NewPool(200)
	//if err!= nil {
	//	panic(err)
	//}
	//for i:=0; i<1000; i++{
	//	// 任务放入池中
	//	pool.Put(&taskPoll.Task{
	//		Handler: func(v ...interface{}) {
	//			fmt.Println(v)
	//		},
	//		Params: []interface{}{i},
	//	})
	//}
	//time.Sleep(1e9)

	//taskC := make(chan string, 10)
	//time.Sleep(time.Second * 2)
	//taskC <- "hha"
	//for {
	//	select {
	//	case data, ok := <-taskC:
	//		if !ok {
	//			return
	//		}
	//		fmt.Println("channel:", data)
	//	default:
	//		fmt.Printf("")
	//	}
	//}

}
