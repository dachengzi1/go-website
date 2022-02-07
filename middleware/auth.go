package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		now := time.Now()
		fmt.Println("hello word mid")
		c.Next()
		end := time.Since(now)
		fmt.Printf("request time: %d", end)
		return
	}
}
