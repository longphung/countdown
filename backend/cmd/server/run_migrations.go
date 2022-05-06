package main

import (
	"github.com/longphung/countdown/countdown/models"
	"gorm.io/gorm"
)

func runMigrations(db *gorm.DB) {
	err := db.AutoMigrate(&models.Countdown{})
	if err != nil {
		return
	}
}
