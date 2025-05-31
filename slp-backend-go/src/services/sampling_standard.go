package services

import (
	"samplelab-go/src/db"
	"samplelab-go/src/dto"
	"samplelab-go/src/models"
)

func GetAllSamplingStandards() ([]dto.SamplingStandardDto, error) {
	var standards []models.SamplingStandard
	if err := db.GetDB().Find(&standards).Error; err != nil {
		return nil, err
	}

	var result []dto.SamplingStandardDto
	for _, s := range standards {
		result = append(result, dto.SamplingStandardDto{
			ID:   s.ID,
			Name: s.Name,
		})
	}

	return result, nil
}

func SaveSamplingStandard(d dto.SamplingStandardDto) error {
	ss := models.SamplingStandard{
		ID:   d.ID,
		Name: d.Name,
	}
	return db.GetDB().Create(&ss).Error
}

func UpdateSamplingStandard(d dto.SamplingStandardDto) error {
	return db.GetDB().Model(&models.SamplingStandard{}).
		Where("id = ?", d.ID).
		Updates(models.SamplingStandard{Name: d.Name}).Error
}

func DeleteSamplingStandard(id int64) error {
	return db.GetDB().Delete(&models.SamplingStandard{}, id).Error
}
