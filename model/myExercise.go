package model

import "time"

type (
	MyExercise struct {
		Id           int          `gorm:"id" json:"id"`
		ExerciseId   int          `gorm:"exercise_id" json:"exerciseId"`
		MyQuestions  []MyQuestion `json:"myQuestions" gorm:"many2many:my_question_group;"`
		UserId       int          `gorm:"user_id" json:"userId"`
		Status       string       `gorm:"status" json:"status"`
		Score        int          `gorm:"score" json:"score"`                //总分数
		CorrectCount int          `gorm:"correct_count" json:"correctCount"` //正确数量
		WrongCount   int          `gorm:"wrong_count" json:"wrongCount"`     //错误数量
		CreatedAt    time.Time    `gorm:"created_at" json:"createdAt"`
		UpdatedAt    time.Time    `gorm:"updated_at" json:"updatedAt"`
	}

	MyQuestion struct {
		//Id         int       `gorm:"id" json:"id"`
		////ExerciseId int       `gorm:"exercise_id" json:"exerciseId"`
		//QuestionId int       `gorm:"question_id" json:"questionId"`
		//UserId     int       `gorm:"user_id" json:"userId"`
		//UserScore  int       `gorm:"user_score" json:"userScore"`
		//UserAnswer string    `gorm:"user_answer" json:"userAnswer"`
		//CreatedAt  time.Time `gorm:"created_at" json:"createdAt"`
		//UpdatedAt  time.Time `gorm:"updated_at" json:"updatedAt"`
		Id         int       `gorm:"id" json:"id"`
		UserId     int       `gorm:"user_id" json:"userId"`
		QuestionId int       `gorm:"question_id" json:"questionId"`
		UserScore  int       `gorm:"user_score" json:"userScore"`
		UserAnswer string    `gorm:"user_answer" json:"userAnswer"`
		Status     int       `gorm:"status" json:"status"` // 正确 1｜ 错误 0
		CreatedAt  time.Time `gorm:"created_at" json:"createdAt"`
		UpdatedAt  time.Time `gorm:"updated_at" json:"updatedAt"`
	}

	MyQuestionGroup struct {
		Id           int       `gorm:"id" json:"id"`
		MyExerciseId int       `gorm:"my_exercise_id" json:"myExerciseId"`
		MyQuestionId int       `gorm:"my_question_id" json:"myQuestionId"`
		CreatedAt    time.Time `gorm:"created_at" json:"createdAt"`
		UpdatedAt    time.Time `gorm:"updated_at" json:"updatedAt"`
	}

	QuestionResult struct {
		QuestionId int    `json:"questionId"`
		UserAnswer string `json:"userAnswer"`
	}
	SubMitExercise struct {
		ExerciseId  int              `json:"exerciseId"`
		MyQuestions []QuestionResult `json:"myQuestions"`
		Status      string           `json:"status"`
	}

	MyQuestionDetail struct {
		Id         int    `json:"id"`
		Context    string `json:"context"`
		Question   string `json:"question"`
		Option1    string `json:"option1"`
		Option2    string `json:"option2"`
		Option3    string `json:"option3"`
		Option4    string `json:"option4"`
		Score      int64  `json:"score"`
		Source     string `json:"source"`
		Year       string `json:"year"`
		Section    string `json:"section"`
		No         int64  `json:"no"`
		Sn         string `json:"sn"`
		QuestionId int    ` json:"questionId"`
		UserScore  int    `json:"userScore"`
		UserAnswer string `json:"userAnswer"`
		Status     int    `json:"status"` // 正确 1｜ 错误 0
	}
	MyExerciseDetail struct {
		Id           int                `json:"id"`
		ExerciseId   int                ` json:"exerciseId"`
		MyQuestions  []MyQuestionDetail ` json:"myQuestions"`
		UserId       int                `json:"userId"`
		Status       string             `json:"status"`
		Score        int                `json:"score"`        //总分数
		CorrectCount int                `json:"correctCount"` //正确数量
		WrongCount   int                `json:"wrongCount"`   //错误数量
		CreatedAt    int64          `json:"createdAt"`
		UpdatedAt    int64          `json:"updatedAt"`
	}

)
