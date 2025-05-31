package models

type ProductGroupSamplingStandard struct {
	GroupID            int64 `gorm:"column:groups_id;not null"`
	SamplingStandardID int64 `gorm:"column:sampling_standards_id;not null"`
}

func (ProductGroupSamplingStandard) TableName() string {
	return "product_group_sampling_standards"
}
