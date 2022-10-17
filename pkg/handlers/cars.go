package handlers

import (
	"cars-pet-project/pkg/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"net/http"
	"strconv"
)

// CarsIndex returns all the existing manufacturers
func (h *Handler) CarsIndex(c *gin.Context) {
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

	var cars []models.Car

	err = h.DB.Preload(clause.Associations).Model(&manufacturer).Association("Cars").Find(&cars)
	if err != nil {
		log.Println(err)
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
	manufacturerID, err := strconv.Atoi(c.Param("manufacturerID"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad identifier",
		})
		return
	}

	carID, err := strconv.Atoi(c.Param("carID"))
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

	// todo: look for another ways to check not found record error
	//err = h.DB.Model(&manufacturer).Association("Cars").Find(&car, carID)
	result = h.DB.
		Where("manufacturer_id = ?", manufacturerID).
		Preload("Manufacturer").
		Preload("Engine").
		Preload("Photos", func(db *gorm.DB) *gorm.DB {
			return db.Order(`"order"`)
		}).
		Find(&car, carID)

	if result.Error != nil || result.RowsAffected == 0 {
		log.Println("Record not found", result.Error)
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Record not found",
		})
		return
	}

	c.JSON(http.StatusOK, car)
}

// CarsUpdate updates a single existing manufacturer
func (h *Handler) CarsUpdate(c *gin.Context) {
	manufacturerID, err := strconv.Atoi(c.Param("manufacturerID"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad identifier",
		})
		return
	}

	carID, err := strconv.Atoi(c.Param("carID"))
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

	// todo: look for another ways to check not found record error
	result = h.DB.Where("manufacturer_id = ?", manufacturerID).Find(&car, carID)
	if result.Error != nil || result.RowsAffected == 0 {
		log.Println("Record not found", result.Error)
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

	// todo: properly handle such operations in other places too
	// todo: validate existence of foreign keys
	result = h.DB.Save(&car)
	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, car)
}

// CarsDelete deletes a single existing manufacturer
func (h *Handler) CarsDelete(c *gin.Context) {
	manufacturerID, err := strconv.Atoi(c.Param("manufacturerID"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad identifier",
		})
		return
	}

	carID, err := strconv.Atoi(c.Param("carID"))
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

	// todo: look for another ways to check not found record error
	result = h.DB.Where("manufacturer_id = ?", manufacturerID).Find(&car, carID)
	if result.Error != nil || result.RowsAffected == 0 {
		log.Println("Record not found", result.Error)
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Record not found",
		})
		return
	}

	result = h.DB.Delete(&car)
	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted car " + strconv.Itoa(carID),
	})
}
