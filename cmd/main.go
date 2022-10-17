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
	}

	router := gin.Default()
	hndlrs := handlers.New(db)

	router.GET("/health-check", hndlrs.HealthCheck)
	manufacturers := router.Group("/manufacturers")
	{
		manufacturers.GET("/", hndlrs.ManufacturersIndex)
		manufacturers.POST("/", hndlrs.ManufacturersStore)
		manufacturers.GET("/:manufacturerID", hndlrs.ManufacturersShow)
		manufacturers.PUT("/:manufacturerID", hndlrs.ManufacturersUpdate)
		manufacturers.DELETE("/:manufacturerID", hndlrs.ManufacturersDelete)

		cars := manufacturers.Group("/:manufacturerID/cars")
		{
			cars.GET("/", hndlrs.CarsIndex)
			cars.POST("/", hndlrs.CarsStore)
			cars.GET("/:carID", hndlrs.CarsShow)
			cars.PUT("/:carID", hndlrs.CarsUpdate)
			cars.DELETE("/:carID", hndlrs.CarsDelete)

			engine := cars.Group("/:carID/engine")
			{
				engine.POST("/", hndlrs.EngineStore)
				engine.GET("/", hndlrs.EngineShow)
				engine.PUT("/", hndlrs.EngineUpdate)
				engine.DELETE("/", hndlrs.EngineDelete)
			}

			photos := cars.Group("/:carID/photos")
			{
				photos.GET("/", hndlrs.PhotosIndex)
				photos.POST("/", hndlrs.PhotosStore)
				photos.PUT("/orders", hndlrs.PhotosUpdateOrders)
				photos.DELETE("/:photoID", hndlrs.PhotosDelete)
			}
		}
	}

	log.Println("Starting server at :4000")
	router.Run(":4000")
}
