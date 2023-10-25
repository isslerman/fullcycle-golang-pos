package matematica

// se essa função for maiúscula ela está como Export, send acessivel externamente.
// se for minúscula é private para esse package.

// isso serve para types, funcs, vars.
func Soma[T int | float64](a, b T) T {
	return a + b
}

var VarAcessivelExternal = 10
var varNotAcessivelExternal = 20
