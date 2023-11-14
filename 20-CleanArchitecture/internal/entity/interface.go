package entity

type OrderRepositoryInterface interface {
	FindAll() ([]Order, error)
	Save(order *Order) error
}
