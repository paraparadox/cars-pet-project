package handlers

import (
	"cars-pet-project/pkg/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) ManufacturersStore(c *gin.Context) {
	var manufacturer models.Manufacturer

	err := c.ShouldBindJSON(&manufacturer)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Validation error",
			"error":   err.Error(),
		})
		return
	}

	result := h.DB.Create(&manufacturer)
	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusCreated, manufacturer)
}
