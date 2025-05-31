package services

import (
	"samplelab-go/src/db"
	"samplelab-go/src/dto"
	"samplelab-go/src/models"
)

func GetAllIndications() ([]dto.IndicationDto, error) {
	var indications []models.Indication
	if err := db.GetDB().Find(&indications).Error; err != nil {
		return nil, err
	}

	var dtos []dto.IndicationDto
	for _, i := range indications {
		dtos = append(dtos, models.ToIndicationDto(i))
	}
	return dtos, nil
}

func GetIndicationByID(id int64) (dto.IndicationDto, error) {
	var indication models.Indication
	if err := db.GetDB().First(&indication, id).Error; err != nil {
		return dto.IndicationDto{}, err
	}
	return models.ToIndicationDto(indication), nil
}

func SaveIndication(input dto.IndicationDto) error {
	model := models.ToIndicationModel(input)
	return db.GetDB().Save(&model).Error
}

func DeleteIndication(id int64) error {
	return db.GetDB().Delete(&models.Indication{}, id).Error
}

//func SelectIndicationsForSample(sampleID int64) ([]dto.IndicationDto, error) {
//	var sample models.Sample
//	if err := db.GetDB().Preload("Assortment.Indications").First(&sample, sampleID).Error; err != nil {
//		return nil, errors.New("próbka nie znaleziona")
//	}
//
//	if sample.Assortment.ID == 0 {
//		return nil, errors.New("próbka nie ma przypisanego asortymentu")
//	}
//
//	indications := sample.Assortment.Indications
//	var result []dto.IndicationDto
//	for _, ind := range indications {
//		result = append(result, mappers.IndicationToDto(ind))
//	}
//
//	return result, nil
//}
