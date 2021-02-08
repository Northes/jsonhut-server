package models

type PostInputJson struct {
	Json string `json:"json" binding:"required,json"`
	Time string `json:"time" binding:"required,number"`
}
