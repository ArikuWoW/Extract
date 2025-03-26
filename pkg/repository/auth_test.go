package repository_test

import (
	"testing"

	"github.com/ArikuWoW/extract/models"
	"github.com/ArikuWoW/extract/pkg/repository"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestAuthRepository_CreateUser(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "postgres")
	repo := repository.NewAuthDB(sqlxDB)

	mock.ExpectQuery("INSERT INTO users").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	user := models.User{
		Login:    "testuser",
		Password: "password123",
		Name:     "Test",
		Surname:  "User",
		Email:    "test@example.com",
	}
	id, err := repo.CreateUser(user)

	assert.NoError(t, err)
	assert.Equal(t, 1, id)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestAuthRepository_GetUser(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "postgres")
	repo := repository.NewAuthDB(sqlxDB)

	mock.ExpectQuery("SELECT id, password FROM users").
		WithArgs("testuser").
		WillReturnRows(sqlmock.NewRows([]string{"id", "password"}).
			AddRow(1, "password"))

	user, err := repo.GetUser("testuser")
	assert.NoError(t, err)
	assert.Equal(t, 1, user.Id)
	assert.Equal(t, "password", user.Password)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
