package calculadora

import "testing"

func TestSomar(t *testing.T) {
	resultado := Somar(1 , 2)
	esperado := 3
	
	if esperado != resultado {
		t.Errorf("esperado '%d', resultado '%d'", esperado, resultado)
	}
}

func TestSubtrair(t *testing.T) {
	resultado := Subtrair(2, 1)
	esperado := 1
	
	if esperado != resultado {
		t.Errorf("esperado '%d', resultado '%d'", esperado, resultado)
	}
}

func TestMultiplicar(t *testing.T) {
	resultado := Multiplicar(3 , 2)
	esperado := 6
	if esperado != resultado {
		t.Errorf("esperado '%d', resultado '%d'", esperado, resultado)
	}
}
	
