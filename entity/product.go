package entity

import "time"

type Product struct {
	ID     uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Name   string `gorm:"type:varchar(255)" json:"name"`
	Images string `gorm:"type:varchar(20000)" json:"profile_image"`

	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
