package models

import "samplelab-go/src/dto"

type Examination struct {
	ID               int64 `gorm:"primaryKey"`
	IndicationID     int64
	SampleID         int64
	Signage          string
	NutritionalValue string
	Specification    string
	Regulation       string
	SamplesNumber    int
	Result           string
	StartDate        string
	EndDate          string
	MethodStatus     string
	Uncertainty      float64
	LOD              float64
	LOQ              float64
}

func (Examination) TableName() string {
	return "examination"
}

func ToExaminationDto(e Examination) dto.ExaminationDto {
	return dto.ExaminationDto{
		ID:               e.ID,
		IndicationID:     e.IndicationID,
		SampleID:         e.SampleID,
		Signage:          e.Signage,
		NutritionalValue: e.NutritionalValue,
		Specification:    e.Specification,
		Regulation:       e.Regulation,
		SamplesNumber:    e.SamplesNumber,
		Result:           e.Result,
		StartDate:        e.StartDate,
		EndDate:          e.EndDate,
		MethodStatus:     e.MethodStatus,
		Uncertainty:      e.Uncertainty,
		LOD:              e.LOD,
		LOQ:              e.LOQ,
	}
}

func ToExaminationModel(d dto.ExaminationDto) Examination {
	return Examination{
		ID:               d.ID,
		IndicationID:     d.IndicationID,
		SampleID:         d.SampleID,
		Signage:          d.Signage,
		NutritionalValue: d.NutritionalValue,
		Specification:    d.Specification,
		Regulation:       d.Regulation,
		SamplesNumber:    d.SamplesNumber,
		Result:           d.Result,
		StartDate:        d.StartDate,
		EndDate:          d.EndDate,
		MethodStatus:     d.MethodStatus,
		Uncertainty:      d.Uncertainty,
		LOD:              d.LOD,
		LOQ:              d.LOQ,
	}
}
