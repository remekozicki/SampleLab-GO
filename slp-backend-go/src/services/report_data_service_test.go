package services

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"samplelab-go/src/dto"
	"samplelab-go/src/testutils"
)

func TestGetAllReportData_Success(t *testing.T) {
	_, mock := testutils.SetupMockDB(t)

	rows := sqlmock.NewRows([]string{"id", "manufacturer_name", "manufacturer_country", "batch_number", "sample_id"}).
		AddRow(1, "Firma A", "Polska", 123, 999)

	mock.ExpectQuery(`SELECT \* FROM "report_data"`).
		WillReturnRows(rows)

	result, err := GetAllReportData()

	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, int64(1), result[0].ID)
	assert.Equal(t, "Firma A", result[0].ManufacturerName)
	assert.Equal(t, 123, result[0].BatchNumber)
}

func TestGetReportDataBySampleID_Success(t *testing.T) {
	_, mock := testutils.SetupMockDB(t)

	rows := sqlmock.NewRows([]string{"id", "manufacturer_name", "manufacturer_country", "batch_number", "sample_id"}).
		AddRow(2, "Firma B", "Niemcy", 456, 1000)

	mock.ExpectQuery(`SELECT \* FROM "report_data" WHERE sample_id = \$1 ORDER BY "report_data"\."id" LIMIT \$2`).
		WithArgs(1000, 1).
		WillReturnRows(rows)

	result, err := GetReportDataBySampleID(1000)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Firma B", result.ManufacturerName)
	assert.Equal(t, 456, result.BatchNumber)
}

func TestSaveReportData_Success(t *testing.T) {
	_, mock := testutils.SetupMockDB(t)

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "report_data"`).
		WithArgs(
			"Firma X", sqlmock.AnyArg(), "PL", // Adresy i country
			"Firma Y", sqlmock.AnyArg(),       // Supplier
			"Firma Z", sqlmock.AnyArg(),       // Seller
			"Odbiorca", sqlmock.AnyArg(),      // Recipient
			"2024-01-01", 1, "100L", "50L",
			"opakowanie", "miejsce", "zbieracz", 2, "mechanizm", "metoda",
			"2024-02-01", "PROTO-123", int64(999),
		).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(123))
	mock.ExpectCommit()

	err := SaveReportData(dto.ReportDataDto{
		ManufacturerName:    "Firma X",
		ManufacturerCountry: "PL",
		SupplierName:        "Firma Y",
		SellerName:          "Firma Z",
		RecipientName:       "Odbiorca",
		ProductionDate:      "2024-01-01",
		BatchNumber:         1,
		BatchSizeProd:       "100L",
		BatchSizeStorehouse: "50L",
		SamplePacking:       "opakowanie",
		SampleCollectionSite: "miejsce",
		SampleCollector:     "zbieracz",
		JobNumber:           2,
		Mechanism:           "mechanizm",
		DeliveryMethod:      "metoda",
		CollectionDate:      "2024-02-01",
		ProtocolNumber:      "PROTO-123",
		SampleID:            999,
	})
	assert.NoError(t, err)
}

func TestDeleteReportData_Success(t *testing.T) {
	_, mock := testutils.SetupMockDB(t)

	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM "report_data" WHERE "report_data"\."id" = \$1`).
		WithArgs(uint(5)).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err := DeleteReportData(5)
	assert.NoError(t, err)
}
