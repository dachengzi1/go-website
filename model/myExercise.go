package model

import "time"

type(
	MyExercise struct {
		Id         int       `gorm:"id" json:"id"`
		ExerciseId int `gorm:"exercise_id" json:"exerciseId"`
		UserId int       `gorm:"user_id" json:"userId"`
		CreatedAt  time.Time `gorm:"created_at" json:"createdAt"`
		UpdatedAt  time.Time `gorm:"updated_at" json:"updatedAt"`
	}
)
