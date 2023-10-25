package usecase

type UsecaseInterface interface {
	Execute(input any) (any, error) // executa o caso do usuario
	Compensate(input any) error     // compensa se algo da errado ( outro sistema )
}
