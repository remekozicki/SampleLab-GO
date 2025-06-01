package services

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"samplelab-go/src/dto"
	"samplelab-go/src/enum"
	"samplelab-go/src/testutils"
)

func TestRegisterUser_EmailTaken(t *testing.T) {
	_, mock := testutils.SetupMockDB(t)

	mock.ExpectQuery(`SELECT \* FROM "users" WHERE email = \$1 ORDER BY "users"\."id" LIMIT \$2`).
		WithArgs("jan@example.com", 1).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	input := dto.RegisterInput{
		Name:  "Jan",
		Email: "jan@example.com",
		Role:  enum.RoleWorker,
	}

	user, err := RegisterUser(input)

	assert.Nil(t, user)
	assert.Equal(t, ErrEmailTaken, err)
}

func TestRegisterUser_Success(t *testing.T) {
	_, mock := testutils.SetupMockDB(t)

	mock.ExpectQuery(`SELECT \* FROM "users" WHERE email = \$1 ORDER BY "users"\."id" LIMIT \$2`).
		WithArgs("jan@example.com", 1).
		WillReturnRows(sqlmock.NewRows([]string{}))

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "users" \("name","email","password","role"\) VALUES \(\$1,\$2,\$3,\$4\) RETURNING "id"`).
		WithArgs("Jan", "jan@example.com", sqlmock.AnyArg(), int64(enum.RoleWorker)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	input := dto.RegisterInput{
		Name:  "Jan",
		Email: "jan@example.com",
		Role:  enum.RoleWorker,
	}

	user, err := RegisterUser(input)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, input.Name, user.Name)
	assert.Equal(t, input.Email, user.Email)
	assert.Equal(t, input.Role, user.Role)
	assert.NotEmpty(t, user.Password)
}
