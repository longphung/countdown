package countdown

import (
	"database/sql"
	"fmt"
	"log"
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository() *SQLiteRepository {
	db, err := sql.Open("sqlite", "./countdown.sqlite")
	if err != nil {
		log.Fatalln(err)
	}
	return &SQLiteRepository{
		db: db,
	}
}

func (repo *SQLiteRepository) GetAllCountdowns() ([]Model, error) {
	var countdowns []Model
	rows, err := repo.db.Query("SELECT * FROM countdown")
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

func (repo *SQLiteRepository) CreateCountdown(countdown Model) (int64, error) {
	result, err := repo.db.Exec("INSERT INTO countdown (Name, time_left) VALUES (?, ?)", countdown.Name, countdown.TimeLeft)
	if err != nil {
		return 0, fmt.Errorf("createCountdown %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("createCountdown %v", err)
	}
	return id, nil
}
