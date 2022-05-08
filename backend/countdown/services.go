package countdown

import (
	"errors"
	"github.com/longphung/countdown/countdown/interfaces"
	"github.com/longphung/countdown/countdown/models"
)

type Services struct {
	repository interfaces.Repository
}

var (
	ErrInvalidId = errors.New("invalid id")
	ErrNotFound  = errors.New("movie not found")
	ErrNoDueDate = errors.New("no due date")
)

func NewService(repository interfaces.Repository) *Services {
	return &Services{
		repository: repository,
	}
}

func (service *Services) GetAllCountdowns() ([]models.Countdown, error) {
	return service.repository.GetAllCountdowns()
}

func (service *Services) GetCountdown(id string) (*models.Countdown, error) {
	return service.repository.GetCountdown(id)
}

func (service *Services) CreateCountdown(countdown models.Countdown) (int64, error) {
	return service.repository.CreateCountdown(countdown)
}

func (service *Services) UpdateCountdown(id string, countdown models.Countdown) (*models.Countdown, int64, error) {
	return service.repository.UpdateCountdown(id, countdown)
}

func (service *Services) DeleteCountdown(id string) error {
	return service.repository.DeleteCountdown(id)
}
