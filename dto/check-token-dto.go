package dto

type CheckTokenDTO struct {
	Token string `json:"token" form:"token" binding:"required"`
}
