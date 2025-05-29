package models

type AddressDTO struct {
	ID      uint   `json:"id"`
	Street  string `json:"street"`
	ZipCode string `json:"zipCode"`
	City    string `json:"city"`
}
