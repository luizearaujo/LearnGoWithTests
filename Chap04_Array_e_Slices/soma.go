package soma

func Soma(lista []int) int {

	resultado := 0
	for _, numero := range lista {
		resultado += numero
	}
	return resultado
}

func somaCauda(lista []int) int {
	resultado := 0
	for i:=1; i < len(lista); i++ {
		resultado += lista[i]
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

func SomaCauda(numerosParaSomar ...[]int) []int {
	var somas []int
	
	for _, numeros := range numerosParaSomar {
		somas = append(somas, somaCauda(numeros))
	}
	
	return somas
}
