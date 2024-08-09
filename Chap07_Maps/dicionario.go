package main

import (
	"errors"
)

var ErrorIndexNotFound = errors.New("Não foi possível encontrar a informação a partir do index.")

type Dicionario map[string]string

func (d Dicionario) Busca(index string) (string, error) {
	definicao, existe := d[index]
	if !existe {
		return "", ErrorIndexNotFound
	}
	return definicao, nil
}
