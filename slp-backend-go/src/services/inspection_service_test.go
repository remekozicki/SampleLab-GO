package services

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"samplelab-go/src/dto"
	"samplelab-go/src/testutils"
)

func TestGetAllInspections_Success(t *testing.T) {
	_, mock := testutils.SetupMockDB(t)

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "Kontrola A").
		AddRow(2, "Kontrola B")

	mock.ExpectQuery(`SELECT \* FROM "inspection"`).
		WillReturnRows(rows)

	result, err := GetAllInspections()

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "Kontrola A", result[0].Name)
	assert.Equal(t, int64(2), result[1].ID)
}

func TestUpdateInspection_Success(t *testing.T) {
	_, mock := testutils.SetupMockDB(t)

	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE "inspection" SET "name"=\$1 WHERE id = \$2`).
		WithArgs("Zmieniona nazwa", int64(5)).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err := UpdateInspection(dto.InspectionDto{
		ID:   5,
		Name: "Zmieniona nazwa",
	})

	assert.NoError(t, err)
}

func TestDeleteInspection_Success(t *testing.T) {
	_, mock := testutils.SetupMockDB(t)

	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM "inspection" WHERE "inspection"\."id" = \$1`).
		WithArgs(int64(10)).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err := DeleteInspection(10)

	assert.NoError(t, err)
}
