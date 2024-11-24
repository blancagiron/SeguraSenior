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
		return nil, errors.New("debe haber un título")
	}
	if len(origen) == 0 {
		return nil, errors.New("el reportaje debe provenir de alguna fuente")
	}
	return &Reportaje{
		titulo:          nom,
		tendencias:      ten,
		datos_poblacion: datospob,
		fuentes:         origen,
	}, nil
}
