package models

type PostInputJson struct {
	Json         string `json:"json" binding:"required,json"`
	DurationDays string `json:"duration_days" binding:"required,number"`
}
