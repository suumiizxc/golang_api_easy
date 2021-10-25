package dto

type ChallengeCreateDTO struct {
	Title       string  `json:"title" form:"title" binding:"required"`
	Description string  `json:"description" form:"description" binding:"required"`
	Goal        float64 `json:"goal" form:"goal" binding:"required"`
	StartDate   string  `json:"start_date" form:"start_date" binding:"required"`
	EndDate     string  `json:"end_date" form:"end_date" binding:"required"`
}
