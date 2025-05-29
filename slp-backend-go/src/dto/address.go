package dto

type Address struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Street  string `json:"street"`
	ZipCode string `json:"zipCode"`
	City    string `json:"city"`
}

func (Address) TableName() string {
	return "address"
}
