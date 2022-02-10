package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "go-website/db"
	"go-website/model"
	"net/http"
	"strconv"
)

func CreateExercise(c *gin.Context) {
	var postExercise model.PostExercise
	db := Db
	err := c.ShouldBind(&postExercise)
	if err != nil {
		fmt.Println("create question err:", err)
		return
	}

	ex := model.Exercise{
		Title:       postExercise.Title,
		Description: postExercise.Description,
		Deleted:     postExercise.Deleted,
		Type:        postExercise.Type,
	}
	result := db.Save(&ex)
	if result.Error != nil {
		fmt.Printf("insert Exercise err:", err)
	}

	fmt.Println(postExercise, "ExerciseId：", ex.Id)

	var questionGroups []model.QuestionGroup
	fmt.Println(postExercise.QuestionIds)
	for k, v := range postExercise.QuestionIds {
		//todo? 未判断question 是否存在
		fmt.Printf("question index: %d, value: %d", k, v)
		questionGroups = append(questionGroups, model.QuestionGroup{
			ExerciseId: ex.Id,
			QuestionId: v,
		})
	}
	err = db.Create(&questionGroups).Error
	if err != nil {
		fmt.Printf("insert question Groups err:", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
	})
	return
}

func GetExercise(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var exs []model.Exercise
	db := Db
	err := db.First(&exs, id).Error
	if err != nil {
		fmt.Printf("find question group  err:", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"code":   0,
		"result": exs,
	})
	return

	//	id, _ := strconv.Atoi(c.Param("id"))
	//	var questionGroup model.QuestionGroup
	//	var question model.Question
	//	var wg sync.WaitGroup
	//	ch := make(chan *model.QuestionGroup)
	//	qch := make(chan *model.Question)
	//	defer close(ch)
	//	defer close(qch)
	//
	//	obj := make(map[string]interface{})
	//	wg.Add(1)
	//	go questionGroup.FindOne(id, ch)
	//	go question.FindOneById(id, qch)
	//LOOP:
	//	for {
	//		select {
	//		case r := <-ch:
	//			fmt.Println("r.Id：", r.Id)
	//			obj["id"] = r.Id
	//			ch = nil
	//
	//		case r := <-qch:
	//			fmt.Println("r.Context：", r.Context)
	//			obj["context"] = r.Context
	//			qch = nil
	//
	//		case <-time.After(time.Second * 1):
	//			fmt.Println("req timeout")
	//			return
	//		}
	//		if ch == nil && qch == nil{
	//			break LOOP
	//		}
	//	}
	//	fmt.Println("for select close：")
	//	wg.Done()
	//	wg.Wait()
	//
	//	c.JSON(http.StatusOK, gin.H{
	//		"code":   0,
	//		"result": obj,
	//	})
	//	return
}
