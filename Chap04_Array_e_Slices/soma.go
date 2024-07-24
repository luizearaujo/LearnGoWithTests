package soma

func Soma(lista []int) int {

	resultado := 0
	for _, numero := range lista {
		resultado += numero
	}
	return resultado
}

func SomaTudo(numerosParaSomar ...[]int) (somas []int) {
	quantidadeDeNumeros := len(numerosParaSomar)
	somas = make([]int, quantidadeDeNumeros)

	for i, numeros := range numerosParaSomar {
		somas[i] = Soma(numeros)
	}

	return
}
