package usecase

import (
	"github.com/isslerman/202308-CursoPosGoFullCycle/22-Interface/entity"
)

type CreateOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface // Interface para falar com o BD.
}

func (c *CreateOrderUseCase) Execute(o entity.Order) {
	o.CalculateFinalPrice()
	c.OrderRepository.Save(&o)
}
