package main

import "testing"


func TestOla(t *testing.T){
	
	verificarMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	}
	
	t.Run("Diz olá para as pessoas", func (t *testing.T){
	
		resultado := Ola("Ben", "")
		esperado := "Olá, Ben"
	
		verificarMensagemCorreta(t, resultado, esperado)
	})	 
	
	t.Run("Diz 'Olá, mundo' quando uma string vazia for passada", 
	func(t *testing.T){
		resultado := Ola("", "")
		esperado := "Olá, mundo"
		
		verificarMensagemCorreta(t, resultado, esperado)
	})
	
	t.Run("Em espanhol", func(t *testing.T){
		
		resultado := Ola("Elodie", "ES")
		esperado := "Hola, Elodie"
		
		verificarMensagemCorreta(t, resultado, esperado)
	})
	
	t.Run("Em frances", func(t *testing.T){
		
		resultado := Ola("Pierre", "FR")
		esperado := "Bonjour, Pierre"
		
		verificarMensagemCorreta(t, resultado, esperado)
	})
	
	t.Run("Em alemao", func(t *testing.T){
		
		resultado := Ola("Michael", "GR")
		esperado := "Hallo, Michael"
		
		verificarMensagemCorreta(t, resultado, esperado)
	})
	
	t.Run("Em ingles", func(t *testing.T){
		
		resultado := Ola("Adam", "EN")
		esperado := "Hello, Adam"
		
		verificarMensagemCorreta(t, resultado, esperado)
	})
		
}
