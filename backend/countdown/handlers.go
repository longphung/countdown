package countdown

import (
	"fmt"
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

func (ch *Handler) GetCountdown(c *gin.Context) {
	countdown, err := ch.service.GetCountdown(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println(countdown)
	fmt.Println(*countdown)
	c.JSON(http.StatusOK, countdown)
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

func (ch *Handler) UpdateCountdown(c *gin.Context) {
	var countdown Model
	err := c.BindJSON(&countdown)
	if err != nil {
		return
	}
	updatedCountdown, rowsAffected, err := ch.service.UpdateCountdown(c.Param("id"), countdown)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Countdown updated",
		"data": gin.H{
			"updatedCountdown": updatedCountdown,
			"rowsAffected":     rowsAffected,
		},
	})
}

func (ch *Handler) DeleteCountdown(c *gin.Context) {
	err := ch.service.DeleteCountdown(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Countdown deleted",
		"id":      c.Param("id"),
	})
}
