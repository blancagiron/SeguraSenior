package segurasenior

import "errors"

type Reportaje struct {
	titulo          string
	tendencias      []Tendencia
	datos_poblacion []Poblacion
	fuentes         []Fuente
}

func NewRepo(nom string, ten []Tendencia, datospob []Poblacion, origen []Fuente) (*Reportaje, error) {
	if nom == "" {
		return nil, errors.New("nombre no puede estar vacío")
	}
	return &Reportaje{
		titulo:          nom,
		tendencias:      ten,
		datos_poblacion: datospob,
		fuentes:         origen,
	}, nil
}
