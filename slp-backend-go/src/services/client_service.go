package services

import (
	"errors"
	"samplelab-go/src/db"
	"samplelab-go/src/dto"
	"samplelab-go/src/models"
)

func GetAllClients() ([]dto.ClientDto, error) {
	var clients []models.Client
	err := db.GetDB().Preload("Address").Find(&clients).Error
	if err != nil {
		return nil, err
	}

	var result []dto.ClientDto
	for _, c := range clients {
		result = append(result, models.ClientToDto(c))
	}

	return result, nil
}

func SaveClient(dto dto.ClientDto) error {
	dbConn := db.GetDB()

	client := models.ClientToModel(dto)

	address := client.Address

	if err := dbConn.Create(&address).Error; err != nil {
		return err
	}
	client.AddressID = address.ID
	client.Address = models.Address{}

	if err := dbConn.Create(&client).Error; err != nil {
		return err
	}
	return nil
}

func DeleteClient(id uint) error {
	conn := db.GetDB()
	if err := conn.Delete(&models.Client{}, id).Error; err != nil {
		return errors.New("nie można usunąć klienta – możliwe zależności")
	}
	return nil
}
