package usecase

import (
	"github.com/isslerman/202308-CursoPosGoFullCycle/20-CleanArchitecture/internal/entity"
	"github.com/isslerman/202308-CursoPosGoFullCycle/20-CleanArchitecture/pkg/events"
)

// DTO - entradas e saida de dados
type OrderInputDTO struct {
	ID    string  `json:"id"`
	Price float64 `json:"price"`
	Tax   float64 `json:"tax"`
}

type OrderOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

// Os componentes do nosso use case.
type CreateOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface // Interface para falar com o BD.
	OrderCreated    events.EventInterface           // Interface de ordem criada - evento.
	EventDispatcher events.EventDispatcherInterface // Interface para disparo do evento.
}

// Criação dos componentes / usecase
func NewCreateOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	OrderCreated events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		OrderRepository: OrderRepository,
		OrderCreated:    OrderCreated,
		EventDispatcher: EventDispatcher,
	}
}

// Método principal do usecase.
// Recebe o inputDTO, retorna
func (c *CreateOrderUseCase) Execute(input OrderInputDTO) (OrderOutputDTO, error) {
	// cria a ordem usando os dados recebidos pela inputDTO
	order := entity.Order{
		ID:    input.ID,
		Price: input.Price,
		Tax:   input.Tax,
	}
	// tendo a ordem criada, calcula o preço final.
	order.CalculateFinalPrice()
	// usa o repositorio para salvar em banco. Não importa se é BD, file, ele está sendo salvo.
	// se existir um erro, retorno o outputDTO em branco + erro.
	if err := c.OrderRepository.Save(&order); err != nil {
		return OrderOutputDTO{}, err
	}

	// prepara o DTO de output.
	dto := OrderOutputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}

	// prepara o evento
	c.OrderCreated.SetPayload(dto) // interface
	// dispara o evento
	c.EventDispatcher.Dispatch(c.OrderCreated) // interface

	// retorna o outputDTO
	return dto, nil
}
