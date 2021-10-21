package dto

import "gorm.io/datatypes"

type OrderCreateDTO struct {
	PharmacistID uint64         `json:"pharmacist_id" form:"pharmacist_id" binding:"required"`
	DoctorID     uint64         `json:"doctor_id" form:"doctor_id" binding:"required"`
	List         datatypes.JSON `json:"list" form:"list" binding:"required"`
}
