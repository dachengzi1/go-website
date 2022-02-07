package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	"go-website/config"
	"time"
)

var (
	rdb  *redis.Client
	ip   = config.Conf.Redis.Ip
	port = config.Conf.Redis.Port
	db = config.Conf.Redis.Db
)

//func getRedisAddr() (s string){
//	var buf bytes.Buffer
//	buf.WriteString(ip)
//	buf.WriteString(":")
//	buf.WriteString(string(port))
//    s = buf.String()
//    return
//}
func RedisConn() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", ip, port),
		Password: "", // no password set
		DB:       db,  // use default DB
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		fmt.Println("redis client err:", err.Error())
		return err
	}
	fmt.Println("redis client success")
	return
}

func Set(key string, data interface{}, expire time.Duration) (err error) {
	err = rdb.Set(key, data, expire).Err()
	if err != nil {
		fmt.Printf("redis set err: %v", err)
		return
	}
	return
}

func Get(key string) (data string, err error) {
	data, err = rdb.Get(key).Result()
	if err != nil {
		fmt.Printf("redis set err: %v", err)
		return
	}
	return
}

//a = 'eee'
//b = 'fff'
//拼接字符串的方法
//1、a + ":" + b
//2、fmt.Sprintf
//3、string.Jon([]interface{a,b},":")
//4、
//var buffer bytes.Buffer
//buffer.WriteString(a)
//buffer.WriteString(":")
//buffer.WriteString(b)