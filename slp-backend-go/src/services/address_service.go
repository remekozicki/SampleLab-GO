package services

import (
	"samplelab-go/src/db"
	"samplelab-go/src/dto"
)

func GetAllAddresses() ([]dto.Address, error) {
	conn := db.GetDB()
	var address []dto.Address
	result := conn.Find(&address)
	return address, result.Error
}
