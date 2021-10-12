package entity

import "time"

type Pharm struct {
	ID        uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Name      string    `gorm:"type:varchar(255)" json:"name"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
