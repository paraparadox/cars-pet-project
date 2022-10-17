package handlers

import (
	"cars-pet-project/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
)

// PhotosIndex gets all the photos of a specified car
func (h *Handler) PhotosIndex(c *gin.Context) {
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

	var photos []models.Photo

	err = h.DB.Model(&car).Association("Photos").Find(&photos)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, photos)
}

// PhotosStore returns an engine of specified car
func (h *Handler) PhotosStore(c *gin.Context) {
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
	//err = h.DB.Model(&manufacturer).Association("Photo").Find(&car, carID)
	result = h.DB.Where("manufacturer_id = ?", manufacturerID).Find(&car, carID)
	if result.Error != nil || result.RowsAffected == 0 {
		log.Println("Record not found", result.Error)
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Record not found",
		})
		return
	}

	uploadedPhoto, err := c.FormFile("photo")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}
	extension := filepath.Ext(uploadedPhoto.Filename)
	newFileName := "assets/cars-photos/"
	newFileName += uuid.New().String() + extension

	err = c.SaveUploadedFile(uploadedPhoto, newFileName)
	if err != nil {
		c.String(http.StatusInternalServerError, "unknown error")
		return
	}

	photo := models.Photo{
		Path:  "http://127.0.0.1:4000/" + newFileName,
		Order: 0,
	}

	err = h.DB.Model(&car).Association("Photos").Append(&photo)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusCreated, photo)
}

// PhotosDelete deletes a single existing manufacturer
func (h *Handler) PhotosDelete(c *gin.Context) {
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

	var engine models.Photo

	result = h.DB.Where("car_id = ?", carID).Find(&engine)
	if result.Error != nil || result.RowsAffected == 0 {
		log.Println("Record not found", result.Error)
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Record not found",
		})
		return
	}

	result = h.DB.Delete(&engine)
	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted engine " + strconv.Itoa(int(engine.ID)),
	})
}

// PhotosUpdateOrders deletes a single existing manufacturer
func (h *Handler) PhotosUpdateOrders(c *gin.Context) {
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

	var engine models.Photo

	result = h.DB.Where("car_id = ?", carID).Find(&engine)
	if result.Error != nil || result.RowsAffected == 0 {
		log.Println("Record not found", result.Error)
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Record not found",
		})
		return
	}

	result = h.DB.Delete(&engine)
	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted engine " + strconv.Itoa(int(engine.ID)),
	})
}
