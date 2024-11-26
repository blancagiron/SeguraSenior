package segurasenior

import "errors"

type Fuente struct {
	Nombre_fuente           string
	Nombre_poblacion_fuente string
	Formato_fuente          FormatoFuente
	Direccion               string
}

func NewFuente(nom string, nompoblacion string, ext FormatoFuente, dir string) (*Fuente, error) {
	if nompoblacion == "" {
		return nil, errors.New("una fuente debe hablar de alguna población")
	}
	if ext != HTML && ext != PDF && ext != EXCEL {
		return nil, errors.New("formato indefinido")
	}
	return &Fuente{
		Nombre_fuente:           nom,
		Nombre_poblacion_fuente: nompoblacion,
		Formato_fuente:          ext,
		Direccion:               dir,
	}, nil
}
