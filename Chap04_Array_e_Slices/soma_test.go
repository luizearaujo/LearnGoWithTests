package soma

import "testing"

func TestSoma(t *testing.T) {
	
	numeros := [3]int{1,2,3}
	resultado := Soma(numeros)
	esperado := 6
	
	if resultado != esperado {
		t.Errorf("Resultado '%d', Esperado '%d', Dado '%v'", resultado, 
			esperado, numeros)
	}
	
}
