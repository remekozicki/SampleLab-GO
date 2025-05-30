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

	// 1. Zamień DTO na Model
	client := models.ClientToModel(dto)

	address := client.Address

	// 2. Najpierw zapisz address (bez ID == fail)
	if err := dbConn.Create(&address).Error; err != nil {
		return err
	}

	// 3. Przypisz wygenerowane ID adresu
	client.AddressID = client.Address.ID

	// 4. Zapisz klienta
	return dbConn.Create(&client).Error
}

func DeleteClient(id uint) error {
	conn := db.GetDB()
	if err := conn.Delete(&models.Client{}, id).Error; err != nil {
		return errors.New("nie można usunąć klienta – możliwe zależności")
	}
	return nil
}
