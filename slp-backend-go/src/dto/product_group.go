package dto

type ProductGroupDto struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type ProductGroupSaveDto struct {
	Name string `json:"name"`
}
