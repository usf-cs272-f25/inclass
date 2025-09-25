package main

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Car struct {
	ID       int `gorm:"primaryKey"`
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

	// Creates the schema for the types of args passed
	db.AutoMigrate(&Car{}, &Word{}, &Doc{})

	accord := Car{
		Make:     "Honda",
		CarModel: "Accord",
		Year:     2025,
	}
	db.Create(&accord)
	fmt.Println("created accord ID: ", accord.ID)

	var foundCar Car
	// SELECT * FROM cars WHERE id = ?
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
	// SELECT * FROM cars where id = ?, ?
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
