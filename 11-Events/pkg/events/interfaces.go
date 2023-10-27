package events

import (
	"sync"
	"time"
)

type EventInterface interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{}
}

type EventHandlerInterface interface {
	Handle(event EventInterface, wg *sync.WaitGroup)
}

type EventDispatcherInterface interface {
	Register(eventName string, handler EventHandlerInterface) // Registra
	Dispatch(event EventInterface)                            // Dispara
	Remove(eventName string, handler EventHandlerInterface)   // Remove
	Has(eventName string, handler EventHandlerInterface) bool // Check if exist
	Clear()                                                   // Limpa todos eventos
}
