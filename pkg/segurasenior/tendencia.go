package segurasenior

import (
	"errors"
)

type Estado int

const (
	decreciente Estado = iota
	creciente
)

func (e Estado) String() string {
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
	estadotendencia         Estado
	poblacion               Poblacion
}

func NewTendencia(nom string, estado Estado, pobl Poblacion) (*Tendencia, error) {
	if estado != decreciente && estado != creciente {
		return nil, errors.New("estado indefinido")
	}

	return &Tendencia{
		nombre_pueblo_tendencia: nom,
		estadotendencia:         estado,
		poblacion:               pobl,
	}, nil
}
