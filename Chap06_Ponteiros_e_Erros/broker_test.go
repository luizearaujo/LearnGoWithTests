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

	confirmarSaldo := func(t *testing.T, carteira Carteira, esperado Bitcoin) {
		t.Helper()

		resultado := carteira.Saldo()

		if resultado != esperado {
			t.Errorf("Resultado '%d', Esperado '%d'", resultado, esperado)
		}
	}

	t.Run("depositar", func(t *testing.T) {

		carteira := Carteira{}
		carteira.Depositar(Bitcoin(10))

		esperado := Bitcoin(10)

		confirmarSaldo(t, carteira, esperado)
	})

	t.Run("retirar", func(t *testing.T) {

		carteira := Carteira{}
		carteira.Depositar(Bitcoin(10))
		carteira.Retirar(Bitcoin(5))

		esperado := Bitcoin(5)

		confirmarSaldo(t, carteira, esperado)

	})

	confirmarErro := func(t *testing.T, erro error, esperado string) {
		t.Helper()
		if erro == nil {
			t.Fatal("Esperava um erro, mas nenhum ocorreu.")
		}
		if erro.Error() != esperado {
			t.Errorf("Resultado '%s', Esperado '%s'", erro, esperado)
		}

	}

	t.Run("retirar com saldo insuficiente", func(t *testing.T) {

		saldoInicial := Bitcoin(20)
		carteira := Carteira{saldoInicial}
		erro := carteira.Retirar(Bitcoin(100))

		confirmarSaldo(t, carteira, saldoInicial)

		confirmarErro(t, erro, "Saldo insuficiente")
	})

}
