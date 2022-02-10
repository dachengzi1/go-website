package model

import (
	"go-website/db"
	"time"
)

type (
	QuestionGroup struct {
		Id         int       `gorm:"id" json:"id"`
		ExerciseId int `gorm:"exercise_id" json:"exerciseId"`
		QuestionId int       `gorm:"question_id" json:"questionId"`
		CreatedAt  time.Time `gorm:"created_at" json:"createdAt"`
		UpdatedAt  time.Time `gorm:"updated_at" json:"updatedAt"`
	}
)

func (qg *QuestionGroup) FindOne(id int, ch chan *QuestionGroup) (err error) {
	err = db.Db.Where("id = ?", id).First(qg).Error
	if err != nil {
		return err
	}
	ch <- qg
	return
}
