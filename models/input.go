package models

type PostInputJson struct {
	Json string `json:"json" binding:"required,json"`
	Day string `json:"day" binding:"required,number"`
}
