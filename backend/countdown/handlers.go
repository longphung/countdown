package countdown

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	service *Services
}

func NewHandler(service *Services) *Handler {
	return &Handler{
		service: service,
	}
}

func (ch *Handler) GetAllCountdowns(c *gin.Context) {
	countdowns, err := ch.service.GetAllCountdowns()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, countdowns)
}

func (ch *Handler) CreateCountdown(c *gin.Context) {
	var countdown Model
	err := c.BindJSON(&countdown)
	if err != nil {
		return
	}
	id, err := ch.service.AddCountdown(countdown)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}
