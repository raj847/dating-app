package main

import (
	"dating-app/handler"
	"dating-app/repository"
	"dating-app/service"
	"dating-app/utils"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db, err := utils.ConnectDB()
	if err != nil {
		panic(err)
	}

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Routes
	e.POST("/v1/signup", userHandler.Create)
	e.POST("/v1/login", userHandler.Login)

	// Start server
	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}
