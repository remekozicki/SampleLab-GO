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
	AssortmentID   int64
}

func (Indication) TableName() string {
	return "indication"
}

func IndicationToDto(i Indication) dto.IndicationDto {
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

func MapIndicationsToDto(list []Indication) []dto.IndicationDto {
	result := make([]dto.IndicationDto, len(list))
	for i, item := range list {
		result[i] = IndicationToDto(item)
	}
	return result
}

func MapIndicationsFromDto(list []dto.IndicationDto) []Indication {
	result := make([]Indication, len(list))
	for i, item := range list {
		result[i] = ToIndicationModel(item)
	}
	return result
}
