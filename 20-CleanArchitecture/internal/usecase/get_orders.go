package usecase

import (
	"github.com/isslerman/202308-CursoPosGoFullCycle/20-CleanArchitecture/internal/entity"
	"github.com/isslerman/202308-CursoPosGoFullCycle/20-CleanArchitecture/pkg/events"
)

// DTO - entradas e saida de dados
type OrdersInputDTO struct{}

type OrdersOutputDTO struct {
	Orders []entity.Order
}

// Os componentes do nosso use case.
type GetOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface // Interface para falar com o BD.
	OrderCreated    events.EventInterface           // Interface de ordem criada - evento.
	EventDispatcher events.EventDispatcherInterface // Interface para disparo do evento.
}

// Criação dos componentes / usecase
func NewGetOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	OrderCreated events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *GetOrdersUseCase {
	return &GetOrdersUseCase{
		OrderRepository: OrderRepository,
		OrderCreated:    OrderCreated,
		EventDispatcher: EventDispatcher,
	}
}

// Método principal do usecase.
// Recebe o inputDTO, retorna outputDTO
func (c *GetOrdersUseCase) Execute() (OrdersOutputDTO, error) {
	// cria a ordem usando os dados recebidos pela inputDTO
	orders, err := c.OrderRepository.FindAll()
	if err != nil {
		return OrdersOutputDTO{}, err
	}

	// prepara o DTO de output.
	dto := OrdersOutputDTO{
		Orders: orders,
	}

	// prepara o evento
	c.OrderCreated.SetPayload(dto) // interface
	// dispara o evento
	c.EventDispatcher.Dispatch(c.OrderCreated) // interface

	// retorna o outputDTO
	return dto, nil
}
