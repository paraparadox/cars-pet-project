package main

import (
	"cars-pet-project/pkg/handlers"
	"cars-pet-project/pkg/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func DBInit() (*gorm.DB, error) {
	dsn := "host=localhost" +
		" user=developer" +
		" password=developer" +
		" dbname=cars_pet_project_db" +
		" port=5432" +
		" sslmode=disable" +
		" TimeZone=Asia/Dushanbe"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&models.Manufacturer{},
		&models.Car{},
		&models.Photo{},
		&models.Engine{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	db, err := DBInit()
	if err != nil {
		log.Fatalln(err)
		return
	}

	router := gin.Default()
	hndlrs := handlers.New(db)

	router.GET("/health-check", hndlrs.HealthCheck)
	manufacturers := router.Group("/manufacturers")
	{
		manufacturers.GET("/", hndlrs.ManufacturersIndex)
		manufacturers.POST("/", hndlrs.ManufacturersStore)
		manufacturers.GET("/:id", hndlrs.ManufacturersShow)
		manufacturers.PUT("/:id", hndlrs.ManufacturersUpdate)
		manufacturers.DELETE("/:id", hndlrs.ManufacturersDelete)
	}

	log.Println("Starting server at :4000")
	router.Run(":4000")
}
