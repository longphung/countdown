package main

import (
	"github.com/gin-gonic/gin"
	"github.com/longphung/countdown/countdown"
	_ "modernc.org/sqlite"
)

func main() {
	countdownSQLiteRepository := countdown.NewSQLiteRepository()
	countdownService := countdown.NewService(countdownSQLiteRepository)
	countdownHandler := countdown.NewHandler(countdownService)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/countdowns", countdownHandler.GetAllCountdowns)
	r.GET("/countdown/:id", countdownHandler.GetCountdown)
	r.POST("/countdown", countdownHandler.CreateCountdown)
	r.PATCH("/countdown/:id", countdownHandler.UpdateCountdown)
	r.DELETE("/countdown/:id", countdownHandler.DeleteCountdown)

	apiErr := r.Run()
	if apiErr != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
