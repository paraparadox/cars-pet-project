package handlers

import (
	"cars-pet-project/pkg/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) ManufacturersShow(c *gin.Context) {
	manufacturerID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad identifier",
		})
	}

	var manufacturer models.Manufacturer

	result := h.DB.First(&manufacturer, manufacturerID)
	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Record not found",
		})
		return
	}

	c.JSON(http.StatusOK, manufacturer)
}