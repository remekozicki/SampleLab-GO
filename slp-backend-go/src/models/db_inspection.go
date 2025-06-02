package models

import "samplelab-go/src/dto"

type Inspection struct {
	ID   int64  `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(255)"`
}

func (Inspection) TableName() string {
	return "inspection"
}

func InspectionToDto(i Inspection) dto.InspectionDto {
	return dto.InspectionDto{
		ID:   i.ID,
		Name: i.Name,
	}
}

func InspectionToModel(d dto.InspectionDto) Inspection {
	return Inspection{
		ID:   d.ID,
		Name: d.Name,
	}
}
