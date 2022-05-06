package countdown

type Repository interface {
	GetCountdown(id string) (Model, error)
	GetAllCountdowns() ([]Model, error)
	CreateCountdown(countdown Model) (Model, error)
	UpdateCountdown(id string, countdown Model) (Model, error)
	DeleteCountdown(id string) error
}
