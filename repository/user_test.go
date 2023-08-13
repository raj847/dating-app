package repository_test

import (
	"context"
	"dating-app/entity"
	"dating-app/repository"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupUserRepositoryTest() (*repository.UserRepository, sqlmock.Sqlmock) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(postgres.New(postgres.Config{Conn: mockDB}))
	if err != nil {
		panic(err)
	}
	userRepo := repository.NewUserRepository(db)
	return userRepo, mock
}

func TestUserRepositoryRegister_Success(t *testing.T) {
	userRepo, mock := SetupUserRepositoryTest()
	user := entity.User{
		Username: "admin",
		Email:    "aryadevara@gmail.com",
		Password: "Ary4123#",
		Name:     "Arya",
		Gender:   "L",
		DOB:      time.Now(),
		Nickname: "raj",
		Domicile: "Semarang",
		Photo:    "xxxx",
		Job:      "Student",
		Interest: "martial arts",
	}
	expectedQuery := `INSERT INTO "users" ("created_at","updated_at","deleted_at","username","email","password","name","gender","dob","nickname","domicile","photo","job","interest") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14) RETURNING "id"`
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), user.Username, user.Email, user.Password, user.Name, user.Gender, user.DOB, user.Nickname, user.Domicile, user.Photo, user.Job, user.Interest).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(101))

	mock.ExpectCommit()

	err := userRepo.Register(context.Background(), &user)
	assert.Nil(t, err)
	assert.Equal(t, uint(101), user.ID)
}

func TestUserRepositoryRegister_Error(t *testing.T) {
	userRepo, mock := SetupUserRepositoryTest()
	user := entity.User{}
	expectedQuery := `INSERT INTO "users" ("created_at","updated_at","deleted_at","username","email","password","name","gender","dob","nickname","domicile","photo","job","interest") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14) RETURNING "id"`
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), user.Username, user.Email, user.Password, user.Name, user.Gender, user.DOB, user.Nickname, user.Domicile, user.Photo, user.Job, user.Interest).
		WillReturnError(errors.New("db error"))

	mock.ExpectRollback()

	err := userRepo.Register(context.Background(), &user)
	assert.Equal(t, errors.New("db error"), err)
}

func TestUserRepository_GetUserByUsername_Success(t *testing.T) {
	userRepo, mock := SetupUserRepositoryTest()

	expectedUsername := "admin"
	expectedQuery := `SELECT * FROM "users" WHERE username = $1`

	expectedUser := entity.User{
		Model: gorm.Model{
			ID: 101,
		},
		Username: "admin",
		Email:    "aryadevara@gmail.com",
		Password: "Ary4123#",
		Name:     "Arya",
		Gender:   "L",
		DOB:      time.Now(),
		Nickname: "raj",
		Domicile: "Semarang",
		Photo:    "xxxx",
		Job:      "Student",
		Interest: "martial arts",
	}

	rows := sqlmock.NewRows([]string{"id", "username", "email", "password", "name", "gender", "dob", "nickname", "domicile", "photo", "job", "interest"}).
		AddRow(expectedUser.ID, expectedUser.Username, expectedUser.Email, expectedUser.Password, expectedUser.Name, expectedUser.Gender, expectedUser.DOB, expectedUser.Nickname, expectedUser.Domicile, expectedUser.Photo, expectedUser.Job, expectedUser.Interest)

	mock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).WithArgs(expectedUsername).WillReturnRows(rows)

	user, err := userRepo.GetUserByUsername(context.Background(), expectedUsername)
	assert.Nil(t, err)
	assert.Equal(t, expectedUser, user)
}

func TestUserRepository_GetUserByUsername_NotFound(t *testing.T) {
	userRepo, mock := SetupUserRepositoryTest()

	expectedUsername := "admin"
	expectedQuery := `SELECT * FROM "users" WHERE username = $1`

	mock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).WithArgs(expectedUsername).WillReturnError(gorm.ErrRecordNotFound)

	_, err := userRepo.GetUserByUsername(context.Background(), expectedUsername)
	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
}

func TestUserRepository_GetUserByEmail_Success(t *testing.T) {
	userRepo, mock := SetupUserRepositoryTest()

	expectedEmail := "aryadevara@gmail.com"
	expectedQuery := `SELECT * FROM "users" WHERE email = $1`

	expectedUser := entity.User{
		Model: gorm.Model{
			ID: 101,
		},
		Username: "admin",
		Email:    "aryadevara@gmail.com",
		Password: "Ary4123#",
		Name:     "Arya",
		Gender:   "L",
		DOB:      time.Now(),
		Nickname: "raj",
		Domicile: "Semarang",
		Photo:    "xxxx",
		Job:      "Student",
		Interest: "martial arts",
	}

	rows := sqlmock.NewRows([]string{"id", "username", "email", "password", "name", "gender", "dob", "nickname", "domicile", "photo", "job", "interest"}).
		AddRow(expectedUser.ID, expectedUser.Username, expectedUser.Email, expectedUser.Password, expectedUser.Name, expectedUser.Gender, expectedUser.DOB, expectedUser.Nickname, expectedUser.Domicile, expectedUser.Photo, expectedUser.Job, expectedUser.Interest)

	mock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).WithArgs(expectedEmail).WillReturnRows(rows)

	user, err := userRepo.GetUserByEmail(context.Background(), expectedEmail)
	assert.Nil(t, err)
	assert.Equal(t, expectedUser, user)
}

func TestUserRepository_GetUserByEmail_NotFound(t *testing.T) {
	userRepo, mock := SetupUserRepositoryTest()

	expectedEmail := "aryadevara@gmail.com"
	expectedQuery := `SELECT * FROM "users" WHERE email = $1`

	mock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).WithArgs(expectedEmail).WillReturnError(gorm.ErrRecordNotFound)

	_, err := userRepo.GetUserByEmail(context.Background(), expectedEmail)
	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
}
