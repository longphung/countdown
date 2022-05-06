package countdown

type Services struct {
	repository Repository
}

func NewService(repository Repository) *Services {
	return &Services{
		repository: repository,
	}
}

func (service *Services) GetAllCountdowns() ([]Model, error) {
	return service.repository.GetAllCountdowns()
}

func (service *Services) GetCountdown(id string) (*Model, error) {
	return service.repository.GetCountdown(id)
}

func (service *Services) AddCountdown(countdown Model) (int64, error) {
	return service.repository.CreateCountdown(countdown)
}

func (service *Services) UpdateCountdown(id string, countdown Model) (*Model, int64, error) {
	return service.repository.UpdateCountdown(id, countdown)
}

func (service *Services) DeleteCountdown(id string) error {
	return service.repository.DeleteCountdown(id)
}
