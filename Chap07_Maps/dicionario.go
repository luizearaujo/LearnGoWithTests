package main


type Dicionario map[string]string

func (d Dicionario) Busca(key string) (string, error) {
	definicao, existe := d[key]
	if !existe {
		return "", ErrorIndexNotFound
	}
	return definicao, nil
}

func (d Dicionario) Adiciona(key, value string) error {
	
	_, err := d.Busca(key)
	
	switch err {
		case ErrorIndexNotFound:
			d[key] = value
		case nil :
			return ErrorIndexAlreadyExist
		default :
			return err
	}
	
	return nil
}
