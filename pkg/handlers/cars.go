package handlers

import (
	"cars-pet-project/pkg/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"log"
	"net/http"
	"strconv"
)

// CarsIndex returns all the existing manufacturers
func (h *Handler) CarsIndex(c *gin.Context) {
	var cars []models.Car
	result := h.DB.Preload(clause.Associations).Find(&cars)
	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, cars)
}

// CarsStore creates a single manufacturer
func (h *Handler) CarsStore(c *gin.Context) {
	manufacturerID, err := strconv.Atoi(c.Param("manufacturerID"))
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

	var car models.Car

	err = c.ShouldBindJSON(&car)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Validation error",
			"error":   err.Error(),
		})
		return
	}

	err = h.DB.Model(&manufacturer).Association("Cars").Append(&car)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusCreated, car)
}

// CarsShow returns a single existing manufacturer
func (h *Handler) CarsShow(c *gin.Context) {
	carID, err := strconv.Atoi(c.Param("carID"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad identifier",
		})
		return
	}

	var car models.Car

	result := h.DB.First(&car, carID)
	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Record not found",
		})
		return
	}

	c.JSON(http.StatusOK, car)
}

// CarsUpdate updates a single existing manufacturer
func (h *Handler) CarsUpdate(c *gin.Context) {
	carID, err := strconv.Atoi(c.Param("carID"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad identifier",
		})
		return
	}

	var car models.Car

	result := h.DB.First(&car, carID)
	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Record not found",
		})
		return
	}

	err = c.ShouldBindJSON(&car)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Validation error",
			"error":   err.Error(),
		})
		return
	}

	h.DB.Save(&car)

	c.JSON(http.StatusOK, car)
}

// CarsDelete deletes a single existing manufacturer
func (h *Handler) CarsDelete(c *gin.Context) {
	carID, err := strconv.Atoi(c.Param("carID"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad identifier",
		})
		return
	}

	var car models.Car

	result := h.DB.First(&car, carID)
	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Record not found",
		})
		return
	}

	h.DB.Delete(&car)

	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted car " + strconv.Itoa(carID),
	})

}
