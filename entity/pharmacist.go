package entity

import "time"

type Pharmacist struct {
	ID       uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Name     string `gorm:"type:varchar(255)" json:"name"`
	Email    string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password string `gorm:"->;<-;not null" json:"-"`
	Token    string `gorm:"-" json:"token,omitempty"`

	Registration      string    `gorm:"type:varchar(100)" json:"registration"`
	PhoneNumber       string    `gorm:"type:varchar(20)" json:"phone_number"`
	FirstName         string    `gorm:"type:varchar(255)" json:"first_name"`
	LastName          string    `gorm:"type:varchar(255)" json:"last_name"`
	BankName          string    `gorm:"type:varchar(255)" json:"bank_name"`
	BankAccountNumber string    `gorm:"type:varchar(100)" json:"bank_account_number"`
	BankAccountHolder string    `gorm:"type:varchar(255)" json:"bank_account_holder"`
	ProfileImage      string    `gorm:"type:varchar(255)" json:"profile_image"`
	UserType          string    `gorm:"type:varchar(20)" json:"user_type"`
	Balance           float64   `gorm:"type:float" json:"balance"`
	PharmName         string    `gorm:"type:varchar(255)" json:"pharm_name"`
	PharmAddress      string    `gorm:"type:varchar(255)" json:"pharm_address"`
	ClaimedPoint      float64   `gorm:"type:float" json:"claimed_point"`
	UpdatedAt         time.Time `gorm:"column:updated_at" json:"updated_at"`
}
