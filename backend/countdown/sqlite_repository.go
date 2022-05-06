package countdown

import (
	"fmt"
	"github.com/longphung/countdown/countdown/models"
	"gorm.io/gorm"
)

type SQLiteRepository struct {
	db *gorm.DB
}

func NewSQLiteRepository(db *gorm.DB) *SQLiteRepository {
	return &SQLiteRepository{
		db: db,
	}
}

func (repo *SQLiteRepository) GetAllCountdowns() ([]models.Countdown, error) {
	var countdowns []models.Countdown
	result := repo.db.Find(&countdowns)
	if result.Error != nil {
		return nil, result.Error
	}
	return countdowns, nil
}

func (repo *SQLiteRepository) GetCountdown(id string) (*models.Countdown, error) {
	return repo.getCountdown(id)
}

func (repo *SQLiteRepository) CreateCountdown(countdown models.Countdown) (int64, error) {
	result := repo.db.Create(&countdown)
	if result.Error != nil {
		return 0, fmt.Errorf("createCountdown %v", result.Error)
	}
	return int64(countdown.ID), nil
}

func (repo *SQLiteRepository) UpdateCountdown(id string, countdown models.Countdown) (*models.Countdown, int64, error) {
	//result, err := repo.ldb.Exec("UPDATE countdown SET name = ?, time_left = ? WHERE id = ?", countdown.Name, countdown.TimeLeft, id)
	result := repo.db.Model(&countdown).Where("id = ?", id).Updates(models.Countdown{Name: countdown.Name, DueDate: countdown.DueDate})
	if result.Error != nil {
		return nil, 0, result.Error
	}
	updatedCountdown, err := repo.getCountdown(id)
	if err != nil {
		return nil, 0, err
	}
	return updatedCountdown, result.RowsAffected, nil
}

func (repo *SQLiteRepository) DeleteCountdown(id string) error {
	result := repo.db.Delete(&models.Countdown{}, id)
	if result.Error != nil {
		return fmt.Errorf("deleteCountdown %v", result.Error)
	}
	return nil
}

//==============INTERNAL=============

func (repo *SQLiteRepository) getCountdown(id string) (*models.Countdown, error) {
	var countdown models.Countdown
	result := repo.db.Find(&countdown, "id = ?", id)
	if result.Error != nil {
		return nil, fmt.Errorf("getCountdown %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("getCountdown %v", "countdown not found")
	}
	return &countdown, nil
}
