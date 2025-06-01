package services

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"samplelab-go/src/testutils"
)

func TestGetAllAddresses(t *testing.T) {
	_, mock := testutils.SetupMockDB(t)

	rows := sqlmock.NewRows([]string{"id", "street", "zip_code", "city"}).
		AddRow(1, "ul. Prosta", "00-001", "Warszawa").
		AddRow(2, "ul. Długa", "31-002", "Kraków")

	mock.ExpectQuery(`SELECT \* FROM "address"`).
		WillReturnRows(rows)

	addresses, err := GetAllAddresses()

	assert.NoError(t, err)
	assert.Len(t, addresses, 2)

	assert.Equal(t, int64(1), addresses[0].ID)
	assert.Equal(t, "ul. Prosta", addresses[0].Street)
	assert.Equal(t, "00-001", addresses[0].ZipCode)
	assert.Equal(t, "Warszawa", addresses[0].City)

	assert.Equal(t, int64(2), addresses[1].ID)
	assert.Equal(t, "ul. Długa", addresses[1].Street)
	assert.Equal(t, "31-002", addresses[1].ZipCode)
	assert.Equal(t, "Kraków", addresses[1].City)
}
