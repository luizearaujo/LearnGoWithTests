package soma

func Soma(lista []int) int {

	resultado := 0
	for _, numero := range lista {
		resultado += numero
	}
	return resultado
}

func SomaTudo(numerosParaSomar ...[]int) []int {

	var somas []int
	for _, numeros := range numerosParaSomar {
		somas = append(somas, Soma(numeros))
	}

	return somas
}
