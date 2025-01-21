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

func CargarDatosDesdeJSON[Tipo any](nombreArchivo string) (map[string]Tipo, error) {

	var datos map[string]Tipo

	file, err := os.Open(nombreArchivo)
	if err != nil {
		return nil, fmt.Errorf("no se pudo abrir el archivo: %w", err)
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&datos); err != nil {
		return nil, fmt.Errorf("error al decodificar JSON: %w", err)
	}

	return datos, nil
}

func LeerIdentificadorDesdeJSON(nombreArchivo, nombrePoblacion string) (IdentificadorDatos, error) {

	datos, err := CargarDatosDesdeJSON[struct {
		NombrePueblo string `json:"NombrePueblo"`
		FechaDatos   string `json:"FechaDatos"`
	}](nombreArchivo)
	if err != nil {
		return IdentificadorDatos{}, fmt.Errorf("error al cargar datos desde JSON: %w", err)
	}

	dato, existe := datos[nombrePoblacion]
	if !existe {
		return IdentificadorDatos{}, fmt.Errorf("población '%s' no encontrada en el archivo", nombrePoblacion)
	}

	fecha, err := time.Parse("02/01/2006", dato.FechaDatos)
	if err != nil {
		return IdentificadorDatos{}, fmt.Errorf("formato de fecha inválido ('%s'): %w", dato.FechaDatos, err)
	}

	return IdentificadorDatos{
		NombrePoblacion: dato.NombrePueblo,
		FechaDeDatos: FechaObtencionDeDatos{
			Dia:  uint16(fecha.Day()),
			Mes:  fecha.Month(),
			Anio: uint16(fecha.Year()),
		},
	}, nil
}

func LeerDatosDesdeJSON(nombreArchivo, nombrePoblacion string) (DatosPoblacion, error) {

	var datosPoblacion DatosPoblacion

	datos, err := CargarDatosDesdeJSON[struct {
		PoblacionTotal uint32  `json:"PoblacionTotal"`
		Hombres        uint32  `json:"Hombres"`
		Mujeres        uint32  `json:"Mujeres"`
		EdadMedia      float32 `json:"EdadMedia"`
		Menor20        float64 `json:"Menor20"`
		Mayor65        float64 `json:"Mayor65"`
		Nacimientos    uint32  `json:"Nacimientos"`
		Defunciones    uint32  `json:"Defunciones"`
	}](nombreArchivo)
	if err != nil {
		return DatosPoblacion{}, fmt.Errorf("error al cargar datos desde JSON: %w", err)
	}

	dato, existe := datos[nombrePoblacion]
	if !existe {
		return DatosPoblacion{}, fmt.Errorf("población '%s' no encontrada en el archivo", nombrePoblacion)
	}

	if err := ValidarDatos(dato); err != nil {
		return DatosPoblacion{}, fmt.Errorf("datos inválidos: %w", err)
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

	return datosPoblacion, nil
}

func ValidarDatos(dato struct {
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
