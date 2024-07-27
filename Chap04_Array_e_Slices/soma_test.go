package soma

import (
	"reflect"
	"testing"
)

func TestSoma(t *testing.T) {

	t.Run("coleção de 3 números", func(t *testing.T) {

		numeros := []int{1, 2, 3}
		resultado := Soma(numeros)
		esperado := 6

		if resultado != esperado {
			t.Errorf("Resultado '%d', Esperado '%d', Dado '%v'", resultado,
				esperado, numeros)
		}
	})

	t.Run("coleção com qualquer tamanho", func(t *testing.T) {
		lista := []int{1, 2, 3, 4}
		resultado := Soma(lista)
		esperado := 10

		if resultado != esperado {
			t.Errorf("Resultado '%d', Esperado '%d', Dados '%v'", resultado,
				esperado, lista)
		}
	})
}

func TestSomaTudo(t *testing.T) {

	resultado := SomaTudo([]int{1, 2}, []int{0, 9})
	esperado := []int{3, 9}

	if !reflect.DeepEqual(resultado, esperado) {
		t.Errorf("Resultado '%v', Esperado '%v'", resultado, esperado)
	}

}

func TestSomaCauda(t *testing.T) {
	
	verificarSomas := func(t *testing.T, resultado, esperado []int) {
		t.Helper()
		if !reflect.DeepEqual(resultado, esperado) {
			t.Errorf("Resultado '%v', Esperado '%v'", resultado, esperado)
		}
	}
	
	t.Run("Soma cauda de um slice de duas posições", func(t *testing.T) {
	
		resultado := SomaCauda([]int{1,2}, []int{0,9})
		esperado := []int{2,9}
		
		verificarSomas(t, resultado, esperado)
	})
	
	t.Run("Soma cauda de slices de diferentes tamanhos", func(t *testing.T) {
		
		resultado := SomaCauda([]int{1,2,3}, []int{1,1,1}, []int{2,9,7,8})
		esperado := []int{5,2,24}
		
		verificarSomas(t, resultado, esperado)
	})
	
	t.Run("Soma cauda de slices vazios", func(t *testing.T) {
		resultado := SomaCauda([]int{1,2}, []int{})
		esperado := []int{2,0}
		
		verificarSomas(t, resultado, esperado)
	})
}
