package calculadora

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
	
	if sinal == "-" {
		return Subtrair(x ,y)
	}
	if sinal == "*" {
		return Multiplicar(x ,y)
	}
	if sinal == "/" {
		return Dividir(x , y)
	}
	return Somar(x ,y)
}
