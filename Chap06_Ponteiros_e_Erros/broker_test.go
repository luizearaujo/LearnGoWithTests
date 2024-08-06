package stock_broker

import (
	"testing"
)

func TestCarteira(t *testing.T) {

	t.Run("formatar", func(t *testing.T) {
		carteira := Carteira{}
		carteira.Depositar(Bitcoin(10))

		resultado := carteira.Saldo().String()
		esperado := "10 BTC"

		if resultado != esperado {
			t.Errorf("Resultado '%s', Esperado '%s'", resultado, esperado)
		}
	})

	t.Run("depositar", func(t *testing.T) {

		carteira := Carteira{}
		carteira.Depositar(Bitcoin(10))

		resultado := carteira.Saldo()
		esperado := Bitcoin(10)

		if resultado != esperado {
			t.Errorf("Resultado '%d', Esperado '%d'", resultado, esperado)
		}
	})

}
