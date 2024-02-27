package main

import (
	"github.com/isslerman/202308-CursoPosGoFullCycle/22-Interface/DataBase-Project-Example/database"
	"github.com/isslerman/202308-CursoPosGoFullCycle/22-Interface/DataBase-Project-Example/entity"
	"github.com/isslerman/202308-CursoPosGoFullCycle/22-Interface/DataBase-Project-Example/file"
)

func main() {
	bd := database.OrderRepository{
		DB: "config",
	}
	file := file.OrderRepository{
		FileIO: "config",
	}

	order := &entity.Order{
		ID:    "id",
		Price: 10.0,
		Tax:   5.0,
	}
	order.CalculateFinalPrice()
	bd.Save(order)
	file.Save(order)
}
