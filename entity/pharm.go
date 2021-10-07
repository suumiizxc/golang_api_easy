package entity

type Pharm struct {
	ID          uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Name        string `gorm:"type:varchar(255)" json:"name"`
	PhoneNumber string `gorm:"type:varchar(20)" json:"phone_number"`
	Address     string `gorm:"type:varchar(255)" json:"address"`
	Longitude   string `gorm:"type:varchar(100)" json:"longitude"`
	Latitude    string `gorm:"type:varchar(100)" json:"latitude"`
	Image       string `gorm:"type:varchar(255)" json:"image"`
}
