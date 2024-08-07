package main

import (
	"testing"
)

func TestBusca(t *testing.T) {
	
	dicionario := map[string]string{"teste":"isso é apenas um teste"}
	
	actual := Busca(dicionario, "teste")
	expected := "isso é apenas um teste"
	
	comparaString(t, actual, expected)
}

func comparaString(t *testing.T, actual, expected string) {
	t.Helper()
	if actual != expected {
		t.Errorf("Resultad '%s', Esperado '%s'", actual, expected)
	}
}
