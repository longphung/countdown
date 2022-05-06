package models

type Countdown struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	TimeLeft int    `json:"timeLeft"`
}
