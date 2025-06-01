package services

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"samplelab-go/src/dto"
	"samplelab-go/src/testutils"
)

func TestGetAllProductGroups_Success(t *testing.T) {
	_, mock := testutils.SetupMockDB(t)

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "Nabiał").
		AddRow(2, "Mięso")

	mock.ExpectQuery(`SELECT \* FROM "product_group"`).
		WillReturnRows(rows)

	result, err := GetAllProductGroups()

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "Nabiał", result[0].Name)
	assert.Equal(t, int64(2), result[1].ID)
}

func TestSaveProductGroup_Success(t *testing.T) {
	_, mock := testutils.SetupMockDB(t)

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "product_group"`).
		WithArgs("Słodycze").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(3))
	mock.ExpectCommit()

	err := SaveProductGroup(dto.ProductGroupSaveDto{Name: "Słodycze"})
	assert.NoError(t, err)
}

func TestUpdateProductGroup_Success(t *testing.T) {
	_, mock := testutils.SetupMockDB(t)

	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE "product_group" SET "name"=\$1 WHERE id = \$2`).
		WithArgs("Nowa nazwa", int64(4)).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err := UpdateProductGroup(4, dto.ProductGroupSaveDto{Name: "Nowa nazwa"})
	assert.NoError(t, err)
}

func TestDeleteProductGroup_Success(t *testing.T) {
	_, mock := testutils.SetupMockDB(t)

	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM "product_group" WHERE "product_group"\."id" = \$1`).
		WithArgs(int64(5)).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err := DeleteProductGroup(5)
	assert.NoError(t, err)
}
