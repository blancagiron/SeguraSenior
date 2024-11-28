package segurasenior

import (
	"errors"
)

type Estado int

const (
	Decreciente Estado = iota
	Creciente
)

func (e Estado) String() string {
	switch e {
	case Decreciente:
		return "decreciente"
	case Creciente:
		return "creciente"
	}
	return "unknown"
}

type Tendencia struct {
	Estadotendencia Estado
}

func NewTendencia(estado Estado) (*Tendencia, error) {
	if nom == "" {
		return nil, errors.New("nombre no puede estar vacío")
	}
	return &Tendencia{
		Estadotendencia: estado,
	}, nil
}
