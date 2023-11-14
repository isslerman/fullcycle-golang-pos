package main

import (
	"github.com/isslerman/202308-CursoPosGoFullCycle/22-Interface/database"
	"github.com/isslerman/202308-CursoPosGoFullCycle/22-Interface/entity"
	"github.com/isslerman/202308-CursoPosGoFullCycle/22-Interface/file"
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
