// Cria um template e printa no stdout usando um arquivo .html como template.
package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria interface{}
}

type Cursos []Curso

func main() {
	// the func Must give us the possibility to execute and store an error if exist.
	t := template.Must(template.New("template.html").ParseFiles("template.html"))
	err := t.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"Python", 20},
		{"C++", 140},
	})
	if err != nil {
		panic(err)
	}
}
