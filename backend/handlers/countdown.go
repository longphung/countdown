package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/longphung/countdown/models"
	"github.com/longphung/countdown/services"
	"net/http"
)

type CountdownHandler struct {
	service *services.CountdownService
}

func NewCountdownHandler(service *services.CountdownService) *CountdownHandler {
	return &CountdownHandler{
		service: service,
	}
}

func (ch *CountdownHandler) GetAllCountdowns(c *gin.Context) {
	countdowns, err := ch.service.GetAllCountdowns()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, countdowns)
}

func (ch *CountdownHandler) CreateCountdown(c *gin.Context) {
	var countdown models.Countdown
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
