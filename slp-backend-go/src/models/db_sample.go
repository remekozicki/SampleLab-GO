package models

import "samplelab-go/src/dto"

type Sample struct {
	ID     int64  `gorm:"primaryKey" json:"id"`
	CodeID string `json:"codeId"`
	Code   Code   `gorm:"foreignKey:CodeID"`

	ClientID int64  `json:"clientId"`
	Client   Client `gorm:"foreignKey:ClientID"`

	AssortmentID int64      `json:"assortmentId"`
	Assortment   Assortment `gorm:"foreignKey:AssortmentID"`

	AdmissionDate              string `json:"admissionDate"`
	ExpirationDate             string `json:"expirationDate"`
	ExpirationComment          string `json:"expirationComment"`
	ExaminationExpectedEndDate string `json:"examinationExpectedEndDate"`
	Size                       string `json:"size"`
	State                      string `json:"state"`
	Analysis                   bool   `json:"analysis"`

	InspectionID int64      `json:"inspectionId"`
	Inspection   Inspection `gorm:"foreignKey:InspectionID"`

	SamplingStandardID int64            `json:"samplingStandardId"`
	SamplingStandard   SamplingStandard `gorm:"foreignKey:SamplingStandardID"`

	ReportDataID int64      `json:"reportDataId"`
	ReportData   ReportData `gorm:"foreignKey:ReportDataID"`

	Examinations []Examination `gorm:"foreignKey:SampleID"`

	ProgressStatus string `json:"progressStatus"`
}

func (Sample) TableName() string {
	return "sample"
}

func ToSampleDto(s Sample) dto.SampleDto {
	return dto.SampleDto{
		ID:                         s.ID,
		CodeID:                     s.CodeID,
		ClientID:                   s.ClientID,
		AssortmentID:               s.AssortmentID,
		AdmissionDate:              s.AdmissionDate,
		ExpirationDate:             s.ExpirationDate,
		ExpirationComment:          s.ExpirationComment,
		ExaminationExpectedEndDate: s.ExaminationExpectedEndDate,
		Size:                       s.Size,
		State:                      s.State,
		Analysis:                   s.Analysis,
		InspectionID:               s.InspectionID,
		SamplingStandardID:         s.SamplingStandardID,
		ReportDataID:               s.ReportDataID,
		ProgressStatus:             s.ProgressStatus,
	}
}

func ToSampleModel(d dto.SampleDto) Sample {
	return Sample{
		ID:                         d.ID,
		CodeID:                     d.CodeID,
		ClientID:                   d.ClientID,
		AssortmentID:               d.AssortmentID,
		AdmissionDate:              d.AdmissionDate,
		ExpirationDate:             d.ExpirationDate,
		ExpirationComment:          d.ExpirationComment,
		ExaminationExpectedEndDate: d.ExaminationExpectedEndDate,
		Size:                       d.Size,
		State:                      d.State,
		Analysis:                   d.Analysis,
		InspectionID:               d.InspectionID,
		SamplingStandardID:         d.SamplingStandardID,
		ReportDataID:               d.ReportDataID,
		ProgressStatus:             d.ProgressStatus,
	}
}
