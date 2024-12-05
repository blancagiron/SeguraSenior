package segurasenior

import (
	"errors"
	"time"
)

type FechaObtencionDeDatos struct {
	Dia  uint16
	Mes  time.Month
	Anio uint16
}

type DatosPoblacion struct {
	NombrePoblacion    string
	FechaDeDatos       FechaObtencionDeDatos
	PoblacionTotal     uint32
	PoblacionJoven     uint32
	PoblacionAdulta    uint32
	PoblacionAnciana   uint32
	TasaNatalidad      float64
	TasaEnvejecimiento float64
	TasaMortalidad     float64
}

func NewDatosPoblacion(nombrepoblacion string, fechadatos FechaObtencionDeDatos, poblacion uint32, poblacionMenor18 uint32, poblacionEntre18Y64 uint32, poblacionMayor64 uint32,
	tasaNatalidadSobre1000 float64, tasaEnvejecimientoSobre1000 float64, tasaMortalidadSobre1000 float64) (*DatosPoblacion, error) {
	if nombrepoblacion == "" {
		return nil, errors.New("nombre no puede estar vacío")
	}

	if nombrepoblacion == "" {
		return nil, errors.New("nombre no puede estar vacío")
	}

	if poblacionMenor18+poblacionEntre18Y64+poblacionMayor64 != poblacion {
		return nil, errors.New("valor de poblacion total erróneo, debe ser la suma de la población menor a 18, entre 18 y 64 y mayor a 64")
	}

	if (tasaNatalidadSobre1000 < 0.0) || (tasaNatalidadSobre1000 > 1000.0) {
		return nil, errors.New("valor de tasa de natalidada erróneo, debe estar comprendido entre 0 y 1000")
	}
	if (tasaEnvejecimientoSobre1000 < 0.0) || (tasaEnvejecimientoSobre1000 > 1000.0) {
		return nil, errors.New("valor de tasa de envejecimiento erróneo, debe estar comprendido entre 0 y 1000")
	}
	if (tasaMortalidadSobre1000 < 0.0) || (tasaMortalidadSobre1000 > 1000.0) {
		return nil, errors.New("valor de tasa de mortalidad erróneo, debe estar comprendido entre 0 y 1000")
	}
	return &DatosPoblacion{
		NombrePoblacion:    nombrepoblacion,
		PoblacionTotal:     poblacion,
		PoblacionJoven:     poblacionMenor18,
		PoblacionAdulta:    poblacionEntre18Y64,
		PoblacionAnciana:   poblacionMayor64,
		TasaNatalidad:      tasaNatalidadSobre1000,
		TasaEnvejecimiento: tasaEnvejecimientoSobre1000,
		TasaMortalidad:     tasaMortalidadSobre1000,
	}, nil
}
