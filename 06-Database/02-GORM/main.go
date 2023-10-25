package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	gorm.Model // soft-delete - create the created_at, updated_at, deleted_at
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	// creating one product
	// db.Create(&Product{
	// 	Name:  "Notebook",
	// 	Price: 999.99,
	// })

	// // creating many products with a slice of products.
	// products := []Product{
	// 	{Name: "Notebook", Price: 500.00},
	// 	{Name: "Monitor", Price: 999.33},
	// }
	// db.Create(&products)

	// var products []Product
	// getting the first, second line
	// db.First(&product, 1)
	// fmt.Println(product)

	// query
	// db.First(&products, "name=?", "Monitor")
	// fmt.Println(products)

	// select all with limit = 2 e offset pag = 2
	// db.Limit(2).Offset(2).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// WHERE
	// db.Where("price > ?", 500).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// LIKE
	// db.Where("name LIKE ?", "%book%").Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// var p Product
	// db.First(&p, 1)
	// p.Name = "Nem Mouse"
	// db.Save(&p)

	var p2 Product
	db.First(&p2, 1)
	fmt.Println(p2.Name)
	db.Delete(&p2)
}
