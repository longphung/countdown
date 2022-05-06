package countdown

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Init(router *gin.Engine, db *gorm.DB) {
	countdownSQLiteRepository := NewSQLiteRepository(db)
	countdownService := NewService(countdownSQLiteRepository)
	countdownHandler := NewHandler(countdownService)

	router.GET("/countdowns", countdownHandler.GetAllCountdowns)
	router.GET("/countdown/:id", countdownHandler.GetCountdown)
	router.POST("/countdown", countdownHandler.CreateCountdown)
	router.PATCH("/countdown/:id", countdownHandler.UpdateCountdown)
	router.DELETE("/countdown/:id", countdownHandler.DeleteCountdown)
}
