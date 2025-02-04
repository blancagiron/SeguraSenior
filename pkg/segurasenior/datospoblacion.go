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

type IdentificadorDatos struct {
	NombrePoblacion string
	FechaDeDatos    FechaObtencionDeDatos
}


type DatosPoblacion struct {
	PoblacionTotal          uint32
	Hombres                 uint32
	Mujeres                 uint32
	EdadMedia               float32
	PorcentajeMenorA20      float64
	PorcentajeMayorA65      float64
	Nacimientos             uint32
	Defunciones             uint32
	TasaNatalidadSobre1000  float64
	TasaMortalidadSobre1000 float64
}

func NewDatosPoblacion(poblacion uint32, hombres uint32, mujeres uint32, edadMedia float32,
	porcentajeMenorA20 float64, porcentajeMayorA65 float64, nacimientos uint32, defunciones uint32, tasaNatalidad float64, tasaMortalidad float64) (*DatosPoblacion, error) {

	if hombres+mujeres != poblacion {
		return nil, errors.New("la población total no coincide con la suma de hombres y mujeres")
	}
	if porcentajeMenorA20 < 0 || porcentajeMenorA20 > 100 {
		return nil, errors.New("el porcentaje de menores de 20 años debe estar entre 0 y 100")
	}
	if porcentajeMayorA65 < 0 || porcentajeMayorA65 > 100 {
		return nil, errors.New("el porcentaje de mayores de 65 años debe estar entre 0 y 100")
	}
	if nacimientos > poblacion {
		return nil, errors.New("el número de nacimientos no puede ser mayor que la población total")
	}
	if defunciones > poblacion {
		return nil, errors.New("el número de defunciones no puede ser mayor que la población total")
	}

	return &DatosPoblacion{
		PoblacionTotal:          poblacion,
		Hombres:                 hombres,
		Mujeres:                 mujeres,
		EdadMedia:               edadMedia,
		PorcentajeMenorA20:      porcentajeMenorA20,
		PorcentajeMayorA65:      porcentajeMayorA65,
		Nacimientos:             nacimientos,
		Defunciones:             defunciones,
		TasaNatalidadSobre1000:  tasaNatalidad,
		TasaMortalidadSobre1000: tasaMortalidad,
	}, nil
}

func CrearIdentificadorDatos(nombre string, fecha time.Time) IdentificadorDatos {
	var identificador IdentificadorDatos
	identificador.NombrePoblacion = nombre
	identificador.FechaDeDatos = FechaObtencionDeDatos{
		Dia:  uint16(fecha.Day()),
		Mes:  fecha.Month(),
		Anio: uint16(fecha.Year()),
	}
	return identificador
}


func (datosPoblacion *DatosPoblacion) CalcularTasas() {
	if datosPoblacion.PoblacionTotal > 0 && datosPoblacion.Nacimientos < datosPoblacion.PoblacionTotal && datosPoblacion.Defunciones < datosPoblacion.PoblacionTotal {
		poblacionTotalFloat := float64(datosPoblacion.PoblacionTotal)
		datosPoblacion.TasaNatalidadSobre1000 = math.Round((float64(datosPoblacion.Nacimientos)/poblacionTotalFloat)*FactorTasasPorMil*PrecisionRedondeo) / PrecisionRedondeo
		datosPoblacion.TasaMortalidadSobre1000 = math.Round((float64(datosPoblacion.Defunciones)/poblacionTotalFloat)*FactorTasasPorMil*PrecisionRedondeo) / PrecisionRedondeo
	}
}