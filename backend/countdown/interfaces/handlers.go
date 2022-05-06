package interfaces

import "github.com/gin-gonic/gin"

type Handlers interface {
	GetCountdown(c *gin.Context)
	GetAllCountdowns(c *gin.Context)
	CreateCountdown(c *gin.Context)
	UpdateCountdown(c *gin.Context)
	DeleteCountdown(c *gin.Context)
}
