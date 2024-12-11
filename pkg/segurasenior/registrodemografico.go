package segurasenior

import (
	"errors"
)

type EstadoPoblacion string

const (
	Decreciente EstadoPoblacion = "decreciente"
	Creciente   EstadoPoblacion = "creciente"
)

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

type RegistrosDemografia map[string][]RegistroDemografico
