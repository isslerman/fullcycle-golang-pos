package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"` // uma categoria tem varios produtos.
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	Categories   []Category `gorm:"many2many:products_categories;"` // um produto pode ter varias categorias
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
	category1 := Category{Name: "Cozinha"}
	db.Create(&category1)
	category2 := Category{Name: "Eletronicos"}
	db.Create(&category2)

	// // create product
	db.Create(&Product{
		Name:       "Relogio",
		Price:      18.99,
		Categories: []Category{category1, category2}, // linkamos o produtos a duas categorias.
	})
}
