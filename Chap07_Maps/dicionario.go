package main

import (
	
)

type Dicionario map[string]string

func (d Dicionario) Busca(index string) string {
	return d[index]
}
