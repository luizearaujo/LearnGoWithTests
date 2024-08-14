package main

const (
	ErrorIndexNotFound = ErrorDicionario("Não foi possível encontrar a informação a partir do index.")
	ErrorIndexAlreadyExist = ErrorDicionario("Chave já existente no map.")
)

type ErrorDicionario string

func (e ErrorDicionario) Error() string {
	return string(e)
}
