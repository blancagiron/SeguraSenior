package segurasenior

import (
	"errors"
	"time"
)

type EstadoPoblacion string

const (
	Decreciente EstadoPoblacion = "decreciente"
	Creciente   EstadoPoblacion = "creciente"
)

type FechaObtencionDeDatos struct {
	Dia  uint16
	Mes  time.Month
	Anio uint16
}

type IdentificadorDatos struct {
	NombrePoblacion string
	FechaDeDatos    FechaObtencionDeDatos
}

type RegistroDemografico struct {
	EstadisticasPoblacion map[IdentificadorDatos]DatosPoblacion
	EstadoDeLaPoblacion   EstadoPoblacion
}

func NewRegistroDemografico(datosPoblacion map[IdentificadorDatos]DatosPoblacion, estadoPoblacion EstadoPoblacion) (*RegistroDemografico, error) {
	if len(datosPoblacion) == 0 {
		return nil, errors.New("los datos de población no pueden estar vacíos")
	}
	return &RegistroDemografico{
		EstadisticasPoblacion: datosPoblacion,
		EstadoDeLaPoblacion:   estadoPoblacion,
	}, nil
}
