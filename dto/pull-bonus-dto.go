package dto

type PullBonusCreateDTO struct {
	PharmacistID uint64  `json:"pharmacist_id" form:"pharmacist_id" binding:"required"`
	DoctorID     uint64  `json:"doctor_id" form:"doctor_id" binding:"required"`
	IsDoctor     uint64  `json:"is_doctor" form:"is_doctor" binding:"required"`
	Amount       float64 `json:"amount" form:"amount" binding:"required"`
}
