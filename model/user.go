package model

type User struct {
	BaseModel
	Name string `json:"name"`
	Pw   string `json:"password"`
}