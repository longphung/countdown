package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/longphung/countdown/countdown"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	_ "modernc.org/sqlite"
	"os"
)

func main() {
	dbLocation := os.Getenv("DB_LOCATION")
	if dbLocation == "" {
		dbLocation = "./db/db.sqlite"
	}
	db, err := gorm.Open(sqlite.Open(dbLocation), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln(err)
	}
	runMigrations(db)
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	r.Use(cors.New(config))
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
