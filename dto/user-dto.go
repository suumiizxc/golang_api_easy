package dto

type UserUpdateDTO struct {
	ID                uint64 `json:"id" form:"id" `
	Name              string `json:"name" form:"name" binding:"required"`
	Email             string `json:"email" form:"email" binding:"required" validate:"email"`
	Password          string `json:"password,omitempty" form:"password,omitempty" validate:"min:8"`
	Registration      string `json:"registration" form:"registration" binding:"required" validate:"registration"`
	PhoneNumber       string `json:"phone_number" form:"phone_number" binding:"required" validate:"phone_number"`
	FirstName         string `json:"first_name" form:"first_name" binding:"required" validate:"first_name"`
	LastName          string `json:"last_name" form:"last_name" binding:"required" validate:"last_name"`
	BankName          string `json:"bank_name" form:"bank_name" binding:"required" validate:"bank_name"`
	BankAccountNumber string `json:"bank_account_number" form:"bank_account_name" binding:"required" validate:"bank_account_number"`
	UserType          string `json:"user_type" form:"user_type" binding:"require" validate:"user_type"`
	// UpdatedAt         time.Time `json:"updated_at" form:"updated_at" binding:"required" validate:"UpdatedAt"`
}

// type UserCreateDTO struct {
// 	Name     string `json:"name" form:"name" binding:"required"`
// 	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
// 	Password string `json:"password,omitempty" form:"password,omitempty" validate:"min:6" binding:"required"`
// }

// Registration      string `gorm:"type:varchar(100)" json:"registration"`
// PhoneNumber       string `gorm:"type:varchar(20)" json:"phone_number"`
// FirstName         string `gorm:"type:varchar(255)" json:"first_name"`
// LastName          string `gorm:"type:varchar(255)" json:"last_name"`
// BankName          string `gorm:"type:varchar(255)" json:"bank_name"`
// BankAccountNumber string `gorm:"type:varchar(100)" json:"bank_account_number"`
