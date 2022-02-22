package model

import "time"

type (
	MyExercise struct {
		Id           int          `gorm:"id" json:"id"`
		ExerciseId   int          `gorm:"exercise_id" json:"exerciseId"`
		MyQuestions  []MyQuestion `json:"myQuestions" gorm:"many2many:my_question_group;"`
		UserId       int          `gorm:"user_id" json:"userId"`
		Status       string          `gorm:"status" json:"status"`
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
		Status int    `gorm:"status" json:"status"`// 正确 1｜ 错误 0
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
		QuestionId   int              `json:"questionId"`
		UserAnswer   string              `json:"userAnswer"`
	}
	SubMitExercise struct {
		ExerciseId   int              `json:"exerciseId"`
		MyQuestions  []QuestionResult `json:"myQuestions"`
		Status       string              `json:"status"`
	}

	//QuestionResult struct {
	//	QuestionId int       `json:"questionId"`
	//	UserAnswer string    `json:"userAnswer"`
	//	CreatedAt  time.Time `gorm:"created_at" json:"createdAt"`
	//	UpdatedAt  time.Time `gorm:"updated_at" json:"updatedAt"`
	//}


)
