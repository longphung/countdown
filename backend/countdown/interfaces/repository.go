package interfaces

import (
	"github.com/longphung/countdown/countdown/models"
)

type Repository interface {
	GetAllCountdowns() ([]models.Countdown, error)
	CreateCountdown(countdown models.Countdown) (int64, error)
	GetCountdown(id string) (*models.Countdown, error)
	UpdateCountdown(id string, countdown models.Countdown) (*models.Countdown, int64, error)
	DeleteCountdown(id string) error
}
