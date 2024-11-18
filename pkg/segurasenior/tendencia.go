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
	Poblacion       Poblacion
}

func NewTendencia(nom string, estado Estadopoblacion, pobl Poblacion) (*Tendencia, error) {
	if estado != Decreciente && estado != Creciente {
		return nil, errors.New("estado indefinido")
	}

	return &Tendencia{
		Nombre:          nom,
		Estadotendencia: estado,
		Poblacion:       pobl,
	}, nil
}

func (t *Tendencia) ConflictoInfo(other *Tendencia) bool {
	if t.Nombre == other.Nombre && t.Estadotendencia != other.Estadotendencia {
		return true
	} else {
		return false
	}
}
