package services

import (
	"samplelab-go/src/db"
	"samplelab-go/src/dto"
	"samplelab-go/src/models"
)

func GetAllAssortments() []dto.AssortmentDto {
	var assortments []models.Assortment
	db.GetDB().Find(&assortments)

	var result []dto.AssortmentDto
	for _, a := range assortments {
		result = append(result, models.AssortmentToDto(a))
	}
	if result == nil {
		return []dto.AssortmentDto{}
	}
	return result
}

func SaveAssortment(dto dto.AssortmentDto) error {
	model := models.AssortmentFromDto(dto)
	return db.GetDB().Create(&model).Error
}

func UpdateAssortment(dto dto.AssortmentDto) error {
	model := models.AssortmentFromDto(dto)
	return db.GetDB().Save(&model).Error
}

func DeleteAssortmentByID(id string) error {
	return db.GetDB().Delete(&models.Assortment{}, id).Error
}
