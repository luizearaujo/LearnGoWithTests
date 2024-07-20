package repeat

const quantidadeRepeticao = 5

func ImprimirCaracter(caracter string) string{
	
	var saida string
	
	for i := 0 ; i < quantidadeRepeticao ; i++ {
		saida += caracter
	}
	
	return saida
}
