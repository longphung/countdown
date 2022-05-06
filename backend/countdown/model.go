package countdown

type Model struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	TimeLeft int    `json:"timeLeft"`
}
