package math

// nomes em maiusculo são exportados
type Exported_Math struct {
	A int // aqui também vale a regra do maiusculo
	B int
}

type notexported_Math struct {
	A int // aqui também vale a regra do maiusculo
	B int
}

func (m Exported_Math) Add() int {
	return m.A + m.B
}

func (m notexported_Math) add() int {
	return m.A + m.B
}

func NewAdd(a int, b int) int {
	m := notexported_Math{A: a, B: b}
	return m.add()
}
