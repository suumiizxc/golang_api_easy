package entity

type Challenge struct {
	ID          uint64  `gorm:"primary_key:auto_increment" json:"id"`
	Title       string  `gorm:"type:varchar(255)" json:"title"`
	Description string  `gorm:"type:varchar(1000)" json:"description"`
	Goal        float64 `gorm:"type:float" json:"goal"`
	StartDate   string  `gorm:"type:varchar(50)" json:"start_date"`
	EndDate     string  `gorm:"type:varchar(50)" json:"end_date"`
}
