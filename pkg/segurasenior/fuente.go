package segurasenior

import "errors"

type Fuente struct {
	Nombre_poblacion_fuente string
	Direccion               string
}

func NewFuente(nompoblacion string, dir string) (*Fuente, error) {
	if nompoblacion == "" {
		return nil, errors.New("una fuente debe hablar de alguna población")
	}
	return &Fuente{
		Nombre_poblacion_fuente: nompoblacion,
		Direccion:               dir,
	}, nil
}
