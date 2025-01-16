package segurasenior

import (
	"encoding/json"
	"errors"
	"os"
	"time"
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

type DatosPoblacion struct {
	PoblacionTotal          uint32
	Hombres                 uint32
	Mujeres                 uint32
	EdadMedia               float32
	PorcentajeMenora20      float64
	PorcentajeMayora65      float64
	Nacimientos             uint32
	Defunciones             uint32
	TasaNatalidadSobre1000  float64
	TasaMortalidadSobre1000 float64
}



func NewDatosPoblacion(poblacion uint32, hombres uint32, mujeres uint32, edadMedia float32,
	porcentajeMenora20 float64, porcentajeMayora65 float64, nacimientos uint32, defunciones uint32) (*DatosPoblacion, error) {

	if hombres+mujeres != poblacion {
		return nil, errors.New("la población total no coincide con la suma de hombres y mujeres")
	}
	if porcentajeMenora20 < 0 || porcentajeMenora20 > 100 {
		return nil, errors.New("el porcentaje de menores de 20 años debe estar entre 0 y 100")
	}
	if porcentajeMayora65 < 0 || porcentajeMayora65 > 100 {
		return nil, errors.New("el porcentaje de mayores de 65 años debe estar entre 0 y 100")
	}

	return &DatosPoblacion{
		PoblacionTotal:     poblacion,
		Hombres:            hombres,
		Mujeres:            mujeres,
		EdadMedia:          edadMedia,
		PorcentajeMenora20: porcentajeMenora20,
		PorcentajeMayora65: porcentajeMayora65,
		Nacimientos:        nacimientos,
		Defunciones:        defunciones,
	}, nil
}


