package services

import (
	"database/sql"
	"fmt"
	models "github.com/longphung/countdown/models"
)

type CountdownService struct {
	db *sql.DB
}

func NewCountdownService(db *sql.DB) *CountdownService {
	return &CountdownService{
		db: db,
	}
}

func (countdownService *CountdownService) GetAllCountdowns() ([]models.Countdown, error) {
	var countdowns []models.Countdown
	rows, err := countdownService.db.Query("SELECT * FROM countdown")

	if err != nil {
		return nil, fmt.Errorf("getAllCountdowns %v", err)
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	for rows.Next() {
		var countdown models.Countdown
		if err := rows.Scan(&countdown.Id, &countdown.Name, &countdown.TimeLeft); err != nil {
			return nil, fmt.Errorf("getAllCountdowns %v", err)
		}
		countdowns = append(countdowns, countdown)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getAllCountdowns %v", err)
	}

	return countdowns, nil
}

func (countdownService *CountdownService) AddCountdown(countdown models.Countdown) (int64, error) {
	result, err := countdownService.db.Exec("INSERT INTO countdown (Name, time_left) VALUES (?, ?)", countdown.Name, countdown.TimeLeft)
	if err != nil {
		return 0, fmt.Errorf("addCountdown %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addCountdown %v", err)
	}
	return id, nil
}
