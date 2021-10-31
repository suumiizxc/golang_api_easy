package entity

import "time"

type PullBonus struct {
	ID           uint64    `gorm:"primary_key:auto_increment" json:"id"`
	PharmacistID uint64    `gorm:"type:int" json:"pharmacist_id"`
	DoctorID     uint64    `gorm:"type:int" json:"doctor_id"`
	IsDoctor     string    `gorm:"type:varchar(20)" json:"is_doctor"`
	Amount       float64   `gorm:"type:float" json:"amount"`
	Status       string    `gorm:"type:varchar(50)" json:"status"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
}
