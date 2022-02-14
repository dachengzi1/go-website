package model

import "time"

type User struct {
	Id        int      `gorm:"column:id" json:"id"`
	Mobile    string    `gorm:"mobile" json:"mobile"`
	NickName  string    `gorm:"nick_name" json:"nickName"`
	Password  string    `gorm: "password" json:"password"`
	CreatedAt time.Time `gorm: "created_at" json:"created_at"`
	UpdatedAt time.Time `gorm: "updated_at" json:"updated_at" `
}
