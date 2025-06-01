package services

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"samplelab-go/src/dto"
	"samplelab-go/src/testutils"
)

func TestGetAllAssortments_ReturnsMappedList(t *testing.T) {
	_, mock := testutils.SetupMockDB(t)

	rows := sqlmock.NewRows([]string{"id", "name", "organoleptic_method", "group_id"}).
		AddRow(1, "Kasza", "ZXC-ASD1", 10).
		AddRow(2, "Winogrona", "ZXC-ASD1", 20)

	mock.ExpectQuery(`SELECT \* FROM "assortment"`).
		WillReturnRows(rows)

	result := GetAllAssortments()

	assert.Len(t, result, 2)
	assert.Equal(t, "Kasza", result[0].Name)
	assert.Equal(t, "ZXC-ASD1", result[0].OrganolepticMethod)
	assert.Equal(t, int64(10), result[0].GroupID)

	assert.Equal(t, "Winogrona", result[1].Name)
	assert.Equal(t, "ZXC-ASD1", result[1].OrganolepticMethod)
	assert.Equal(t, int64(20), result[1].GroupID)
}

func TestSaveAssortment_Success(t *testing.T) {
	_, mock := testutils.SetupMockDB(t)

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "assortment"`).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	dto := dto.AssortmentDto{
		Name:               "Kasza",
		OrganolepticMethod: "ZXC-ASD1",
		GroupID:            99,
	}

	err := SaveAssortment(dto)

	assert.NoError(t, err)
}

func TestUpdateAssortment_Success(t *testing.T) {
	_, mock := testutils.SetupMockDB(t)

	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE "assortment"`).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	dto := dto.AssortmentDto{
		ID:                 1,
		Name:               "Winogrona",
		OrganolepticMethod: "ZXC-ASD1",
		GroupID:            101,
	}

	err := UpdateAssortment(dto)

	assert.NoError(t, err)
}

func TestDeleteAssortmentByID_Success(t *testing.T) {
	_, mock := testutils.SetupMockDB(t)

	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM "assortment" WHERE "assortment"\."id" = \$1`).
		WithArgs("5").
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err := DeleteAssortmentByID("5")

	assert.NoError(t, err)
}
