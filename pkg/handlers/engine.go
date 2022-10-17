package handlers

import (
	"cars-pet-project/pkg/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// EngineStore creates a single engine
func (h *Handler) EngineStore(c *gin.Context) {
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

	var engine models.Engine

	err = c.ShouldBindJSON(&engine)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Validation error",
			"error":   err.Error(),
		})
		return
	}

	err = h.DB.Model(&car).Association("Engine").Append(&engine)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusCreated, engine)
}

// EngineShow returns an engine of specified car
func (h *Handler) EngineShow(c *gin.Context) {
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
	//err = h.DB.Model(&manufacturer).Association("Engine").Find(&car, carID)
	result = h.DB.Where("manufacturer_id = ?", manufacturerID).Find(&car, carID)
	if result.Error != nil || result.RowsAffected == 0 {
		log.Println("Record not found", result.Error)
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Record not found",
		})
		return
	}

	var engine models.Engine

	result = h.DB.Where("car_id = ?", carID).Find(&engine)
	if result.Error != nil || result.RowsAffected == 0 {
		log.Println("Record not found", result.Error)
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Record not found",
		})
		return
	}

	c.JSON(http.StatusOK, engine)
}

// EngineUpdate updates a single existing manufacturer
func (h *Handler) EngineUpdate(c *gin.Context) {
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

	var engine models.Engine

	result = h.DB.Where("car_id = ?", carID).Find(&engine)
	if result.Error != nil || result.RowsAffected == 0 {
		log.Println("Record not found", result.Error)
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Record not found",
		})
		return
	}

	err = c.ShouldBindJSON(&engine)
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
	result = h.DB.Save(&engine)
	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, engine)
}

// EngineDelete deletes a single existing manufacturer
func (h *Handler) EngineDelete(c *gin.Context) {
	manufacturerID, err := strconv.Atoi(c.Param(POST"manufacturerID"))
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
