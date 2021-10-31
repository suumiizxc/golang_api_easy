package dto

type PharmacistUpdateDTO struct {
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
	BankAccountHolder string `json:"bank_account_holder" form:"bank_account_holder" binding:"required" validate:"bank_account_holder"`
	ProfileImage      string `json:"profile_image" form:"image" binding:"required" validate:"profile_image"`
	PharmName         string `json:"pharm_name" form:"form_name" binding:"required" validate:"pharm_name"`
	PharmAddress      string `json:"pharm_address" form:"pharm_address" binding:"required" validate:"pharm_address"`
	// UpdatedAt         time.Time `json:"updated_at" form:"updated_at" binding:"required" validate:"updated_at"`
}
