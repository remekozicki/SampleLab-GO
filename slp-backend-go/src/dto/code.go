package dto

type CodeDto struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
