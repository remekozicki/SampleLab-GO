package dto

import "samplelab-go/src/enum"

type SampleDto struct {
	ID int64 `json:"id"`
	//CodeID                     string               `json:"codeId"`
	Code CodeDto `json:"code,omitempty"`
	//ClientID                   int64                `json:"clientId"`
	Client ClientDto `json:"client,omitempty"`
	//AssortmentID               int64                `json:"assortmentId"`
	Assortment                 AssortmentDto `json:"assortment,omitempty"`
	AdmissionDate              string        `json:"admissionDate"`
	ExpirationDate             string        `json:"expirationDate"`
	ExpirationComment          string        `json:"expirationComment"`
	ExaminationExpectedEndDate string        `json:"examinationExpectedEndDate"`
	Size                       string        `json:"size"`
	State                      string        `json:"state"`
	Analysis                   bool          `json:"analysis"`
	//InspectionID               int64                `json:"inspectionId"`
	Inspection InspectionDto `json:"inspection,omitempty"`
	//SamplingStandardID         int64                `json:"samplingStandardId"`
	SamplingStandard SamplingStandardDto `json:"samplingStandard,omitempty"`
	//ReportDataID               int64                `json:"reportDataId"`
	ReportData     ReportDataDto       `json:"reportData,omitempty"`
	ProgressStatus enum.ProgressStatus `json:"progressStatus"`
}

type SampleFilterDto struct {
	FieldName   string        `json:"fieldName"`
	Ascending   bool          `json:"ascending"`
	PageNumber  int           `json:"pageNumber"`
	PageSize    int           `json:"pageSize"`
	Filters     *FilterFields `json:"filters"`
	FuzzySearch string        `json:"fuzzySearch"`
}

type FilterFields struct {
	Code             []string              `json:"codes"`
	Client           []string              `json:"clients"`
	Groups           []string              `json:"groups"`
	ProgressStatuses []enum.ProgressStatus `json:"progressStatuses"` // lub []int jeśli trzymasz enum jako int
}

type SampleSummaryDto struct {
	ID             int64               `json:"id"`
	Code           string              `json:"code"`
	Group          string              `json:"group"`
	Assortment     string              `json:"assortment"`
	ClientName     string              `json:"clientName"`
	AdmissionDate  string              `json:"admissionDate"`
	ProgressStatus enum.ProgressStatus `json:"progressStatus"`
}
