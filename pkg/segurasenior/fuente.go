package segurasenior

import "errors"

type FormatoFuente string

const (
	html  FormatoFuente = "html"
	pdf   FormatoFuente = "pdf"
	excel FormatoFuente = "excel"
)

type Fuente struct {
	nombre_fuente  string
	formato_fuente FormatoFuente
	direccion      string
}

func NewFuente(nom string, ext FormatoFuente, dir string) (*Fuente, error) {
	if ext != html && ext != pdf && ext != excel {
		return nil, errors.New("formato indefinido")
	}
	return &Fuente{
		nombre_fuente:  nom,
		formato_fuente: ext,
		direccion:      dir,
	}, nil
}
