package entity

import "time"

type Product struct {
	ID           uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Title        string    `gorm:"type:varchar(255)" json:"title"`
	Description  string    `gorm:"type:text" json:"description"`
	ChemicalName string    `gorm:"type:text" json:"chemical_name"`
	Image        string    `gorm:"type:varchar(255)" json:"image"`
	StoreCount   int       `gorm:"type:int" json:"store_count"`
	UserID       uint64    `gorm:"not null" json:"-"`
	User         User      `gorm:"foreignkey:UserID; constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
}
