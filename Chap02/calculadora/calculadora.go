package calculadora

const sinalSoma = "+"
const sinalSubtracao = "-"
const sinalMultiplicacao = "*"
const sinalDivisao = "/"

func somar(x , y int) int {
	return x + y
}

func subtrair(x , y int) int {
	return x - y
}

func multiplicar(x, y int) int {
	return x * y
}

func dividir(x, y int) int {
	return x / y
}

func Calcular(x, y int, sinal string) int {
	
	switch sinal {
		case sinalSubtracao:
			return subtrair(x , y)
		case sinalMultiplicacao:
			return multiplicar( x, y)
		case sinalDivisao:
			return dividir(x , y)
		default:
			return somar(x, y)
	}
	
}
