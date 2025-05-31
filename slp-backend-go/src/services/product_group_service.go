package services

import (
	"samplelab-go/src/db"
	"samplelab-go/src/dto"
	"samplelab-go/src/models"
)

func GetAllProductGroups() ([]dto.ProductGroupDto, error) {
	var groups []models.ProductGroup
	if err := db.GetDB().Find(&groups).Error; err != nil {
		return nil, err
	}

	var result []dto.ProductGroupDto
	for _, g := range groups {
		result = append(result, models.ProductGroupToDto(g))
	}
	return result, nil
}

func SaveProductGroup(d dto.ProductGroupSaveDto) error {
	pg := models.ProductGroupFromSaveDto(d)
	return db.GetDB().Create(&pg).Error
}

func UpdateProductGroup(id int64, d dto.ProductGroupSaveDto) error {
	return db.GetDB().Model(&models.ProductGroup{}).Where("id = ?", id).Updates(models.ProductGroupFromSaveDto(d)).Error
}

func DeleteProductGroup(id int64) error {
	return db.GetDB().Delete(&models.ProductGroup{}, id).Error
}
