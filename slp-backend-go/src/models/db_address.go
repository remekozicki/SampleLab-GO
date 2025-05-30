package models

import (
	"samplelab-go/src/dto"
)

type Address struct {
	ID      int64  `gorm:"primaryKey"`
	Street  string `json:"street"`
	ZipCode string `json:"zipCode"`
	City    string `json:"city"`
}

func (Address) TableName() string {
	return "address"
}

func AddressToDto(a Address) dto.AddressDto {
	return dto.AddressDto{
		ID:      a.ID,
		Street:  a.Street,
		ZipCode: a.ZipCode,
		City:    a.City,
	}
}

func AddressToModel(dto dto.AddressDto) Address {
	addr := Address{
		Street:  dto.Street,
		ZipCode: dto.ZipCode,
		City:    dto.City,
	}
	if dto.ID != 0 {
		addr.ID = dto.ID
	}
	return addr
}
