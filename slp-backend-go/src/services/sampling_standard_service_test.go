package services

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"samplelab-go/src/dto"
	"samplelab-go/src/testutils"
)

func TestGetAllSamplingStandards_Success(t *testing.T) {
	_, mock := testutils.SetupMockDB(t)

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "Standard A").
		AddRow(2, "Standard B")

	mock.ExpectQuery(`SELECT \* FROM "sampling_standard"`).
		WillReturnRows(rows)

	result, err := GetAllSamplingStandards()

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "Standard A", result[0].Name)
	assert.Equal(t, int64(2), result[1].ID)
}

func TestUpdateSamplingStandard_Success(t *testing.T) {
	_, mock := testutils.SetupMockDB(t)

	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE "sampling_standard" SET "name"=\$1 WHERE id = \$2`).
		WithArgs("Zmieniony standard", int64(1)).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err := UpdateSamplingStandard(dto.SamplingStandardDto{
		ID:   1,
		Name: "Zmieniony standard",
	})
	assert.NoError(t, err)
}

func TestDeleteSamplingStandard_Success(t *testing.T) {
	_, mock := testutils.SetupMockDB(t)

	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM "sampling_standard" WHERE "sampling_standard"\."id" = \$1`).
		WithArgs(int64(5)).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err := DeleteSamplingStandard(5)
	assert.NoError(t, err)
}
