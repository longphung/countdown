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

func (service *Services) AddCountdown(countdown Model) (int64, error) {
	return service.repository.CreateCountdown(countdown)
}
