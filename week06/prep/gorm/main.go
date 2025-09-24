package main

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Car struct {
	gorm.Model
	Make     string
	CarModel string
	Year     int
}

func main() {
	db, err := gorm.Open(sqlite.Open("demo.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	sqldb, err := db.DB()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer sqldb.Close()

	db.AutoMigrate(&Car{})

	accord := Car{
		Make:     "Honda",
		CarModel: "Accord",
		Year:     2025,
	}
	db.Create(&accord)
	fmt.Println("created accord ID: ", accord.ID)

	var foundCar Car
	db.First(&foundCar, accord.ID)
	fmt.Printf("found: %v\n", foundCar)

	camry := Car{
		Make:     "Toyota",
		CarModel: "Camry",
		Year:     2024,
	}
	db.Create(&camry)
	fmt.Println("created camry ID: ", camry.ID)

	var foundCars []Car
	db.Find(&foundCars, []int{1, 2})
	for _, car := range foundCars {
		fmt.Printf("Found car: %v\n", car)
	}

	db.Where("1=1").Delete(&Car{})
	db.Find(&foundCars, []int{1, 2})
	for _, car := range foundCars {
		fmt.Printf("Found car: %v\n", car)
	}

	recentCutoff := time.Now().AddDate(0, 0, -7)
	var recentlyDeleted []Car
	db.Unscoped().
		Where("deleted_at > ? AND deleted_at IS NOT NULL", recentCutoff).
		Find(&recentlyDeleted)
	for _, car := range recentlyDeleted {
		fmt.Printf("deleted car: %v\n", car)
	}
}
