package dto

type ProductUpdateDTO struct {
	ID          uint64 `json:"id" form:"id" binding:"required"`
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserID      uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}

type ProductCreateDTO struct {
	Title        string `json:"title" form:"title" binding:"required"`
	Description  string `json:"description" form:"description" binding:"required"`
	ChemicalName string `json:"chemical_name" form:"chemical_name" binding:"required"`
	Image        string `json:"image" form:"image" binding:"required"`
	StoreCount   int    `json:"store_count" form:"store_count" binding:"required"`
	UserID       uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}
