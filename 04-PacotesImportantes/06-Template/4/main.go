// Cria um template e usa o http writer para dar a saida. webserver localhost 8282.
package main

import (
	"html/template"
	"net/http"
)

type Curso struct {
	Nome         string
	CargaHoraria interface{}
}

type Cursos []Curso

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("template.html").ParseFiles("template.html"))
		err := t.Execute(w, Cursos{
			{"Go", 40},
			{"Python", 20},
			{"C++", 140},
		})
		if err != nil {
			panic(err)
		}

	})
	http.ListenAndServe(":8282", nil)
}
