package services

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"samplelab-go/src/dto"
	"samplelab-go/src/testutils"
)

func TestGetAllCodes_Success(t *testing.T) {
	_, mock := testutils.SetupMockDB(t)

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow("ABC", "Kod ABC").
		AddRow("XYZ", "Kod XYZ")

	mock.ExpectQuery(`SELECT \* FROM "code"`).
		WillReturnRows(rows)

	result, err := GetAllCodes()

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "ABC", result[0].ID)
	assert.Equal(t, "Kod XYZ", result[1].Name)
}

func TestGetAllCodes_DBError(t *testing.T) {
	_, mock := testutils.SetupMockDB(t)

	mock.ExpectQuery(`SELECT \* FROM "code"`).
		WillReturnError(errors.New("db error"))

	result, err := GetAllCodes()

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestSaveCode_Success(t *testing.T) {
	_, mock := testutils.SetupMockDB(t)

	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE "code" SET`).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	code := dto.CodeDto{
		ID:   "ABC",
		Name: "Kod ABC",
	}

	err := SaveCode(code)
	assert.NoError(t, err)
}

func TestDeleteCode_Success(t *testing.T) {
	_, mock := testutils.SetupMockDB(t)

	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM "code" WHERE id = \$1`).
		WithArgs("XYZ").
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err := DeleteCode("XYZ")
	assert.NoError(t, err)
}

func TestDeleteCode_DBError(t *testing.T) {
	_, mock := testutils.SetupMockDB(t)

	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM "code" WHERE id = \$1`).
		WithArgs("ERR").
		WillReturnError(errors.New("constraint error"))
	mock.ExpectRollback()

	err := DeleteCode("ERR")
	assert.Error(t, err)
}
