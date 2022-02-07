package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "go-website/db"
	"go-website/model"
	"net/http"
	"strconv"
	"time"
)

func CreateQuestion(c *gin.Context) {

	var question model.Question

	err := c.ShouldBind(&question)
	if err != nil {
		fmt.Println("create question err:", err)
		return
	}
	question.CreatedAt = time.Now()
	question.UpdatedAt = time.Now()

	fmt.Println(question)
	err = question.Insert()
	if err != nil {
		fmt.Printf("insert question err:", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
	})
	return
}

func SearchQuestions(c *gin.Context) {
	questions := make([]model.Question, 0)
	var total int64
	db := Db

	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))

	if year, exits := c.GetQuery("year"); exits == true {
		fmt.Println("year:", year)
		db = db.Where("year = ?", year)
	}
	if page > 0 && pageSize > 0 {
		db = db.Limit(pageSize).Offset((page - 1) * pageSize)
	}
	db.Order("year asc")
	db.Find(&questions).Count(&total)
	c.JSON(http.StatusOK, gin.H{
		"code":      0,
		"questions": questions,
		"total":     total,
	})
	return
}

func UpdateQuestions(c *gin.Context) {
	db := Db
	id := c.Param("id")
	var question model.Question

	//更新 1
	//var qInter interface{}
	//
	//body, _ := ioutil.ReadAll(c.Request.Body)
	//err := json.Unmarshal(body, &qInter)
	//if err != nil {
	//	fmt.Println("update err:", err)
	//}
	//fmt.Println(body, qInter)
	//db.Model(&question).Where("id = ?", id).Updates(qInter)

	//更新 2
	//todo? 没有检查question是否存在
	err := db.First(&question, id).Error
	if err != nil {
		fmt.Println("update err:", err)
	}
	err = c.BindJSON(&question)
	if err != nil {
		fmt.Println("bind json err:", err)
	}
	db.Save(&question)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
	})
	return
}
