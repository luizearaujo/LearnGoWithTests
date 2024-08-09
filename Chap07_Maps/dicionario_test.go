package main

import (
	"testing"
)

func TestBusca(t *testing.T) {
	
	dicionario := Dicionario{"teste":"isso é apenas um teste"}
		
	t.Run("chave existente", func(t *testing.T) {
		actual, _ := dicionario.Busca("teste")
		expected := "isso é apenas um teste"
	
		comparaString(t, actual, expected)	
	})
	
	t.Run("chave inexistente", func(t *testing.T) {
		_, err := dicionario.Busca("desconhecido")
		
		comparaString(t, err.Error(), ErrorIndexNotFound.Error())
	})
	
}

func comparaString(t *testing.T, actual, expected string) {
	t.Helper()
	if actual != expected {
		t.Errorf("Resultad '%s', Esperado '%s'", actual, expected)
	}
}
