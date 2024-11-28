package segurasenior

import "errors"

type Fuente struct {
	Direccion string
}

func NewFuente(dir string) (*Fuente, error) {
	if nompoblacion == "" {
		return nil, errors.New("una fuente debe hablar de alguna población")
	}
	return &Fuente{
		Direccion: dir,
	}, nil
}
