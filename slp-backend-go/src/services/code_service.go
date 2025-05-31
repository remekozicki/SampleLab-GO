package services

import (
	"samplelab-go/src/db"
	"samplelab-go/src/dto"
	"samplelab-go/src/models"
)

func GetAllCodes() ([]dto.CodeDto, error) {
	var codes []models.Code
	if err := db.GetDB().Find(&codes).Error; err != nil {
		return nil, err
	}
	result := make([]dto.CodeDto, len(codes))
	for i, c := range codes {
		result[i] = models.CodeToDto(c)
	}
	return result, nil
}

func SaveCode(codeDto dto.CodeDto) error {
	code := models.CodeToModel(codeDto)
	return db.GetDB().Save(&code).Error
}

func DeleteCode(id string) error {
	if err := db.GetDB().Delete(&models.Code{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
