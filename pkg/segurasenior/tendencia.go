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
	Nombre_pueblo_tendencia string
	Estadotendencia         Estado
}

func NewTendencia(nom string, estado Estado) (*Tendencia, error) {
	if nom == "" {
		return nil, errors.New("nombre no puede estar vacío")
	}
	if estado != decreciente && estado != creciente {
		return nil, errors.New("estado indefinido")
	}
	return &Tendencia{
		Nombre_pueblo_tendencia: nom,
		Estadotendencia:         estado,
	}, nil
}
