package model

import (
	"time"
)

type (
	Exercise struct {
		Id          int       `gorm:"id" json:"id"`
		CreatedAt   time.Time `gorm:"created_at" json:"createdAt"`
		UpdatedAt   time.Time `gorm:"updated_at" json:"updatedAt"`
		Title       string    `gorm:"title" json:"title"`
		Description string    `gorm:"description" json:"description"`
		//Count       int       `gorm:"count,omitempty" json:"count"`
		Deleted   int        `gorm:"deleted,omitempty" json:"deleted"`
		Type      string     `gorm:"type,omitempty" json:"type"`
		Questions []Question `json:"questions" gorm:"many2many:question_group;"`
	}
	PostExercise struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Deleted     int    `json:"deleted"`
		Type        string `json:"type"`
		QuestionIds []int  `json:"questionIds"`
	}
	//ExerciseInfo struct {
	//	Id          int        ` json:"id"`
	//	CreatedAt   time.Time  ` json:"createdAt"`
	//	UpdatedAt   time.Time  `json:"updatedAt"`
	//	Title       string     ` json:"title"`
	//	Description string     `json:"description"`
	//	Deleted     int        ` json:"deleted"`
	//	Type        string     ` json:"type"`
	//	Questions   []Question `json:"questions"`
	//}
)

//func (ex *Exercise) FindOne(id int, ch chan *Exercise) (err error) {
//	err = db.Db.Where("id = ?", id).First(ex).Error
//	if err != nil {
//		return err
//	}
//	ch <- ex
//	return
//}
