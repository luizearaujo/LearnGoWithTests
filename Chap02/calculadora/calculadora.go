package calculadora

const sinalSoma = "+"
const sinalSubtracao = "-"
const sinalMultiplicacao = "*"
const sinalDivisao = "/"

func Somar(x , y int) int {
	return x + y
}

func Subtrair(x , y int) int {
	return x - y
}

func Multiplicar(x, y int) int {
	return x * y
}

func Dividir(x, y int) int {
	return x / y
}

func Calcular(x, y int, sinal string) int {
	
	switch sinal {
		case sinalSubtracao:
			return Subtrair(x , y)
		case sinalMultiplicacao:
			return Multiplicar( x, y)
		case sinalDivisao:
			return Dividir(x , y)
		default:
			return Somar(x, y)
	}
	
}
