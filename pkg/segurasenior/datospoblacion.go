package segurasenior

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

type FechaObtencionDeDatos struct {
	Dia  uint16
	Mes  time.Month
	Anio uint16
}

type IdentificadorDatos struct {
	NombrePoblacion string                `json:"NombrePoblacion"`
	FechaDeDatos    FechaObtencionDeDatos `json:"FechaDeDatos"`
}

type DatosPoblacion struct {
	PoblacionTotal          uint32  `json:"PoblacionTotal"`
	Hombres                 uint32  `json:"Hombres"`
	Mujeres                 uint32  `json:"Mujeres"`
	EdadMedia               float32 `json:"EdadMedia"`
	PorcentajeMenora20      float64 `json:"Menor20"`
	PorcentajeMayora65      float64 `json:"Mayor65"`
	Nacimientos             uint32  `json:"Nacimientos"`
	Defunciones             uint32  `json:"Defunciones"`
	TasaNatalidadSobre1000  float64 `json:"TasaNatalidadSobre1000"`
	TasaMortalidadSobre1000 float64 `json:"TasaMortalidadSobre1000"`
}

func NewIdentificadorDatos(nombrePoblacion string, fecha FechaObtencionDeDatos) IdentificadorDatos {
	return IdentificadorDatos{
		NombrePoblacion: nombrePoblacion,
		FechaDeDatos:    fecha,
	}
}


const (
	FactorTasasPorMil = 1000
	PrecisionRedondeo = 100
)



func (datosPoblacion *DatosPoblacion) CalcularTasas() {
	if datosPoblacion.PoblacionTotal > 0 && datosPoblacion.Nacimientos < datosPoblacion.PoblacionTotal && datosPoblacion.Defunciones < datosPoblacion.PoblacionTotal {
		datosPoblacion.TasaNatalidadSobre1000 = math.Round((float64(datosPoblacion.Nacimientos)/float64(datosPoblacion.PoblacionTotal))*FactorTasasPorMil*PrecisionRedondeo) / PrecisionRedondeo
		datosPoblacion.TasaMortalidadSobre1000 = math.Round((float64(datosPoblacion.Defunciones)/float64(datosPoblacion.PoblacionTotal))*FactorTasasPorMil*PrecisionRedondeo) / PrecisionRedondeo
	}
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
		PorcentajeMenora20:      porcentajeMenorA20,
		PorcentajeMayora65:      porcentajeMayorA65,
		Nacimientos:             nacimientos,
		Defunciones:             defunciones,
		TasaNatalidadSobre1000:  tasaNatalidad,
		TasaMortalidadSobre1000: tasaMortalidad,
	}, nil
}

func LeerIdentificadorDesdeJSON(nombreArchivo, nombrePoblacion string) (IdentificadorDatos, error) {
	
}

func LeerDatosDesdeJSON(nombreArchivo, nombrePoblacion string) (IdentificadorDatos, DatosPoblacion, error) {
	var identificador IdentificadorDatos
	var datosPoblacion DatosPoblacion

	file, err := os.Open(nombreArchivo)
	if err != nil {
		return identificador, datosPoblacion, fmt.Errorf("no se pudo abrir el archivo: %w", err)
	}
	defer file.Close()

	var datos map[string]struct {
		NombrePueblo   string  `json:"NombrePueblo"`
		FechaDatos     string  `json:"FechaDatos"`
		PoblacionTotal uint32  `json:"PoblacionTotal"`
		Hombres        uint32  `json:"Hombres"`
		Mujeres        uint32  `json:"Mujeres"`
		EdadMedia      float32 `json:"EdadMedia"`
		Menor20        float64 `json:"Menor20"`
		Mayor65        float64 `json:"Mayor65"`
		Nacimientos    uint32  `json:"Nacimientos"`
		Defunciones    uint32  `json:"Defunciones"`
	}

	if err := json.NewDecoder(file).Decode(&datos); err != nil {
		return identificador, datosPoblacion, fmt.Errorf("error al decodificar JSON: %w", err)
	}

	dato, existe := datos[nombrePoblacion]
	if !existe {
		return identificador, datosPoblacion, fmt.Errorf("población '%s' no encontrada en el archivo", nombrePoblacion)
	}

	if err := ValidarDatos(dato); err != nil {
		return identificador, datosPoblacion, err
	}

	fecha, err := time.Parse("02/01/2006", dato.FechaDatos)
	if err != nil {
		return identificador, datosPoblacion, fmt.Errorf("formato de fecha inválido ('%s'): %w", dato.FechaDatos, err)
	}

	identificador = IdentificadorDatos{
		NombrePoblacion: dato.NombrePueblo,
		FechaDeDatos: FechaObtencionDeDatos{
			Dia:  uint16(fecha.Day()),
			Mes:  fecha.Month(),
			Anio: uint16(fecha.Year()),
		},
	}

	datosPoblacion = DatosPoblacion{
		PoblacionTotal:     dato.PoblacionTotal,
		Hombres:            dato.Hombres,
		Mujeres:            dato.Mujeres,
		EdadMedia:          dato.EdadMedia,
		PorcentajeMenora20: dato.Menor20,
		PorcentajeMayora65: dato.Mayor65,
		Nacimientos:        dato.Nacimientos,
		Defunciones:        dato.Defunciones,
	}

	datosPoblacion.CalcularTasas()

	return identificador, datosPoblacion, nil
}

func ValidarDatos(dato struct {
	NombrePueblo   string  `json:"NombrePueblo"`
	FechaDatos     string  `json:"FechaDatos"`
	PoblacionTotal uint32  `json:"PoblacionTotal"`
	Hombres        uint32  `json:"Hombres"`
	Mujeres        uint32  `json:"Mujeres"`
	EdadMedia      float32 `json:"EdadMedia"`
	Menor20        float64 `json:"Menor20"`
	Mayor65        float64 `json:"Mayor65"`
	Nacimientos    uint32  `json:"Nacimientos"`
	Defunciones    uint32  `json:"Defunciones"`
}) error {
	var errores []string

	if dato.NombrePueblo == "" {
		errores = append(errores, "nombre de la población está vacío")
	}
	if dato.FechaDatos == "" {
		errores = append(errores, "la fecha de datos está vacía")
	}
	if dato.PoblacionTotal == 0 {
		errores = append(errores, "la población total no puede ser 0")
	}
	if dato.Hombres+dato.Mujeres != dato.PoblacionTotal {
		errores = append(errores, "la población total no coincide con la suma de hombres y mujeres")
	}
	if dato.Menor20 < 0 || dato.Menor20 > 100 {
		errores = append(errores, "el porcentaje de menores de 20 años debe estar entre 0 y 100")
	}
	if dato.Mayor65 < 0 || dato.Mayor65 > 100 {
		errores = append(errores, "el porcentaje de mayores de 65 años debe estar entre 0 y 100")
	}
	if dato.Nacimientos > dato.PoblacionTotal {
		errores = append(errores, "el número de nacimientos no puede ser mayor que la población total")
	}
	if dato.Defunciones > dato.PoblacionTotal {
		errores = append(errores, "el número de defunciones no puede ser mayor que la población total")
	}

	if len(errores) > 0 {
		return fmt.Errorf("errores de validación: %s", strings.Join(errores, ", "))
	}

	return nil
}
