package models

import (
	"github.com/longphung/countdown/utils"
	"time"
)

type Countdown struct {
	utils.CommonModelFields
	Name    string    `json:"name" binding:"required" gorm:"not null"`
	DueDate time.Time `json:"dueDate" binding:"required" gorm:"not null"`
}
