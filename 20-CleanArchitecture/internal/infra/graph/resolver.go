package graph

import "github.com/isslerman/202308-CursoPosGoFullCycle/20-CleanArchitecture/internal/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CreateOrderUseCase usecase.CreateOrderUseCase
}
