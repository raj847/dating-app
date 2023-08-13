package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string    `gorm:"type:varchar(11);not null"`
	Email    string    `gorm:"type:varchar(255);not null"`
	Password string    `gorm:"type:varchar(255);not null"`
	Name     string    `gorm:"type:varchar(255);not null"`
	Gender   string    `gorm:"type:varchar(255);not null"`
	DOB      time.Time `gorm:"type:date;not null"`
	Nickname string    `gorm:"type:varchar(255)"`
	Domicile string    `gorm:"type:varchar(255);not null"`
	Photo    string    `gorm:"type:varchar(255)"`
	Job      string    `gorm:"type:varchar(255)"`
	Interest string    `gorm:"type:varchar(255)"`
}
