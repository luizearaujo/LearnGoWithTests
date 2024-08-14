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
		
		comparaErro(t, err, ErrorIndexNotFound)
	})
	
}

func TestAdicionar(t *testing.T) {
	
		
	t.Run("Adiciona nova palavra", func(t *testing.T) {
		dicionario := Dicionario{}
			
		key := "teste"
		value := "isso é apenas um teste"
		
		err := dicionario.Adiciona(key, value)
		
		comparaErro(t, err, nil)
		comparaDefinicao(t, dicionario, "teste", "isso é apenas um teste")
	})
	
	t.Run("Adiciona palavra já existente", func(t *testing.T){
		key := "teste"
		value := "isso é apenas um teste"
		dicionario := Dicionario{key: value}
		
		err := dicionario.Adiciona(key, "Esta chave já existe")
		
		comparaErro(t, err, ErrorIndexAlreadyExist)
		comparaDefinicao(t, dicionario, key, value)
	})
	
}

func comparaDefinicao(t *testing.T, dicionario Dicionario, palavra, definicao string){

	t.Helper()
	
	resultado, err := dicionario.Busca(palavra)
	if err != nil {
		t.Fatal("Deveria ter encontrado uma palavra adicionada: ", err)
	}
	
	if definicao != resultado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, definicao)
	}
	
}

func comparaString(t *testing.T, actual, expected string) {
	t.Helper()
	if actual != expected {
		t.Errorf("Resultado '%s', Esperado '%s'", actual, expected)
	}
}

func comparaErro(t *testing.T, actual, expected error){
	t.Helper()
	if actual != expected {
		t.Errorf("Resultado '%s', Esperado '%s'", actual, expected)
	}
}
