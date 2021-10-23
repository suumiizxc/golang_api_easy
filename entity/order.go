package entity

import (
	"time"

	"gorm.io/datatypes"
)

type Order struct {
	ID           uint64 `gorm:"primary_key:auto_increment" json:"id"`
	PharmacistID uint64 `gorm:"int" json:"pharmacist_id"`
	// Pharmacist   Pharmacist `gorm:"foreignkey:PharmacistID; constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"pharmacist"`
	DoctorID uint64 `gorm:"int" json:"doctor_id"`
	// Doctor       Doctor     `gorm:"foreignkey:DoctorID; constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"doctor"`
	List             datatypes.JSON `json:"list" gorm:"type:json"`
	TotalPrice       float64        `gorm:"type:float" json:"total_price"`
	Status           string         `gorm:"type:varchar(20)" json:"status"`
	CouponDoctor     float64        `gorm:"type:float" json:"coupon_doctor"`
	CouponPharmacist float64        `gorm:"type:float" json:"coupon_pharmacist"`
	UpdatedAt        time.Time      `gorm:"column:updated_at" json:"updated_at"`
}
