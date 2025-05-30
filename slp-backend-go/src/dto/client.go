package dto

type ClientDto struct {
	ID          int64      `json:"id"`
	WijharsCode string     `json:"wijharsCode" binding:"required"`
	Name        string     `json:"name" binding:"required"`
	Address     AddressDto `json:"address" binding:"required"`
}
