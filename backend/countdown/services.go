package countdown

import (
	"database/sql"
	"fmt"
)

type Services struct {
	db *sql.DB
}

func NewService(db *sql.DB) *Services {
	return &Services{
		db: db,
	}
}

func (countdownService *Services) GetAllCountdowns() ([]Model, error) {
	var countdowns []Model
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
		var countdown Model
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

func (countdownService *Services) AddCountdown(countdown Model) (int64, error) {
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
