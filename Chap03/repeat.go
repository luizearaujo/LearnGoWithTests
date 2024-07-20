package repeat

func ImprimirCaracter(caracter string) string{
	
	var saida string
	
	for i := 0 ; i < 5 ; i++ {
		saida = saida + caracter
	}
	
	return saida
}
