package model

import "time"

type BaseModel struct {
	Id        uint      `json:"id" gorm:"column:id"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}
