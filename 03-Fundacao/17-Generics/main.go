package main

// Aqui usamos o generics aonde indicamos que podemos receber na variavel T, o tipo int ou float64.
func Soma[T int | float64](m map[string]T) T {
	var Soma T
	for _, v := range m {
		Soma += v
	}
	return Soma
}

// Um segundo método é criar uma constrain
type Number interface {
	// para aceitarmos o type MyNumber que também é int, temos que aqui na declaração usar o ~
	~int | ~float64
}

type MyNumber int

// Aqui usamos o generics aonde indicamos que podemos receber na variavel T, o tipo int ou float64.
func Soma2[T Number](m map[string]T) T {
	var Soma T
	for _, v := range m {
		Soma += v
	}
	return Soma
}

func compare[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	// Aqui criamos duas variaveis mapeadas uma com int e outra com float e usamos a mesma funcao para somar.
	mapInt := map[string]int{"Marcao": 6000, "Boss": 10000, "Bekka": 14550, "Zé": 12}
	mapFloat := map[string]float64{"Marcao": 6000.12, "Boss": 10000.50, "Bekka": 14550.33, "Zé": 1.20}

	println(Soma(mapInt))
	println(Soma(mapFloat))

	println(Soma2(mapInt))
	println(Soma2(mapFloat))

	compare(10, 10)
}
