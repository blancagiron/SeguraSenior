package segurasenior

import (
	"errors"
)

type Estadopoblacion int

const (
	decreciente = iota
	creciente
)

func (e Estadopoblacion) String() string {
	switch e {
	case decreciente:
		return "decreciente"
	case creciente:
		return "creciente"
	}
	return "unknown"
}

type Tendencia struct {
	nombre_pueblo_tendencia string
	estadotendencia         Estadopoblacion
	poblacion               Poblacion
}

func NewTendencia(nom string, estado Estadopoblacion, pobl Poblacion) (*Tendencia, error) {
	if estado != decreciente && estado != creciente {
		return nil, errors.New("estado indefinido")
	}

	return &Tendencia{
		nombre_pueblo_tendencia: nom,
		estadotendencia:         estado,
		poblacion:               pobl,
	}, nil
}
