package handlers

import (
	"cars-pet-project/pkg/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
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

func (h *Handler) ManufacturersUpdate(c *gin.Context) {
	manufacturerID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad identifier",
		})
		return
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

	err = c.ShouldBindJSON(&manufacturer)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Validation error",
			"error":   err.Error(),
		})
		return
	}

	h.DB.Save(&manufacturer)

	c.JSON(http.StatusOK, manufacturer)
}

func (h *Handler) ManufacturersDelete(c *gin.Context) {
	manufacturerID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad identifier",
		})
		return
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

	h.DB.Delete(&manufacturer)

	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted manufacturer " + strconv.Itoa(manufacturerID),
	})

}
