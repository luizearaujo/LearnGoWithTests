package main

import (
	"testing"
)

func TestBusca(t *testing.T) {
	
	dicionario := map[string]string{"teste":"isso é apenas um teste"}
	
	actual := Busca(dicionario, "teste")
	expected := "isso é apenas um teste"
	
	if actual != expected {
		t.Errorf("Resultado '%s', Esperado '%s'", actual, expected)
	}
}
