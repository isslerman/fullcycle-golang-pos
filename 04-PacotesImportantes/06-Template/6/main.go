// Cria um template e usa o http writer para dar a saida. webserver localhost 8282.
// Aqui temos a criação de uma função usada dentro do template.
package main

import (
	"html/template" // better to use html/template than text/template. More secure.
	"net/http"
	"strings"
)

type Curso struct {
	Nome         string
	CargaHoraria interface{}
}

type Cursos []Curso

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// content is our base template that include the other files.
		t := template.New("content.html")
		t.Funcs(template.FuncMap{"ToUpper": ToUpper})
		// t.Funcs(template.FuncMap{"ToUpper": strings.ToUpper })
		t = template.Must(t.ParseFiles(templates...))
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
