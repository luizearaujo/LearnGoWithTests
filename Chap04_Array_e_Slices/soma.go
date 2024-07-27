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

func SomaCauda(numerosParaSomar ...[]int) []int {
	var somas []int
	
	for _, numeros := range numerosParaSomar {
		if len(numeros) == 0 {
			somas = append(somas, 0)
		} else {
			final := numeros[1:]
			somas = append(somas, Soma(final))
		}
	}
	
	return somas
}
