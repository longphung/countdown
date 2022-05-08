package countdown

import (
	"github.com/gin-gonic/gin"
	"github.com/longphung/countdown/countdown/interfaces"
	"github.com/longphung/countdown/countdown/models"
	"net/http"
)

type Handler struct {
	service interfaces.Services
}

func NewHandler(service interfaces.Services) *Handler {
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
	var params struct {
		Id string `uri:"id" binding:"numeric"`
	}
	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	countdown, err := ch.service.GetCountdown(params.Id)
	switch err {
	case ErrNotFound:
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	case nil:
		c.JSON(http.StatusOK, countdown)
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}

func (ch *Handler) CreateCountdown(c *gin.Context) {
	var countdown models.Countdown
	err := c.BindJSON(&countdown)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	id, err := ch.service.CreateCountdown(countdown)
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
	var countdown models.Countdown
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
	var params struct {
		Id string `uri:"id" binding:"numeric"`
	}
	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err := ch.service.DeleteCountdown(params.Id)
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
