package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "go-website/db"
	"go-website/model"
	"net/http"
	"strings"
	"time"
)

func GetConfig(c *gin.Context) {
	interData, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
		})
		return
	}
	u := interData.(model.User)
	fmt.Println("user:", u, u.Id)

	db := Db

	var count int64
	var myExercise model.MyExercise

	err := db.Where("user_id = ? ", u.Id).Find(&myExercise).Count(&count).Error

	if err != nil {
		fmt.Println("find my exercise err:", err)
		panic(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": gin.H{
			"version":         "v1.0.0",
			"upgrade":         strings.Compare("v1.0.2", "v1.0.1"), //是否升级
			"myExerciseCount": count,
			"user": gin.H{
				"userId":   u.Id,
				"mobile":   u.Mobile,
				"nickName": u.NickName,
			},
			"currentTime": time.Now().UnixMilli(),
		},
	})
	return
}
