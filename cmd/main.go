package main

import (
	"cars-pet-project/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
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
}
