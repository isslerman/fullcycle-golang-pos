// Cria um template e printa no stdout o nosso struct passado. Usamos aqui o .Must do template, assim podemos checar por erros.
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
	// the func Must give us the possibility to execute and store an error if exist.
	t := template.Must(template.New("CursoTemplate").Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}"))
	err := t.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}
