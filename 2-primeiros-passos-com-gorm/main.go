package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
	gorm.Model
}

func main() {
	dsn := "root:@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Product{})

	// Create
	//db.Create(&Product{Name: "Notebook", Price: 1000.00})

	// Create batch
	//products := []Product{
	//	{Name: "Notebook2", Price: 1000.00},
	//	{Name: "Smartphone1", Price: 500.00},
	//	{Name: "Mouse", Price: 300.00},
	//}
	//db.Create(&products)

	// select one
	//var product Product
	//db.First(&product, 2)
	//fmt.Println(product)
	//db.First(&product, "name = ?", "Mouse")
	//fmt.Println(product)

	// select all
	//var products []Product
	//db.Find(&products)
	//for _, product := range products {
	//	fmt.Println(product)
	//}

	// Limit and Offset (pagination)
	//var products []Product
	//db.Limit(2).Offset(2).Find(&products)
	//for _, product := range products {
	//	fmt.Println(product)
	//}

	// Where
	//var products []Product
	//db.Where("price > ?", 500).Find(&products)
	//for _, product := range products {
	//	fmt.Println(product)
	//}

	// Like
	//var products []Product
	//db.Where("name LIKE ?", "%book%").Find(&products)
	//for _, product := range products {
	//	fmt.Println(product)
	//}

	// Update & Delete
	//var product Product
	//db.First(&product, 1)
	//product.Name = "New Mouse"
	//db.Save(&product)

	var product2 Product
	db.First(&product2, 1)
	fmt.Println(product2.Name)
	db.Delete(&product2) // Soft delete
}
