//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/isslerman/202308-CursoPosGoFullCycle/20-CleanArchitecture/internal/entity"
	"github.com/isslerman/202308-CursoPosGoFullCycle/20-CleanArchitecture/internal/event"
	"github.com/isslerman/202308-CursoPosGoFullCycle/20-CleanArchitecture/internal/infra/database"
	"github.com/isslerman/202308-CursoPosGoFullCycle/20-CleanArchitecture/internal/infra/web"
	"github.com/isslerman/202308-CursoPosGoFullCycle/20-CleanArchitecture/internal/usecase"
	"github.com/isslerman/202308-CursoPosGoFullCycle/20-CleanArchitecture/pkg/events"
)

var setOrderRepositoryDependency = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)),
)

var setEventDispatcherDependency = wire.NewSet(
	events.NewEventDispatcher,
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
)

var setOrderCreatedEvent = wire.NewSet(
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
)

func NewCreateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		usecase.NewCreateOrderUseCase,
	)
	return &usecase.CreateOrderUseCase{}
}

func NewWebOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebOrderHandler {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		web.NewWebOrderHandler,
	)
	return &web.WebOrderHandler{}
}
