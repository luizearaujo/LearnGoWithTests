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
	
	resultado := SomaCauda([]int{1,2}, []int{0,9})
	esperado := []int{2,9}
	
	if !reflect.DeepEqual(resultado, esperado) {
		t.Errorf("Resultado '%v', Esperado '%v'", resultado, esperado)
	}
}
