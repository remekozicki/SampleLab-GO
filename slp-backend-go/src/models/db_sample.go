package models

import (
	"samplelab-go/src/dto"
	"samplelab-go/src/enum"
)

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

	ProgressStatus enum.ProgressStatus `json:"progressStatus"`
}

func (Sample) TableName() string {
	return "sample"
}

func ToSampleDto(s Sample) dto.SampleDto {
	return dto.SampleDto{
		ID:                         s.ID,
		Code:                       CodeToDto(s.Code),
		Client:                     ClientToDto(s.Client),
		Assortment:                 AssortmentToDto(s.Assortment),
		AdmissionDate:              s.AdmissionDate,
		ExpirationDate:             s.ExpirationDate,
		ExpirationComment:          s.ExpirationComment,
		ExaminationExpectedEndDate: s.ExaminationExpectedEndDate,
		Size:                       s.Size,
		State:                      s.State,
		Analysis:                   s.Analysis,
		Inspection:                 InspectionToDto(s.Inspection),
		SamplingStandard:           SamplingStandardToDto(s.SamplingStandard),
		ReportData:                 ToReportDataDto(s.ReportData),
		ProgressStatus:             s.ProgressStatus,
	}
}

func ToSampleModel(d dto.SampleDto) Sample {
	expirationComment := d.ExpirationComment
	if expirationComment == "" {
		expirationComment = "Brak"
	}

	state := d.State
	if state == "" {
		state = "Bez zastrzeżeń"
	}

	return Sample{
		ID:                         d.ID,
		CodeID:                     d.Code.ID,
		ClientID:                   d.Client.ID,
		AssortmentID:               d.Assortment.ID,
		AdmissionDate:              d.AdmissionDate,
		ExpirationDate:             d.ExpirationDate,
		ExpirationComment:          expirationComment,
		ExaminationExpectedEndDate: d.ExaminationExpectedEndDate,
		Size:                       d.Size,
		State:                      state,
		Analysis:                   d.Analysis,
		InspectionID:               d.Inspection.ID,
		SamplingStandardID:         d.SamplingStandard.ID,
		ReportDataID:               d.ReportData.ID,
		ProgressStatus:             d.ProgressStatus,
	}
}
