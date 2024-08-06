package stock_broker

import (
	"testing"
)

func TestCarteira(t *testing.T) {

	carteira := Carteira{}
	carteira.Depositar(10)

	resultado := carteira.Saldo()
	esperado := Bitcoin(10)

	if resultado != esperado {
		t.Errorf("Resultado '%d', Esperado '%d'", resultado, esperado)
	}

}
