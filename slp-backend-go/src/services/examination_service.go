package services

import (
	"samplelab-go/src/db"
	"samplelab-go/src/dto"
	"samplelab-go/src/models"
)

func GetAllExaminationsForSample(sampleID int64) ([]dto.ExaminationDto, error) {
	var list []models.Examination
	err := db.GetDB().Where("sample_id = ?", sampleID).Find(&list).Error
	if err != nil {
		return nil, err
	}

	var result []dto.ExaminationDto
	for _, e := range list {
		result = append(result, models.ToExaminationDto(e))
	}
	return result, nil
}

func GetExaminationByID(id int64) (*dto.ExaminationDto, error) {
	var model models.Examination
	err := db.GetDB().First(&model, id).Error
	if err != nil {
		return nil, err
	}
	dto := models.ToExaminationDto(model)
	return &dto, nil
}

//func SaveExamination(d dto.ExaminationDto) error {
//	model := models.ToExaminationModel(d)
//	if err := db.GetDB().Save(&model).Error; err != nil {
//		return err
//	}
//	return updateSampleProgress(model.SampleID)
//}

func DeleteExamination(id int64) error {
	return db.GetDB().Delete(&models.Examination{}, id).Error
}

//func updateSampleProgress(sampleID uint) error {
//	var examinations []models.Examination
//	db := db.GetDB()
//
//	err := db.Where("sample_id = ?", sampleID).Find(&examinations).Error
//	if err != nil {
//		return err
//	}
//
//	completed := false
//	for _, e := range examinations {
//		if e.Result != "" {
//			completed = true
//			break
//		}
//	}
//
//	status := "IN_PROGRESS"
//	if completed {
//		status = "DONE"
//	}
//
//	return db.Model(&models.Sample{}).Where("id = ?", sampleID).Update("progress_status", status).Error
//}
