package database

import (
	"fmt"

	"github.com/isslerman/202308-CursoPosGoFullCycle/22-Interface/entity"
)

type OrderRepository struct {
	DB string
}

func (r *OrderRepository) Save(order *entity.Order) error {
	fmt.Println("Saving DB")
	return nil
}
