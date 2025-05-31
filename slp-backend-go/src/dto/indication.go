package dto

type IndicationDto struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	Method         string `json:"method"`
	Unit           string `json:"unit"`
	Laboratory     string `json:"laboratory"`
	IsOrganoleptic bool   `json:"isOrganoleptic"`
}
