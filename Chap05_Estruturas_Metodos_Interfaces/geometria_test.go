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
	
	verificarArea := func(t *testing.T, forma Forma, esperado float64) {
		
		t.Helper()
		resultado := forma.Area()
		
		if resultado != esperado {
			t.Errorf("Resultado '%.2f', Esperado '%.2f'", resultado, esperado)
		}
	}
	
	
	t.Run("Retangulo", func(t *testing.T) {
		retangulo := Retangulo{12.0, 6.0}
		esperado := 72.0
		
		verificarArea(t, retangulo, esperado)	
	})
	
	t.Run("Circulo", func(t *testing.T) {
		circulo := Circulo{10.0}
		esperado := 314.1592653589793
		
		verificarArea(t, circulo, esperado)
	})
}

func TestAreaTable(t *testing.T) {
	
	testesArea := []struct {
		forma Forma
		esperado float64
	}{
		{forma: Retangulo{Largura: 12.0, Altura: 6.0}, esperado: 72.0},
		{forma: Circulo{Raio: 10.0}, esperado: 314.1592653589793},
		{forma: Triangulo{Altura: 12, Base: 6}, esperado: 36.0}, 
	}
	
	for _, tt := range testesArea {
		
		resultado := tt.forma.Area()
		
		if resultado != tt.esperado {
			t.Errorf("'%#v' - Esperado '%.2f', Resultado '%.2f'", tt.forma, tt.esperado, resultado)
		}
	} 
}
