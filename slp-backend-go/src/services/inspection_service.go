package services

import (
	"samplelab-go/src/db"
	"samplelab-go/src/dto"
	"samplelab-go/src/models"
)

func GetAllInspections() ([]dto.InspectionDto, error) {
	var inspections []models.Inspection
	if err := db.GetDB().Find(&inspections).Error; err != nil {
		return nil, err
	}

	var result []dto.InspectionDto
	for _, i := range inspections {
		result = append(result, dto.InspectionDto{
			ID:   i.ID,
			Name: i.Name,
		})
	}

	return result, nil
}

func SaveInspection(d dto.InspectionDto) error {
	inspection := models.Inspection{
		ID:   d.ID,
		Name: d.Name,
	}
	return db.GetDB().Create(&inspection).Error
}

func UpdateInspection(d dto.InspectionDto) error {
	return db.GetDB().Model(&models.Inspection{}).
		Where("id = ?", d.ID).
		Updates(models.Inspection{Name: d.Name}).Error
}

func DeleteInspection(id int64) error {
	return db.GetDB().Delete(&models.Inspection{}, id).Error
}
