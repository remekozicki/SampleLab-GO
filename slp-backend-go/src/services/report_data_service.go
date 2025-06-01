package services

import (
	"samplelab-go/src/db"
	"samplelab-go/src/dto"
	"samplelab-go/src/models"
)

func GetAllReportData() ([]dto.ReportDataDto, error) {
	var data []models.ReportData
	err := db.GetDB().Find(&data).Error
	if err != nil {
		return nil, err
	}

	var result []dto.ReportDataDto
	for _, d := range data {
		result = append(result, models.ToReportDataDto(d))
	}
	return result, nil
}

func GetReportDataBySampleID(sampleID uint) (*dto.ReportDataDto, error) {
	var data models.ReportData
	err := db.GetDB().Where("sample_id = ?", sampleID).First(&data).Error
	if err != nil {
		return nil, err
	}
	dtoData := models.ToReportDataDto(data)
	return &dtoData, nil
}

func SaveReportData(d dto.ReportDataDto) error {
	model := models.ToReportDataModel(d)
	return db.GetDB().Save(&model).Error
}

func DeleteReportData(id uint) error {
	return db.GetDB().Delete(&models.ReportData{}, id).Error
}
