package main

import (
	"cars-pet-project/pkg/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	dsn := "host=localhost" +
		" user=developer" +
		" password=developer" +
		" dbname=cars_pet_project_db" +
		" port=5432" +
		" sslmode=disable" +
		" TimeZone=Asia/Dushanbe"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
		return
	}

	err = db.AutoMigrate(
		&models.Manufacturer{},
		&models.Car{},
		&models.Photo{},
		&models.Engine{},
	)
	if err != nil {
		log.Fatalln(err)
		return
	}

	router := gin.Default()

	router.GET("/health-check", HealthCheck)
	manufacturers := router.Group("/manufacturers")
	{
		manufacturers.GET("/", ManufacturersIndex)
		manufacturers.POST("/", ManufacturersStore)
		manufacturers.GET("/:id", ManufacturersShow)
		manufacturers.PUT("/:id", ManufacturersUpdate)
		manufacturers.DELETE("/:id", ManufacturersDelete)
	}

	log.Println("Starting server at :4000")
	router.Run(":4000")
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "all is OK",
	})
}

func ManufacturersIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Manufacturers index",
	})
}

func ManufacturersStore(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Manufacturers store",
	})
}

func ManufacturersShow(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Manufacturers show",
	})
}

func ManufacturersUpdate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Manufacturers update",
	})
}

func ManufacturersDelete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Manufacturers delete",
	})
}
