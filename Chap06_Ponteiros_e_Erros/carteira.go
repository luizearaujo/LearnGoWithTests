package stock_broker

import "errors"

type Carteira struct {
	saldo Bitcoin
}

func (c *Carteira) Depositar(quantidade Bitcoin) {
	c.saldo += quantidade
}

func (c *Carteira) Saldo() Bitcoin {
	return c.saldo
}

func (c *Carteira) Retirar(quantidade Bitcoin) error {

	if quantidade > c.saldo {
		return errors.New("Saldo insuficiente")
	}

	c.saldo -= quantidade
	return nil
}
