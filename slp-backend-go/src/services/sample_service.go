package services

import (
	"errors"
	"fmt"
	"samplelab-go/src/enum"
	"strings"

	"gorm.io/gorm"
	"samplelab-go/src/db"
	"samplelab-go/src/dto"
	"samplelab-go/src/models"
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
		//Distinct("sample.id").
		Joins("JOIN code ON sample.code_id = code.id").
		Joins("JOIN client ON sample.client_id = client.id").
		Joins("JOIN assortment ON sample.assortment_id = assortment.id").
		Joins("JOIN product_group ON assortment.group_id = product_group.id").
		Joins("LEFT JOIN examination ON examination.sample_id = sample.id").
		Joins("LEFT JOIN indication ON examination.indication_id = indication.id").
		Group("sample.id")

	// Fuzzy search
	if strings.TrimSpace(filter.FuzzySearch) != "" {
		pattern := "%" + strings.ToLower(filter.FuzzySearch) + "%"
		query = query.Where(
			dbConn.Where("LOWER(code.id::text) LIKE ?", pattern).
				Or("LOWER(client.name) LIKE ?", pattern).
				Or("LOWER(assortment.name) LIKE ?", pattern).
				Or("LOWER(product_group.name) LIKE ?", pattern).
				Or("LOWER(indication.name) LIKE ?", pattern).
				Or("LOWER(indication.method) LIKE ?", pattern),
		)
	}

	// Filtry po polach
	if filter.Filters != nil {
		if len(filter.Filters.Code) > 0 {
			query = query.Where("code.id IN ?", filter.Filters.Code)
		}
		if len(filter.Filters.Client) > 0 {
			query = query.Where("client.name IN ?", filter.Filters.Client)
		}
		if len(filter.Filters.Groups) > 0 {
			query = query.Where("product_group.name IN ?", filter.Filters.Groups)
		}
		if len(filter.Filters.ProgressStatuses) > 0 {
			query = query.Where("sample.progress_status IN ?", filter.Filters.ProgressStatuses)
		}
	}

	// Zlicz rekordy do paginacji
	var total int64
	query.Count(&total)

	fmt.Println("FieldName received:", filter.FieldName)

	// Sortowanie
	sortField := "sample.id"
	if filter.FieldName != "" {
		sortField = filter.FieldName
	}
	direction := "ASC"
	if !filter.Ascending {
		direction = "DESC"
	}
	query = query.Order(sortField + " " + direction + ", sample.id " + direction)

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
			ProgressStatus: s.ProgressStatus,
		}
		result = append(result, summary)
	}

	return result, total, nil
}

func GetFilters() (dto.FilterFields, error) {
	conn := db.GetDB()

	var codes []string
	if err := conn.Table("code").Select("id").Scan(&codes).Error; err != nil {
		return dto.FilterFields{}, err
	}

	var clients []string
	if err := conn.Table("client").Select("name").Scan(&clients).Error; err != nil {
		return dto.FilterFields{}, err
	}

	var groups []string
	if err := conn.
		Table("product_group").
		Select("DISTINCT product_group.name").
		Joins("JOIN assortment ON assortment.group_id = product_group.id").
		Scan(&groups).Error; err != nil {
		return dto.FilterFields{}, err
	}

	return dto.FilterFields{
		Code:   codes,
		Client: clients,
		Groups: groups,
	}, nil
}

func CountSamples() (int64, error) {
	var count int64
	err := db.GetDB().Model(&models.Sample{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func UpdateSampleStatus(id uint, progress enum.ProgressStatus) (*models.Sample, error) {
	dbConn := db.GetDB()

	var sample models.Sample
	if err := dbConn.First(&sample, id).Error; err != nil {
		return nil, errors.New("próbka o podanym ID nie istnieje")
	}

	sample.ProgressStatus = progress

	if err := dbConn.Save(&sample).Error; err != nil {
		return nil, err
	}

	return &sample, nil
}
