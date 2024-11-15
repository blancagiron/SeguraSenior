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
	Direccion       string
}

func NewTendencia(nom string, estado Estadopoblacion, dir string) (*Tendencia, error) {
	if estado != Decreciente && estado != Creciente {
		return nil, errors.New("Estado indefinido")
	}

	return &Tendencia{
		Nombre:          nom,
		Estadotendencia: estado,
		Direccion:       dir,
	}, nil
}
