package countdown

type Repository interface {
	GetAllCountdowns() ([]Model, error)
	CreateCountdown(countdown Model) (int64, error)
	GetCountdown(id string) (*Model, error)
	//UpdateCountdown(id string, countdown Model) (Model, error)
	//DeleteCountdown(id string) error
}
