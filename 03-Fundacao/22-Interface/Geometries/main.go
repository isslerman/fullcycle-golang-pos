package main

import "fmt"

type rect struct {
	width  int
	height int
	name   string
}

type circle struct {
	radius int
}

type geometry interface {
	area() int
}

func (c circle) area() int {
	return 10
}

func (r rect) area() int {
	return 1
}

func (r rect) String() string {
	return r.name
}

func main() {
	r := rect{width: 10, height: 10}
	r.name = "retangulo"
	c := circle{radius: 15}

	objList := []geometry{r, c}

	fmt.Println("A area do rect é:", r.area())
	fmt.Println("A area do circle é:", c.area())

	fmt.Println("A area do rect é:", area(c))
	fmt.Println("A area do circle é:", area(r))

	PrintAllAreas(objList)

	fmt.Println("Vamos testar aqui essa loucura:", r)
}

func area(g geometry) int {
	return g.area()
}

func PrintAllAreas(gs []geometry) {
	for i, g := range gs {
		fmt.Println("A area do elemento", i, "é:", g.area())
	}
}
