package models

type Inspection struct {
	ID   int64  `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(255)"`
}

func (Inspection) TableName() string {
	return "inspection"
}
