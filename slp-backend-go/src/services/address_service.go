package services

import (
	"samplelab-go/src/db"
	"samplelab-go/src/dto"
)

func GetAllAddresses() ([]dto.AddressDto, error) {
	conn := db.GetDB()
	var address []dto.AddressDto
	result := conn.Find(&address)
	return address, result.Error
}
