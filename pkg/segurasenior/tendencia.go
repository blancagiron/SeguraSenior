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
	PoblacionTotal  int
	Tasanat         int //medido en tantos por mil
	Tasaenv         int //medido en tantos por mil
	Tasamort        int //medido en tantos por mil
}

func NewTendencia(nom string, estado Estadopoblacion, poblacion int, tasan int, tasae int, tasam int) (*Tendencia, error) {
	if estado != Decreciente && estado != Creciente {
		return nil, errors.New("estado indefinido")
	}

	if (tasan < 0) || (tasan > 1000) || (tasae < 0) || (tasae > 1000) || (tasam < 0) || (tasam > 1000) {
		return nil, errors.New("valor de tasas erróneo, debe estar comprendido entre 0 y 1000")
	}

	return &Tendencia{
		Nombre:          nom,
		Estadotendencia: estado,
		PoblacionTotal:  poblacion,
		Tasanat:         tasan,
		Tasaenv:         tasae,
		Tasamort:        tasam,
	}, nil
}

func (t *Tendencia) ConflictoInfo(other *Tendencia) bool {
	if t.Nombre == other.Nombre && t.Estadotendencia != other.Estadotendencia {
		return true
	} else {
		return false
	}
}
