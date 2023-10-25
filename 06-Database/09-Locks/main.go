// Lock otimista
// existe um version na linha da tabela e vemos se ainda estamos com a mesma versao no inicio e no fim do processo.

// Lock pessimista
// lock a linha da tabela até que o processo tenha sido completado.
// custo mais alto.

package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	db.AutoMigrate(&Product{}, &Category{})

	// lock pessimista // essa aula ficou bem por cima mostrando o conteudo, mas não mto clara.
	// fala da otimista e pessimista, mas aplica apenas a pessimista.
	// pesquisar qual nome em ingles.

	// Starting the tx
	tx := db.Begin()
	var c Category
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&c, 1).Error
	if err != nil {
		panic(err)
	}
	c.Name = "Eletronicos"
	tx.Debug().Save(&c)
	tx.Commit()

}
