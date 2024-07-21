package soma

func Soma(lista []int) int {
	
	resultado := 0
	for _,numero := range lista {
		resultado += numero
	}
	return resultado
}
