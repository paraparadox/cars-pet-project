package handlers

import (
	"cars-pet-project/pkg/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) ManufacturersIndex(c *gin.Context) {
	var manufacturers []models.Manufacturer
	result := h.DB.Find(&manufacturers)
	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, manufacturers)
}
