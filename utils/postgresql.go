package utils

import (
	"dating-app/entity"
	"os"

	// _ "github.com/jackc/pgx/v4/stdlib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	// connect using gorm pgx
	conn, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: "pgx",
		DSN:        os.Getenv("DATABASE_URL"),
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	conn.AutoMigrate(
		entity.User{},
	)

	SetupDBConnection(conn)
	return conn, nil
}

func SetupDBConnection(DB *gorm.DB) {
	db = DB
}

func GetDBConnection() *gorm.DB {
	return db
}