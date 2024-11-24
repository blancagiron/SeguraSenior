package segurasenior

import "errors"

type FormatoFuente string

const (
	html  FormatoFuente = "html"
	pdf   FormatoFuente = "pdf"
	excel FormatoFuente = "excel"
)

type Fuente struct {
	nombre_fuente           string
	nombre_poblacion_fuente string
	formato_fuente          FormatoFuente
	direccion               string
}

func NewFuente(nom string, nompoblacion string, ext FormatoFuente, dir string) (*Fuente, error) {
	if nompoblacion == "" {
		return nil, errors.New("una fuente debe hablar de alguna población")
	}
	if ext != html && ext != pdf && ext != excel {
		return nil, errors.New("formato indefinido")
	}
	return &Fuente{
		nombre_fuente:           nom,
		nombre_poblacion_fuente: nompoblacion,
		formato_fuente:          ext,
		direccion:               dir,
	}, nil
}
