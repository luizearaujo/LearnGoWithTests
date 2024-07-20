package repeat

import "testing"

func TestImprimirCaracter(t *testing.T) {
	resultado := ImprimirCaracter("*")
	esperado := "*****"
	
	if resultado != esperado {
		t.Errorf("Resultado '%s', Esperado '%s'", resultado, esperado)
	}
}

func TestImprimirCaracterInformado(t *testing.T) {
	resultado := ImprimirCaracter("a")
	esperado := "aaaaa"
	
	if resultado != esperado {
		t.Errorf("Resultado '%s', Esperado '%s'", resultado, esperado)
	} 
}
