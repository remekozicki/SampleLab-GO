package models

import "samplelab-go/src/dto"

type SamplingStandard struct {
	ID   int64  `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(255)"`
}

func (SamplingStandard) TableName() string {
	return "sampling_standard"
}

func SamplingStandardToModel(d dto.SamplingStandardDto) SamplingStandard {
	return SamplingStandard{
		ID:   d.ID,
		Name: d.Name,
	}
}

func SamplingStandardToDto(m SamplingStandard) dto.SamplingStandardDto {
	return dto.SamplingStandardDto{
		ID:   m.ID,
		Name: m.Name,
	}
}
