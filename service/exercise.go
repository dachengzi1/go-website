package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "go-website/db"
	"go-website/model"
	"go-website/res"
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

	//todo ? 事物处理
	var questions []model.Question

	db.Where("(id) IN ?", postExercise.QuestionIds).Find(&questions)
	if err != nil {
		fmt.Printf("insert question Groups err:", err)
	}

	ex := model.Exercise{
		Title:       postExercise.Title,
		Description: postExercise.Description,
		Deleted:     postExercise.Deleted,
		Type:        postExercise.Type,
		Questions:   questions,
	}
	result := db.Save(&ex)
	if result.Error != nil {
		fmt.Printf("insert Exercise err:", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
	})
	return
}

func GetExercise(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var ex model.Exercise
	db := Db

	ex.Id = id
	result := db.Preload("Questions").Find(&ex)

	if result.Error != nil {
		fmt.Printf("find question group  err:", result.Error)
	}
	c.JSON(http.StatusOK, gin.H{
		"code":     0,
		"exercise": ex,
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

func SaveUserExercise(c *gin.Context) {
	db := Db
	interData, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
		})
		return
	}
	u := interData.(model.User)

	var SubMitExercise model.SubMitExercise
	var qIds = make([]int, len(SubMitExercise.MyQuestions))
	var questions []model.Question

	err := c.ShouldBindJSON(&SubMitExercise)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.MarshalJsonErr)
		return
	}

	for _, v := range SubMitExercise.MyQuestions {
		qIds = append(qIds, v.QuestionId)
	}

	err = db.Where("(id) IN ?", qIds).Find(&questions).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, res.ParamErr)
		return
	}

	questionObj := make(map[int]interface{}, len(questions))

	for _, v := range questions {
		questionObj[v.Id] = v
	}
	userScore := 0
	correctCount := 0
	wrongCount := 0
	myQuestion := make([]model.MyQuestion, len(SubMitExercise.MyQuestions))
	for _, v := range SubMitExercise.MyQuestions {
		interQ, _ := questionObj[v.QuestionId]
		question := interQ.(model.Question)
		score := 0
		status := 0
		if question.Answer == v.UserAnswer {
			score = int(question.Score)
			userScore = userScore + score
			correctCount++
			status = 1
		} else {
			wrongCount++
		}
		myQuestion = append(myQuestion, model.MyQuestion{
			UserId:     u.Id,
			QuestionId: v.QuestionId,
			UserScore:  score,
			UserAnswer: v.UserAnswer,
			Status:     status,
		})
	}

	ex := model.MyExercise{
		ExerciseId:   SubMitExercise.ExerciseId,
		UserId:       u.Id,
		Status:       "finished",
		Score:        userScore,
		CorrectCount: correctCount,
		WrongCount:   wrongCount,
		MyQuestions:  myQuestion,
	}
	result := db.Save(&ex)
	if result.Error != nil {
		fmt.Printf("save Exercise err:", result.Error)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
	})
	return
}

func GetUserExercise(c *gin.Context) {
	interData, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
		})
		return
	}
	u := interData.(model.User)
	fmt.Println("user:", u, u.Id)

	id, _ := strconv.Atoi(c.Param("id"))
	var ex model.MyExercise
	db := Db

	ex.Id = id
	ex.UserId = u.Id
	result := db.Preload("MyQuestions", "user_id = ? ", u.Id).Find(&ex)
	if result.Error != nil {
		fmt.Printf("find my question group  err:", result.Error)
	}

	var questions []model.Question
	var qIds = make([]int, len(ex.MyQuestions))
	for _, v := range ex.MyQuestions {
		qIds = append(qIds, v.QuestionId)
	}

	err := db.Where("(id) IN ?", qIds).Find(&questions).Error
	if err != nil {
		fmt.Printf("find my question group  err:", result.Error)
	}

	questionObj := make(map[int]interface{}, len(questions))

	for _, v := range questions {
		questionObj[v.Id] = v
	}

	var myQuestions []model.MyQuestionDetail

	fmt.Println("len(ex.MyQuestions):", len(ex.MyQuestions))
	for _, v := range ex.MyQuestions {
		interQ, _ := questionObj[v.QuestionId]
		question := interQ.(model.Question)
		fmt.Println("ex.MyQuestions  id:", v.Id)
		myQuestions = append(myQuestions, model.MyQuestionDetail{
			Id:         v.Id,
			Context:    question.Context,
			Question:   question.Question,
			Option1:    question.Option1,
			Option2:    question.Option2,
			Option3:    question.Option3,
			Option4:    question.Option4,
			Score:      question.Score,
			Source:     question.Source,
			Year:       question.Year,
			Section:    question.Section,
			No:         question.No,
			Sn:         question.Sn,
			QuestionId: v.QuestionId,
			UserScore:  v.UserScore,
			UserAnswer: v.UserAnswer,
			Status:     v.Status,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"exercise": model.MyExerciseDetail{
			Id:           ex.Id,
			ExerciseId:   ex.ExerciseId,
			MyQuestions:  myQuestions,
			UserId:       ex.UserId,
			Status:       ex.Status,
			Score:        ex.Score, //总分数
			CorrectCount: ex.CorrectCount,
			WrongCount:   ex.WrongCount,
			CreatedAt:    ex.CreatedAt.UnixMilli(),
			UpdatedAt:    ex.UpdatedAt.UnixMilli(),
		},
	})
	return
}
