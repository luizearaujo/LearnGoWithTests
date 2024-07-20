package repeat

import "testing"

func TestImprimirCaracter(t *testing.T) {
	resultado := ImprimirCaracter("*", 5)
	esperado := "*****"
	
	if resultado != esperado {
		t.Errorf("Resultado '%s', Esperado '%s'", resultado, esperado)
	}
}

func TestImprimirCaracterInformado(t *testing.T) {
	resultado := ImprimirCaracter("a", 5)
	esperado := "aaaaa"
	
	if resultado != esperado {
		t.Errorf("Resultado '%s', Esperado '%s'", resultado, esperado)
	} 
}

func TestImprimirCaracterInformadoComQuantidade(t *testing.T) {
	resultado := ImprimirCaracter("-",10)
	esperado := "----------"

	if resultado != esperado {
		t.Errorf("Resultado '%s', Esperado '%s'", resultado, esperado)
	}
}

func BenchmarkImprimirCaracter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ImprimirCaracter("b", 3)
	}
}
