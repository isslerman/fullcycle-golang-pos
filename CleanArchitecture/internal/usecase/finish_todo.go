package usecase

// sempre publicar um evento quando sua entidade mudar
// gerar logs para a mudança de estado

// DTO (Data Transfer Object)
type InputFinishTodo struct {
	ID string
}

type OutputFinishTodo struct {
	ID string
}

type InputCompensateFinishTodo struct {
	ID     string
	Reason string
}

// type FinishTodoUseCase struct {
// 	TodoRepository  TodoRepository
// 	CompensateEvent CompensateEvent
// }

// func (f *FinishTodoUseCase) Execute(input any) (any, error) {

// }

// func (f *FinishTodoUseCase) Compensate(input any) error {

// }
