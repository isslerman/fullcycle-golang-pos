// Cria um template e printa no stdout o nosso struct.
package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria interface{}
}

func main() {
	curso := Curso{"Go", 40}
	tmp := template.New("CursoTemplate")
	tmp, _ = tmp.Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}")
	err := tmp.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}
