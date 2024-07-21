package main

import "fmt"

const ES = "ES"
const FR = "FR"
const GR = "GR"
const EN = "EN"

const prefixoOlaPTBR = "Olá, "
const prefixoOlaES = "Hola, "  
const prefixoOlaFR = "Bonjour, "
const prefixoOlaGR = "Hallo, "
const prefixoOlaEN = "Hello, "

func Ola(nome, idioma string) string {
	if nome == ""{
		nome = "mundo"
	}
		
	return prefixoSaudacao(idioma) + nome
}

func prefixoSaudacao(idioma string)(prefixo string){
	
	switch idioma {
		case FR:
			prefixo = prefixoOlaFR
		case ES:
			prefixo = prefixoOlaES
		case GR:
			prefixo = prefixoOlaGR
		case EN:
			prefixo = prefixoOlaEN
		default:
			prefixo = prefixoOlaPTBR
	}
		
	return
	}

func main(){
	fmt.Println(Ola("mundo", "PT"))
}
