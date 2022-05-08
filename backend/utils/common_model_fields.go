package utils

import (
	"gorm.io/gorm"
	"time"
)

type CommonModelFields struct {
	ID        uint           `gorm:"primaryKey" json:"id,omitempty"`
	CreatedAt time.Time      `json:"createdAt,omitempty"`
	UpdatedAt time.Time      `json:"updatedAt,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
}
