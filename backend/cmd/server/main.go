package main

import (
	"github.com/gin-gonic/gin"
	"github.com/longphung/countdown/countdown"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	_ "modernc.org/sqlite"
)

func main() {
	db, err := gorm.Open(sqlite.Open("./countdown.sqlite"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln(err)
	}
	runMigrations(db)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	countdown.Init(r, db)

	apiErr := r.Run()
	if apiErr != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
