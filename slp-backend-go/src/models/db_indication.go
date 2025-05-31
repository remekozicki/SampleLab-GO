package models

import (
	"samplelab-go/src/dto"
)

type Indication struct {
	ID             int64 `gorm:"primaryKey"`
	Name           string
	Method         string
	Unit           string
	Laboratory     string
	IsOrganoleptic bool `json:"isOrganoleptic"`
}

func (Indication) TableName() string {
	return "indication"
}

func ToIndicationDto(i Indication) dto.IndicationDto {
	return dto.IndicationDto{
		ID:             i.ID,
		Name:           i.Name,
		Method:         i.Method,
		Unit:           i.Unit,
		Laboratory:     i.Laboratory,
		IsOrganoleptic: i.IsOrganoleptic,
	}
}

func ToIndicationModel(d dto.IndicationDto) Indication {
	return Indication{
		ID:             d.ID,
		Name:           d.Name,
		Method:         d.Method,
		Unit:           d.Unit,
		Laboratory:     d.Laboratory,
		IsOrganoleptic: d.IsOrganoleptic,
	}
}
