package dto

type RegisterDTO struct {
	Name     string `json:"name" form:"name" biding:"required" validate:"min:1"`
	Email    string `json:"email" form:"email" binding:"required,email" validate:"email"`
	Password string `json:"password" form:"password" binding:"required" validate:"min:6"`
}
