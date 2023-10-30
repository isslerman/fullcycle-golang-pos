package model

type Category struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	// aqui removemos os cursos para deixar nos arquivos separados para o graphql reconhecer que os cursos tem uma categoria.
}
