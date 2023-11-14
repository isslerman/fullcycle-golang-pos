package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/isslerman/202308-CursoPosGoFullCycle/20-CleanArchitecture/internal/entity"
	"github.com/isslerman/202308-CursoPosGoFullCycle/20-CleanArchitecture/internal/usecase"
	"github.com/isslerman/202308-CursoPosGoFullCycle/20-CleanArchitecture/pkg/events"
)

type WebOrderHandler struct {
	EventDispatcher   events.EventDispatcherInterface
	OrderRepository   entity.OrderRepositoryInterface
	OrderCreatedEvent events.EventInterface
}

func NewWebOrderHandler(
	EventDispatcher events.EventDispatcherInterface,
	OrderRepository entity.OrderRepositoryInterface,
	OrderCreatedEvent events.EventInterface,
) *WebOrderHandler {
	return &WebOrderHandler{
		EventDispatcher:   EventDispatcher,
		OrderRepository:   OrderRepository,
		OrderCreatedEvent: OrderCreatedEvent,
	}
}

func (h *WebOrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto usecase.OrderInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createOrder := usecase.NewCreateOrderUseCase(h.OrderRepository, h.OrderCreatedEvent, h.EventDispatcher)
	output, err := createOrder.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *WebOrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Debug: GetOrdersHandler")
	getOrders := usecase.NewGetOrdersUseCase(h.OrderRepository, h.OrderCreatedEvent, h.EventDispatcher)
	outputDTO, err := getOrders.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Printf("Debug: GetOrdersHandler - output %v", outputDTO)

	err = json.NewEncoder(w).Encode(outputDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
