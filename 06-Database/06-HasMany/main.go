package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model   // soft-delete - create the created_at, updated_at, deleted_at
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// create category
	category := Category{Name: "Eletronicos"}
	db.Create(&category)
	category = Category{Name: "Privadas"}
	db.Create(&category)

	// // create product
	db.Create(&Product{
		Name:       "Mouse",
		Price:      38.99,
		CategoryID: 1,
	})
	db.Create(&Product{
		Name:       "Assento",
		Price:      1.99,
		CategoryID: 2,
	})
	db.Create(&Product{
		Name:       "Descarga",
		Price:      8.99,
		CategoryID: 2,
	})

	var products []Product
	// preload with belongs to
	db.Preload("Category").Find(&products)
	for _, product := range products {
		fmt.Println(product.Name, product.Category.Name, product.SerialNumber.Number)
	}

	var categories []Category
	// obs for sub relation between category and serial number that was inside products.
	err = db.Model(&Category{}).Preload("Products").Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			fmt.Println("- ", product.Name, category.Name)
		}
	}

}
