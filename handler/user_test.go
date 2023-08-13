package handler_test

import (
	"context"
	"dating-app/entity"
	"dating-app/handler"
	"dating-app/service"
	"dating-app/service/mock_service"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUserHandlerRegister_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userService := mock_service.NewMockUserService(ctrl)
	handler := handler.NewUserHandler(userService)

	payload := `{"username": "admin",
    "email": "aryadevara@gmail.com",
    "password": "Ary4123#",
    "name": "Arya",
    "gender": "L",
    "date_of_birth": "2001-03-03",
    "nickname": "raj",
    "domicile": "Semarang",
    "photo": "xxxx",
    "job": "Student",
    "interest": "martial arts"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	dob, _ := time.Parse("2006-01-02", "2001-03-03")

	userReq := entity.User{
		Username: "admin",
		Email:    "aryadevara@gmail.com",
		Password: "Ary4123#",
		Name:     "Arya",
		Gender:   "L",
		DOB:      dob,
		Nickname: "raj",
		Domicile: "Semarang",
		Photo:    "xxxx",
		Job:      "Student",
		Interest: "martial arts",
	}

	userService.EXPECT().Register(context.Background(), &userReq).Return(nil)

	err := handler.Create(c)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
}

func TestUserHandlerRegister_EmailInvalidErr(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userService := mock_service.NewMockUserService(ctrl)
	handler := handler.NewUserHandler(userService)

	payload := `{"username": "admin",
    "email": "aryadevara@gmail.com",
    "password": "Ary4123#",
    "name": "Arya",
    "gender": "L",
    "date_of_birth": "2001-03-03",
    "nickname": "raj",
    "domicile": "Semarang",
    "photo": "xxxx",
    "job": "Student",
    "interest": "martial arts"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	expectedRes := "{\"errors\":[{\"message\":\"failed to create user\",\"code\":\"USER-DOMAIN-INVALID_CREATE-ERROR\"}]}\n"

	userService.EXPECT().Register(gomock.Any(), gomock.Any()).Return(service.ErrEmailInvalid)

	err := handler.Create(c)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, expectedRes, rec.Body.String())
}

func TestUserHandlerRegister_EmptyPassErr(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userService := mock_service.NewMockUserService(ctrl)
	handler := handler.NewUserHandler(userService)

	payload := `{"username": "admin",
    "email": "aryadevara@gmail.com",
    "name": "Arya",
    "gender": "L",
    "date_of_birth": "2001-03-03",
    "nickname": "raj",
    "domicile": "Semarang",
    "photo": "xxxx",
    "job": "Student",
    "interest": "martial arts"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	expectedRes := "{\"errors\":[{\"message\":\"Key: 'User.Password' Error:Field validation for 'Password' failed on the 'required' tag\",\"code\":\"USER_INVALID\"}]}\n"

	err := handler.Create(c)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, expectedRes, rec.Body.String())
}
