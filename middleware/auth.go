package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-website/db"
	"go-website/model"
	"go-website/res"
	"net/http"
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

func ApiAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := c.Cookie("userId")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, res.ParamErr.Error("Invalid UserId"))
			return
		}
		conn := db.Db
		var user model.User
		err = conn.Where("id = ?", userId).First(&user).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, res.ParamErr.Error("Invalid user"))
			return
		}
		fmt.Println("user id:", userId, user)
		c.Set("userId", userId)
		c.Set("user", user)
		return
	}
}
