package dto

type AddressDto struct {
	ID      int64  `json:"id"`
	Street  string `json:"street"`
	ZipCode string `json:"zipCode"`
	City    string `json:"city"`
}

func (AddressDto) TableName() string {
	return "address"
}
