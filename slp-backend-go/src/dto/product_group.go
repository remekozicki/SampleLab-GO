package dto

type ProductGroupDto struct {
	ID                int64                 `json:"id"`
	Name              string                `json:"name"`
	SamplingStandards []SamplingStandardDto `json:"samplingStandards"`
	Assortments       []AssortmentDto       `json:"assortments"`
}

type ProductGroupSaveDto struct {
	Name string `json:"name"`
}
