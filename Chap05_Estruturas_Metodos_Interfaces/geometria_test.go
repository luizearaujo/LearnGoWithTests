package geometria

import (
	"testing"
)

func TestPerimetro(t *testing.T) {
	retangulo := Retangulo{10.0, 10.0}
	resultado := retangulo.Perimetro()
	esperado := 40.0
	
	if resultado != esperado {
		t.Errorf("Resultado '%.2f, Esperado '%.2f'", resultado, esperado)
	}
}

func TestArea(t *testing.T) {
	
	t.Run("Retangulo", func(t *testing.T) {
		retangulo := Retangulo{12.0, 6.0}
		resultado := retangulo.Area()
		esperado := 72.0
		
		if resultado != esperado {
			t.Errorf("Resultado '%.2f', Esperado '%.2f'", resultado, esperado)
		}	
	})
	
	t.Run("Circulo", func(t *testing.T) {
		circulo := Circulo{10.0}
		resultado := circulo.Area()
		esperado := 314.1592653589793
		
		if resultado != esperado {
			t.Errorf("Resultado '%.2f', Esperado '%.2f'", resultado, esperado)
		}
	})
}
