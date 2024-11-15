package segurasenior

import (
	"errors"
)

type Estadopoblacion int

const (
	Decreciente = iota
	Creciente
)

func (e Estadopoblacion) String() string {
	switch e {
	case Decreciente:
		return "decreciente"
	case Creciente:
		return "creciente"
	}
	return "unknown"
}

type Tendencia struct {
	Nombre          string
	Estadotendencia Estadopoblacion
}

func NewTendencia(nom string, estado Estadopoblacion) (*Tendencia, error) {
	if estado != Decreciente && estado != Creciente {
		return nil, errors.New("estado indefinido")
	}

	return &Tendencia{
		Nombre:          nom,
		Estadotendencia: estado,
	}, nil
}
