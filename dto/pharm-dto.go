package dto

type PharmUpdateDTO struct {
	ID          uint64 `json:"id" form:"id" binding:"required"`
	Name        string `json:"name" form:"name" binding:"required"`
	PhoneNumber string `json:"phone_number" form:"phone_number" binding:"required"`
	Address     string `json:"address" form:"address" binding:"required"`
	Longitude   string `json:"longitude" form:"longitude" binding:"required"`
	Latitude    string `json:"latitude" form:"latitude" binding:"required"`
	Image       string `json:"image" form:"image" binding:"required"`
}

type PharmCreateDTO struct {
	Name        string `json:"name" form:"name" binding:"required"`
	PhoneNumber string `json:"phone_number" form:"phone_number" binding:"required"`
	Address     string `json:"address" form:"address" binding:"required"`
	Longitude   string `json:"longitude" form:"longitude" binding:"required"`
	Latitude    string `json:"latitude" form:"latitude" binding:"required"`
	Image       string `json:"image" form:"image" binding:"required"`
}
