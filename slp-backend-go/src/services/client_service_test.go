package services

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"samplelab-go/src/testutils"
)

func TestGetAllClients_ReturnsList(t *testing.T) {
	_, mock := testutils.SetupMockDB(t)

	rows := sqlmock.NewRows([]string{"id", "wijhars_code", "name", "address_id"}).
		AddRow(1, "WJ-001", "Klient 1", 10)

	addressRows := sqlmock.NewRows([]string{"id", "street", "zip_code", "city"}).
		AddRow(10, "ul. Prosta", "00-001", "Warszawa")

	mock.ExpectQuery(`SELECT \* FROM "client"`).WillReturnRows(rows)
	mock.ExpectQuery(`SELECT \* FROM "address" WHERE "address"\."id" = \$1`).
		WithArgs(10).WillReturnRows(addressRows)

	result, err := GetAllClients()

	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "Klient 1", result[0].Name)
	assert.Equal(t, "ul. Prosta", result[0].Address.Street)
}
