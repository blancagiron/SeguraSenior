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

func daysIn(mes time.Month, anio uint16) int {
	return time.Date(int(anio), mes+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

type RegistroDemografico struct {
	EstadisticasPoblacion map[IdentificadorDatos]DatosPoblacion
	EstadoDeLaPoblacion   EstadoPoblacion
}

func NewRegistroDemografico(datosPoblacion map[IdentificadorDatos]DatosPoblacion, estadoPoblacion EstadoPoblacion) (*RegistroDemografico, error) {
	for identificador := range datosPoblacion {
		if identificador.NombrePoblacion == "" {
			return nil, errors.New("el nombre de la población no puede estar vacío")
		}

		if int(identificador.FechaDeDatos.Dia) > daysIn(identificador.FechaDeDatos.Mes, identificador.FechaDeDatos.Anio) {
			return nil, errors.New("día de mes no válido")
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
