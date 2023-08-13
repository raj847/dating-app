package response_test

import (
	"dating-app/entity"
	"dating-app/handler/response"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestBuildUser(t *testing.T) {
	user := entity.User{
		Model: gorm.Model{
			ID: 1,
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

	expectedUser := response.User{
		UserID:    1,
		Username:  "admin",
		Email:     "aryadevara@gmail.com",
		Password:  "Ary4123#",
		Name:      "Arya",
		Gender:    "L",
		DOB:       time.Now(),
		Nickname:  "raj",
		Domicile:  "Semarang",
		Photo:     "xxxx",
		Job:       "Student",
		Interest:  "martial arts",
		CreatedAt: "0001-01-01 00:00:00 +0000 UTC",
		UpdatedAt: "0001-01-01 00:00:00 +0000 UTC",
		DeletedAt: "0001-01-01 00:00:00 +0000 UTC",
	}

	res := response.BuildUser(user)

	assert.Equal(t, expectedUser, res)
}
