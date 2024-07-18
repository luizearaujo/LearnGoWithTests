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
	
	if sinal == sinalSubtracao {
		return Subtrair(x ,y)
	}
	if sinal == sinalMultiplicacao {
		return Multiplicar(x ,y)
	}
	if sinal == sinalDivisao {
		return Dividir(x , y)
	}
	
	return Somar(x ,y)
}
