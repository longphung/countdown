package models

import (
	"gorm.io/gorm"
	"time"
)

type Countdown struct {
	gorm.Model
	Name    string    `json:"name" binding:"required"`
	DueDate time.Time `json:"dueDate" binding:"required"`
}
