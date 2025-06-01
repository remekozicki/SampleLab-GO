package dto

type ExaminationDto struct {
	ID               int64   `json:"id"`
	IndicationID     int64   `json:"indicationId"`
	SampleID         int64   `json:"sampleId"`
	Signage          string  `json:"signage"`
	NutritionalValue string  `json:"nutritionalValue"`
	Specification    string  `json:"specification"`
	Regulation       string  `json:"regulation"`
	SamplesNumber    int     `json:"samplesNumber"`
	Result           string  `json:"result"`
	StartDate        string  `json:"startDate"`
	EndDate          string  `json:"endDate"`
	MethodStatus     string  `json:"methodStatus"`
	Uncertainty      float64 `json:"uncertainty"`
	LOD              float64 `json:"lod"`
	LOQ              float64 `json:"loq"`
}
