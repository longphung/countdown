package interfaces

import (
	"github.com/longphung/countdown/countdown/models"
)

type Services interface {
	GetCountdown(id string) (*models.Countdown, error)
	GetAllCountdowns() ([]models.Countdown, error)
	CreateCountdown(countdown models.Countdown) (int64, error)
	UpdateCountdown(id string, countdown models.Countdown) (*models.Countdown, int64, error)
	DeleteCountdown(id string) error
}
