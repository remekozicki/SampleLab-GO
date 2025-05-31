package models

type SamplingStandard struct {
	ID   int64  `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(255)"`
}

func (SamplingStandard) TableName() string {
	return "sampling_standard"
}
