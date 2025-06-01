package dto

type SampleDto struct {
	ID                         int64  `json:"id"`
	CodeID                     string `json:"codeId"`
	ClientID                   int64  `json:"clientId"`
	AssortmentID               int64  `json:"assortmentId"`
	AdmissionDate              string `json:"admissionDate"`
	ExpirationDate             string `json:"expirationDate"`
	ExpirationComment          string `json:"expirationComment"`
	ExaminationExpectedEndDate string `json:"examinationExpectedEndDate"`
	Size                       string `json:"size"`
	State                      string `json:"state"`
	Analysis                   bool   `json:"analysis"`
	InspectionID               int64  `json:"inspectionId"`
	SamplingStandardID         int64  `json:"samplingStandardId"`
	ReportDataID               int64  `json:"reportDataId"`
	ProgressStatus             string `json:"progressStatus"`
}

type SampleFilterDto struct {
	FieldName   string       `json:"fieldName"`
	Ascending   bool         `json:"ascending"`
	PageNumber  int          `json:"pageNumber"`
	PageSize    int          `json:"pageSize"`
	Filters     FilterFields `json:"filters"`
	FuzzySearch string       `json:"fuzzySearch"`
}

type FilterFields struct {
	Codes            []string `json:"codes"`
	Clients          []string `json:"clients"`
	Groups           []string `json:"groups"`
	ProgressStatuses []string `json:"progressStatuses"` // lub []int jeśli trzymasz enum jako int
}

type SampleSummaryDto struct {
	ID             int64  `json:"id"`
	Code           string `json:"code"`
	Group          string `json:"group"`
	Assortment     string `json:"assortment"`
	ClientName     string `json:"clientName"`
	AdmissionDate  string `json:"admissionDate"`
	ProgressStatus string `json:"progressStatus"`
}
