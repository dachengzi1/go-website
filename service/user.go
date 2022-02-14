package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-website/db"
	"go-website/model"
	"net/http"
)

func SaveUser(c *gin.Context) {
	var u model.User
	err := c.ShouldBind(&u)
	if err != nil {
		fmt.Println("should bind user err:", err)
	}

	fmt.Println("user", u.NickName)
	err = db.Db.Create(&u).Error
	if err != nil {
		fmt.Println("create user err:", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
	})
	return
}
