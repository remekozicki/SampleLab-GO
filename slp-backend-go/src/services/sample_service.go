package services

import (
	"errors"
	"gorm.io/gorm"
	"samplelab-go/src/db"
	"samplelab-go/src/dto"
	"samplelab-go/src/models"
	"strings"
)

func GetAllSamples() ([]dto.SampleDto, error) {
	var samples []models.Sample
	err := db.GetDB().
		Preload("Code").
		Preload("Client").Preload("Client.Address").
		Preload("Assortment").Preload("Assortment.Group").
		Preload("Inspection").
		Preload("SamplingStandard").
		Preload("ReportData").
		Preload("ReportData.ManufacturerAddress").
		Preload("ReportData.SupplierAddress").
		Preload("ReportData.SellerAddress").
		Preload("ReportData.RecipientAddress").
		Preload("Examinations").
		Find(&samples).Error
	if err != nil {
		return nil, err
	}

	var result []dto.SampleDto
	for _, s := range samples {
		result = append(result, models.ToSampleDto(s))
	}

	return result, nil
}

func GetSampleByID(id int64) (dto.SampleDto, error) {
	var sample models.Sample
	err := db.GetDB().
		Preload("Code").
		Preload("Client").Preload("Client.Address").
		Preload("Assortment").Preload("Assortment.Group").
		Preload("Inspection").
		Preload("SamplingStandard").
		Preload("ReportData").
		Preload("ReportData.ManufacturerAddress").
		Preload("ReportData.SupplierAddress").
		Preload("ReportData.SellerAddress").
		Preload("ReportData.RecipientAddress").
		Preload("Examinations").
		First(&sample, id).Error
	if err != nil {
		return dto.SampleDto{}, err
	}

	return models.ToSampleDto(sample), nil
}

func SaveSample(sampleDto dto.SampleDto) error {
	dbConn := db.GetDB()

	sample := models.ToSampleModel(sampleDto)

	// Zapisywanie zagnieżdżonych obiektów (jeśli potrzebne)
	if sample.ReportData.ID == 0 {
		report := sample.ReportData

		// Zapisz adresy raportu (po kolei)
		dbConn.Create(&report.ManufacturerAddress)
		dbConn.Create(&report.SupplierAddress)
		dbConn.Create(&report.SellerAddress)
		dbConn.Create(&report.RecipientAddress)

		// Zapisz cały raport z powiązaniami
		sample.ReportData = report
	}

	// Zapis próbki
	if err := dbConn.Create(&sample).Error; err != nil {
		return err
	}

	return nil
}

func DeleteSample(id int64) error {
	conn := db.GetDB()
	if err := conn.Delete(&models.Sample{}, id).Error; err != nil {
		return errors.New("nie można usunąć próbki – możliwe zależności")
	}
	return nil
}

func UpdateSample(sampleDto dto.SampleDto) error {
	dbConn := db.GetDB()

	// Pobierz istniejącą próbkę
	var existing models.Sample
	if err := dbConn.First(&existing, sampleDto.ID).Error; err != nil {
		return errors.New("próbka nie istnieje")
	}

	// Zamiana DTO na model – nowa wersja
	updated := models.ToSampleModel(sampleDto)
	updated.ID = existing.ID // upewnij się, że ID zostaje to samo

	// Aktualizacja relacji jeśli są zagnieżdżone
	if updated.ReportData.ID != 0 {
		report := updated.ReportData

		dbConn.Save(&report.ManufacturerAddress)
		dbConn.Save(&report.SupplierAddress)
		dbConn.Save(&report.SellerAddress)
		dbConn.Save(&report.RecipientAddress)

		dbConn.Save(&report)
	}

	// Aktualizacja głównego obiektu Sample
	if err := dbConn.Session(&gorm.Session{FullSaveAssociations: true}).Save(&updated).Error; err != nil {
		return err
	}

	return nil
}

func FilterSamples(filter dto.SampleFilterDto) ([]dto.SampleSummaryDto, int64, error) {
	dbConn := db.GetDB()
	query := dbConn.Model(&models.Sample{}).
		Joins("JOIN codes ON samples.code_id = codes.id").
		Joins("JOIN clients ON samples.client_id = clients.id").
		Joins("JOIN assortments ON samples.assortment_id = assortments.id").
		Joins("JOIN product_groups ON assortments.group_id = product_groups.id").
		Joins("LEFT JOIN examinations ON examinations.sample_id = samples.id").
		Joins("LEFT JOIN indications ON examinations.indication_id = indications.id")

	// Fuzzy search
	if filter.FuzzySearch != "" {
		pattern := "%%" + strings.ToLower(filter.FuzzySearch) + "%%"
		query = query.Where(
			dbConn.Where("LOWER(codes.id::text) LIKE ?", pattern).
				Or("LOWER(clients.name) LIKE ?", pattern).
				Or("LOWER(assortments.name) LIKE ?", pattern).
				Or("LOWER(product_groups.name) LIKE ?", pattern).
				Or("LOWER(indications.name) LIKE ?", pattern).
				Or("LOWER(indications.method) LIKE ?", pattern),
		)
	}

	// Filtry po polach
	if len(filter.Filters.Codes) > 0 {
		query = query.Where("codes.id IN ?", filter.Filters.Codes)
	}
	if len(filter.Filters.Clients) > 0 {
		query = query.Where("clients.name IN ?", filter.Filters.Clients)
	}
	if len(filter.Filters.Groups) > 0 {
		query = query.Where("product_groups.name IN ?", filter.Filters.Groups)
	}
	if len(filter.Filters.ProgressStatuses) > 0 {
		query = query.Where("samples.progress_status IN ?", filter.Filters.ProgressStatuses)
	}

	// Zlicz rekordy do paginacji
	var total int64
	query.Count(&total)

	// Sortowanie
	sortField := filter.FieldName
	if sortField == "" {
		sortField = "samples.id"
	}
	direction := "ASC"
	if !filter.Ascending {
		direction = "DESC"
	}
	query = query.Order(sortField + " " + direction + ", samples.id " + direction)

	// Stronicowanie
	offset := filter.PageNumber * filter.PageSize
	query = query.Offset(offset).Limit(filter.PageSize)

	// Pobranie danych
	var samples []models.Sample
	if err := query.Preload("Code").Preload("Assortment.Group").Preload("Client").Find(&samples).Error; err != nil {
		return nil, 0, err
	}

	// Mapowanie na DTO
	var result []dto.SampleSummaryDto
	for _, s := range samples {
		summary := dto.SampleSummaryDto{
			ID:             s.ID,
			Code:           s.Code.ID,
			Group:          s.Assortment.Group.Name,
			Assortment:     s.Assortment.Name,
			ClientName:     s.Client.Name,
			AdmissionDate:  s.AdmissionDate,
			ProgressStatus: string(s.ProgressStatus),
		}
		result = append(result, summary)
	}

	return result, total, nil
}
