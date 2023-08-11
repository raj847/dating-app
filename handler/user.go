package handler

import (
	"dating-app/entity"
	"dating-app/handler/request"
	"dating-app/handler/response"
	"dating-app/service"
	"dating-app/validate"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/minio/minio-go/v7"
)

type UserHandler struct {
	userService *service.UserService
	minioClient *minio.Client
}

func NewUserHandler(
	userService *service.UserService,
	minioClient *minio.Client,
) *UserHandler {
	return &UserHandler{
		userService: userService,
		minioClient: minioClient,
	}
}

func (u *UserHandler) Create(c echo.Context) error {

	var userReq request.User
	err := c.Bind(&userReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: []response.ErrorDetail{
				{
					Message: "failed to read form request",
					Code:    "BAD_REQUEST",
				},
			},
		})
	}
	err = validate.Validate(userReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: []response.ErrorDetail{
				{
					Message: err.Error(),
					Code:    "USER_INVALID",
				},
			},
		})
	}
	fileHeader, err := c.FormFile("photo")
	if err != nil {
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.ErrorResponse{
				Errors: []response.ErrorDetail{
					{
						Message: err.Error(),
						Code:    "FILE_INVALID",
					},
				},
			})
		}
	}
	file, err := fileHeader.Open()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to open the file.")
	}
	defer file.Close()

	_, err = u.minioClient.PutObject(c.Request().Context(), "rajendra", fileHeader.Filename, file, fileHeader.Size, minio.PutObjectOptions{
		UserMetadata: map[string]string{
			"x-amz-acl": "public-read",
		},
		ContentType: "image/jpeg",
	})
	if err != nil {
		log.Println(err)
		return c.String(http.StatusInternalServerError, "Failed to upload to storage.")
	}
	fileName := fmt.Sprintf("https://is3.cloudhost.id/rajendra/%s", fileHeader.Filename)
	t, _ := time.Parse("2006-01-02", userReq.DOB)
	user := entity.User{
		Username: userReq.Username,
		Email:    userReq.Email,
		Password: userReq.Password,
		Name:     userReq.Name,
		Gender:   userReq.Gender,
		DOB:      t,
		Nickname: userReq.Nickname,
		Domicile: userReq.Domicile,
		Photo:    fileName,
		Job:      userReq.Job,
		Interest: userReq.Interest,
	}
	err = u.userService.Register(c.Request().Context(), user)
	if err != nil {
		if errors.Is(err, service.ErrEmailInvalid) {
			return c.JSON(http.StatusBadRequest, response.ErrorResponse{
				Errors: []response.ErrorDetail{
					{
						Message: "failed to create user",
						Code:    "USER-DOMAIN-INVALID_CREATE-ERROR",
					},
				},
			})
		}
		if errors.Is(err, service.ErrUserAlreadyExists) {
			return c.JSON(http.StatusBadRequest, response.ErrorResponse{
				Errors: []response.ErrorDetail{
					{
						Message: "failed to create user",
						Code:    "USER-ALREADY-EXIST_CREATE-ERROR",
					},
				},
			})
		}
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: []response.ErrorDetail{
				{
					Message: "failed to create user",
					Code:    "USER_CREATE-ERROR",
				},
			},
		})
	}
	res := response.BuildUser(user)
	return c.JSON(http.StatusCreated, res)
}

func (u *UserHandler) Login(c echo.Context) error {

	var userReq request.Login
	err := c.Bind(&userReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: []response.ErrorDetail{
				{
					Message: "failed to read json request",
					Code:    "BAD_REQUEST",
				},
			},
		})
	}
	User := entity.User{
		Username: userReq.Chars,
		Email:    userReq.Chars,
		Password: userReq.Password,
	}
	res, err := u.userService.LoginUser(c.Request().Context(), User)
	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			return c.JSON(http.StatusBadRequest, response.ErrorResponse{
				Errors: []response.ErrorDetail{
					{
						Message: "failed to login user",
						Code:    "USER-NOT-FOUND_LOGIN-ERROR",
					},
				},
			})
		}
		if errors.Is(err, service.ErrUserPasswordDontMatch) {
			return c.JSON(http.StatusBadRequest, response.ErrorResponse{
				Errors: []response.ErrorDetail{
					{
						Message: "failed to login user",
						Code:    "PASSWORD-NOT-MATCH_LOGIN-ERROR",
					},
				},
			})
		}
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: []response.ErrorDetail{
				{
					Message: "failed to login user",
					Code:    "USER_LOGIN-ERROR",
				},
			},
		})
	}
	expiresAt := time.Now().Add(5 * time.Hour)
	claims := entity.Claims{
		UserID:   res.ID,
		Username: res.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	tokenString, _ := token.SignedString([]byte("rahasia-perusahaan"))

	resp := map[string]any{
		"data":  res,
		"token": tokenString,
	}
	return c.JSON(http.StatusCreated, resp)
}
