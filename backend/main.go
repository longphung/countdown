package main

import (
	"database/sql"
	"fmt"
	gin "github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"
	"net/http"
)

type Countdown struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	TimeLeft int    `json:"timeLeft"`
}

var db *sql.DB

func getAllCountdowns() ([]Countdown, error) {
	var countdowns []Countdown
	rows, err := db.Query("SELECT * FROM countdown")

	if err != nil {
		return nil, fmt.Errorf("getAllCountdowns %v", err)
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	for rows.Next() {
		var countdown Countdown
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

func addCountdown(countdown Countdown) (int64, error) {
	result, err := db.Exec("INSERT INTO countdown (Name, time_left) VALUES (?, ?)", countdown.Name, countdown.TimeLeft)
	if err != nil {
		return 0, fmt.Errorf("addCountdown %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addCountdown %v", err)
	}
	return id, nil
}

func main() {
	db, _ = sql.Open("sqlite", "./countdown.sqlite")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/countdowns", func(c *gin.Context) {
		countdowns, err := getAllCountdowns()
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, countdowns)
	})

	r.POST("/countdowns", func(c *gin.Context) {
		var countdown Countdown
		if err := c.ShouldBindJSON(&countdown); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		id, err := addCountdown(countdown)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"id": id,
		})
	})

	apiErr := r.Run()
	if apiErr != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
