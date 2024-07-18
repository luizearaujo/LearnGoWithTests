package calculadora

import "testing"

func TestSomar(t *testing.T) {
	resultado := Somar(1 , 2)
	esperado := 3
	
	verificarResultadoTeste(t, resultado, esperado)
}

func TestSubtrair(t *testing.T) {
	resultado := Subtrair(2, 1)
	esperado := 1
	
	verificarResultadoTeste(t, resultado, esperado)
}

func TestMultiplicar(t *testing.T) {
	resultado := Multiplicar(3 , 2)
	esperado := 6
	
	verificarResultadoTeste(t, resultado, esperado)
}

func TestDividir(t *testing.T) {
	resultado := Dividir(10, 5)
	esperado := 2
	
	verificarResultadoTeste(t, resultado, esperado)
}

func TestCalcularSoma(t *testing.T) {
	resultado := Calcular(2, 2, "+")
	esperado := 4
	
	verificarResultadoTeste(t, resultado, esperado)
}

func TestCalcularSubtracao(t *testing.T) {
	resultado := Calcular(3, 1, "-")
	esperado := 2

	verificarResultadoTeste(t, resultado, esperado)
}

func TestCalcularMultiplicacao(t *testing.T) {
	resultado := Calcular(5 , 2, "*")
	esperado := 10
	verificarResultadoTeste(t, resultado, esperado)
}

func verificarResultadoTeste(t *testing.T, resultado, esperado int) {
	t.Helper()
	
	if esperado!=resultado {
		t.Errorf("esperado '%d', resultado '%d'", esperado, resultado)
	}
}
	
