package service

import (
	"context"

	"github.com/isslerman/202308-CursoPosGoFullCycle/20-CleanArchitecture/internal/infra/grpc/pb"
	"github.com/isslerman/202308-CursoPosGoFullCycle/20-CleanArchitecture/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	GetOrdersUseCase   usecase.GetOrdersUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase, getOrdersUseCase usecase.GetOrdersUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
		GetOrdersUseCase:   getOrdersUseCase,
	}
}

// func GetOrdersService(getOrdersUseCase usecase.GetOrdersUseCase) *OrderService{
// 	return &OrderService{
// 		GetOrdersUseCase: getOrdersUseCase,
// 	}
// }

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) GetOrders(ctx context.Context, in *pb.Blank) (*pb.OrderList, error) {
	orders, err := s.GetOrdersUseCase.Execute()
	if err != nil {
		return nil, err
	}
	var ordersResponse []*pb.Order

	for _, order := range orders.Orders {
		orderResponse := &pb.Order{
			Id:         order.ID,
			Price:      float32(order.Price),
			Tax:        float32(order.Price),
			FinalPrice: float32(order.FinalPrice),
		}
		ordersResponse = append(ordersResponse, orderResponse)
	}
	return &pb.OrderList{
		Orders: ordersResponse,
	}, nil
}
