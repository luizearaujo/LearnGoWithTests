package repeat



func ImprimirCaracter(caracter string, quantidadeRepeticao int) string{
	
	var saida string
	
	for i := 0 ; i < quantidadeRepeticao ; i++ {
		saida += caracter
	}
	
	return saida
}
