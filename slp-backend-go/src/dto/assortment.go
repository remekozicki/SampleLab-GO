package dto

type AssortmentDto struct {
	ID                 int64  `json:"id"`
	Name               string `json:"name"`
	OrganolepticMethod string `json:"organolepticMethod"`
	GroupID            int64  `json:"groupId"`
}
