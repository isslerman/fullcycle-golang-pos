package file

import (
	"fmt"

	"github.com/isslerman/202308-CursoPosGoFullCycle/22-Interface/DataBase-Project-Example/entity"
)

type OrderRepository struct {
	FileIO string
}

func (r *OrderRepository) Save(order *entity.Order) error {
	fmt.Println("Saving file")
	return nil
}
