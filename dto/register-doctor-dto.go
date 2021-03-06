package dto

type RegisterDoctorDTO struct {
	Name              string `json:"name" form:"name" biding:"required" validate:"min:1"`
	Email             string `json:"email" form:"email" binding:"required,email" validate:"email"`
	Password          string `json:"password" form:"password" binding:"required" validate:"min:6"`
	Registration      string `json:"registration" form:"registration" binding:"required" validate:"registration"`
	PhoneNumber       string `json:"phone_number" form:"phone_number" binding:"required" validate:"phone_number"`
	FirstName         string `json:"first_name" form:"first_name" binding:"required" validate:"first_name"`
	LastName          string `json:"last_name" form:"last_name" binding:"required" validate:"last_name"`
	BankName          string `json:"bank_name" form:"bank_name" binding:"required" validate:"bank_name"`
	BankAccountNumber string `json:"bank_account_number" form:"bank_account_name" binding:"required" validate:"bank_account_number"`
	BankAccountHolder string `json:"bank_account_holder" form:"bank_account_holder" binding:"required" validate:"bank_account_holder"`
	HospitalName      string `json:"hospital_name" form:"hospital_name" binding:"required" validate:"hospital_name"`
	DoctorAddress     string `json:"doctor_address" form:"doctor_address" binding:"required" validate:"doctor_address"`
	DoctorGuarentee   string `json:"doctor_guarentee" form:"doctor_guarentee" binding:"required" validate:"doctor_guarentee"`
}
