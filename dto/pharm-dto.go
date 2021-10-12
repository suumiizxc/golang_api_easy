package dto

type PharmCreateDTO struct {
	Name string `json:"name" form:"name" binding:"required"`
}
