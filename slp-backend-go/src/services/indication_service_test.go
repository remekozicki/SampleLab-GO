package services

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"samplelab-go/src/dto"
	"samplelab-go/src/testutils"
)

func TestGetAllIndications_Success(t *testing.T) {
	_, mock := testutils.SetupMockDB(t)

	rows := sqlmock.NewRows([]string{"id", "name", "method", "unit", "laboratory", "is_organoleptic"}).
		AddRow(1, "Tłuszcz", "Metoda A", "g", "LAB 1", true).
		AddRow(2, "Białko", "Metoda B", "mg", "LAB 2", false)

	mock.ExpectQuery(`SELECT \* FROM "indication"`).
		WillReturnRows(rows)

	result, err := GetAllIndications()

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "Tłuszcz", result[0].Name)
	assert.True(t, result[0].IsOrganoleptic)
}

func TestGetIndicationByID_Success(t *testing.T) {
	_, mock := testutils.SetupMockDB(t)

	row := sqlmock.NewRows([]string{"id", "name", "method", "unit", "laboratory", "is_organoleptic"}).
		AddRow(5, "Woda", "Metoda C", "ml", "LAB W", false)

	mock.ExpectQuery(`SELECT \* FROM "indication" WHERE "indication"\."id" = \$1 ORDER BY "indication"\."id" LIMIT \$2`).
		WithArgs(5, 1).
		WillReturnRows(row)

	result, err := GetIndicationByID(5)

	assert.NoError(t, err)
	assert.Equal(t, "Woda", result.Name)
	assert.False(t, result.IsOrganoleptic)
}

func TestSaveIndication_Success(t *testing.T) {
	_, mock := testutils.SetupMockDB(t)

	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE "indication" SET`).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	input := dto.IndicationDto{
		ID:             1,
		Name:           "Nowa",
		Method:         "X",
		Unit:           "g",
		Laboratory:     "LAB X",
		IsOrganoleptic: true,
	}

	err := SaveIndication(input)
	assert.NoError(t, err)
}

func TestDeleteIndication_Success(t *testing.T) {
	_, mock := testutils.SetupMockDB(t)

	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM "indication" WHERE "indication"\."id" = \$1`).
		WithArgs(7).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err := DeleteIndication(7)
	assert.NoError(t, err)
}
