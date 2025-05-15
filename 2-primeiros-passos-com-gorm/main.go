package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

func main() {
	dsn := "root:@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Product{})

	// Create
	//db.Create(&Product{Name: "Notebook", Price: 1000.00})

	// Create batch
	products := []Product{
		{Name: "Notebook2", Price: 1000.00},
		{Name: "Smartphone1", Price: 500.00},
		{Name: "Mouse", Price: 300.00},
	}
	db.Create(&products)
}
