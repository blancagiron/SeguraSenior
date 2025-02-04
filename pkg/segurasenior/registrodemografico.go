package segurasenior

import (
	"errors"
)

type EstadoPoblacion string

const (
	Decreciente EstadoPoblacion = "decreciente"
	Creciente   EstadoPoblacion = "creciente"
)

type RegistroDemografico struct {
	EstadisticasPoblacion map[IdentificadorDatos]DatosPoblacion
	EstadoDeLaPoblacion   EstadoPoblacion
}

func ObtenerEstadoPoblacion(datos DatosPoblacion) EstadoPoblacion {
	if datos.TasaMortalidadSobre1000 > datos.TasaNatalidadSobre1000 {
		return Decreciente
	}
	return Creciente
}

func NewRegistroDemografico(datosPoblacion map[IdentificadorDatos]DatosPoblacion, estadoPoblacion EstadoPoblacion) (*RegistroDemografico, error) {
	for identificador := range datosPoblacion {
		if identificador.NombrePoblacion == "" {
			return nil, errors.New("el nombre de la población no puede estar vacío")
		}
	}
	if len(datosPoblacion) == 0 {
		return nil, errors.New("los datos de población no pueden estar vacíos")
	}
	return &RegistroDemografico{
		EstadisticasPoblacion: datosPoblacion,
		EstadoDeLaPoblacion:   estadoPoblacion,
	}, nil
}