// Cria um template e usa o http writer para dar a saida. webserver localhost 8282.
package main

import (
	"html/template" // better to use html/template than text/template. More secure.
	"net/http"
)

type Curso struct {
	Nome         string
	CargaHoraria interface{}
}

type Cursos []Curso

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// content is our base template that include the other files.
		t := template.Must(template.New("content.html").ParseFiles(templates...))
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
