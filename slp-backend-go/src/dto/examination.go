package dto

type ExaminationDto struct {
	ID               uint    `json:"id"`
	IndicationID     uint    `json:"indicationId"`
	SampleID         uint    `json:"sampleId"`
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
