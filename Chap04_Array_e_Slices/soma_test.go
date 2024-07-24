package soma

import (
	"reflect"
	"testing"
)

func TestSomaComArray(t *testing.T) {

	numeros := []int{1, 2, 3}
	resultado := Soma(numeros)
	esperado := 6

	if resultado != esperado {
		t.Errorf("Resultado '%d', Esperado '%d', Dado '%v'", resultado,
			esperado, numeros)
	}

}

func TestaSomaComSlice(t *testing.T) {

	lista := []int{1, 2, 3, 4}
	resultado := Soma(lista)
	esperado := 10

	if resultado != esperado {
		t.Errorf("Resultado '%d', Esperado '%d', Dados '%v'", resultado,
			esperado, lista)
	}
}

func TestSomaTudo(t *testing.T) {

	resultado := SomaTudo([]int{1, 2}, []int{0, 9})
	esperado := []int{3, 9}

	if !reflect.DeepEqual(resultado, esperado) {
		t.Errorf("Resultado '%v', Esperado '%v'", resultado, esperado)
	}

}
