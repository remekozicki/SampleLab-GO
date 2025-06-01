package services

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"samplelab-go/src/testutils"
)

func TestGetAllExaminationsForSample_Success(t *testing.T) {
	_, mock := testutils.SetupMockDB(t)

	rows := sqlmock.NewRows([]string{
		"id", "indication_id", "sample_id", "signage", "nutritional_value",
		"specification", "regulation", "samples_number", "result",
		"start_date", "end_date", "method_status", "uncertainty", "lod", "loq",
	}).AddRow(
		1, 100, 200, "EX-001", "Białko", "spec-A", "EU123",
		3, "OK", "2025-01-01", "2025-01-03", "Zatwierdzono", 0.1, 0.01, 0.05,
	)

	mock.ExpectQuery(`SELECT \* FROM "examination" WHERE sample_id = \$1`).
		WithArgs(200).
		WillReturnRows(rows)

	result, err := GetAllExaminationsForSample(200)

	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, int64(1), result[0].ID)
	assert.Equal(t, "EX-001", result[0].Signage)
	assert.Equal(t, 0.1, result[0].Uncertainty)
}
