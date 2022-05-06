package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/longphung/countdown/countdown"
	_ "modernc.org/sqlite"
)

var db *sql.DB

func main() {
	db, _ = sql.Open("sqlite", "./countdown.sqlite")
	countdownService := countdown.NewService(db)
	countdownHandler := countdown.NewHandler(countdownService)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/countdowns", countdownHandler.GetAllCountdowns)
	r.POST("/countdowns", countdownHandler.CreateCountdown)

	apiErr := r.Run()
	if apiErr != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
