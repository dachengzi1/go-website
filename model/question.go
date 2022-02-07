package model

import (
	"goschool/db"
	"time"
)

type (
	Question struct {
		Id        int `gorm:"id" json:"id"`
		CreatedAt time.Time          `gorm:"created_at" json:"createdAt"`
		UpdatedAt time.Time          `gorm:"updated_at" json:"updatedAt"`
		Context string `gorm:"context" json:"context"`
		Question string `gorm:"question" json:"question"`
		Option1 string `gorm:"option1,omitempty" json:"option1"`
		Option2 string `gorm:"option2,omitempty" json:"option2"`
		Option3 string `gorm:"option3,omitempty" json:"option3"`
		Option4 string `gorm:"option4,omitempty" json:"option4"`
		Score   int64  `gorm:"score,omitempty" json:"score"`
		Source  string `gorm:"source,omitempty" json:"source"`
		Year    string `gorm:"year,omitempty" json:"year"`
		Section string `gorm:"section,omitempty" json:"section"`
		No      int64  `gorm:"no,omitempty" json:"no"`
		Sn      string `gorm:"sn,omitempty" json:"sn"`
		Order   int64  `gorm:"order,omitempty" json:"order"`
		PayFree string `gorm:"pay_free,omitempty" json:"payFree"`
		Deleted int   `gorm:"deleted,omitempty" json:"deleted"`
		Details string `gorm:"details,omitempty" json:"details"`
		Answer  string `gorm:"answer,omitempty" json:"answer"`
	}
)

func (q *Question) Insert() (err error) {
	result := db.Db.Create(&q) // 新插入数据的id
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}
