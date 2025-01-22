package segurasenior

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"os"
	"time"
	"io"
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

const (
	FactorTasasPorMil = 1000
	PrecisionRedondeo = 100
)

func (datosPoblacion *DatosPoblacion) CalcularTasas() {
	if datosPoblacion.PoblacionTotal > 0 && datosPoblacion.Nacimientos < datosPoblacion.PoblacionTotal && datosPoblacion.Defunciones < datosPoblacion.PoblacionTotal {
		poblacionTotalFloat := float64(datosPoblacion.PoblacionTotal)
		datosPoblacion.TasaNatalidadSobre1000 = math.Round((float64(datosPoblacion.Nacimientos)/poblacionTotalFloat)*FactorTasasPorMil*PrecisionRedondeo) / PrecisionRedondeo
		datosPoblacion.TasaMortalidadSobre1000 = math.Round((float64(datosPoblacion.Defunciones)/poblacionTotalFloat)*FactorTasasPorMil*PrecisionRedondeo) / PrecisionRedondeo
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
		PorcentajeMenorA20:      porcentajeMenorA20,
		PorcentajeMayorA65:      porcentajeMayorA65,
		Nacimientos:             nacimientos,
		Defunciones:             defunciones,
		TasaNatalidadSobre1000:  tasaNatalidad,
		TasaMortalidadSobre1000: tasaMortalidad,
	}, nil
}

func DecodificarJSON[Tipo any](contenido []byte) (map[string]Tipo, error) {
	var datos map[string]Tipo
	if err := json.Unmarshal(contenido, &datos); err != nil {
		return nil, fmt.Errorf("error al decodificar JSON: %w", err)
	}
	return datos, nil
}


func LeerArchivo(nombreArchivo string) ([]byte, error) {
	file, err := os.Open(nombreArchivo)
	if err != nil {
		return nil, fmt.Errorf("no se pudo abrir el archivo: %w", err)
	}
	defer file.Close()

	contenido, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error al leer el archivo: %w", err)
	}

	return contenido, nil
}

func CargarDatosDesdeArchivo[Tipo any](nombreArchivo string) (map[string]Tipo, error) {
	contenido, err := LeerArchivo(nombreArchivo)
	if err != nil {
		return nil, fmt.Errorf("error al leer datos desde el archivo: %w", err)
	}
	datosDecodificados, err := DecodificarJSON[Tipo](contenido)
	if err != nil {
		return nil, fmt.Errorf("error al decodificar datos desde JSON: %w", err)
	}
	return datosDecodificados, nil
}

func ValidarPoblacionExiste[Tipo any](datos map[string]Tipo, nombrePoblacion string) (Tipo, error) {
	datoPoblacionEspecifica, existe := datos[nombrePoblacion]
	if !existe {
		return datoPoblacionEspecifica, fmt.Errorf("población '%s' no encontrada", nombrePoblacion)
	}
	return datoPoblacionEspecifica, nil
}

func ParsearFecha(fechaEnString string) (time.Time, error) {
	fecha, err := time.Parse("02/01/2006", fechaEnString)
	if err != nil {
		return time.Time{}, fmt.Errorf("formato de fecha inválido ('%s'): %w", fechaEnString, err)
	}
	return fecha, nil
}

func CrearIdentificadorDatos(nombre string, fecha time.Time) IdentificadorDatos {
	return IdentificadorDatos{
		NombrePoblacion: nombre,
		FechaDeDatos: FechaObtencionDeDatos{
			Dia:  uint16(fecha.Day()),
			Mes:  fecha.Month(),
			Anio: uint16(fecha.Year()),
		},
	}
}



func LeerDatosPoblacionDesdeArchivo(nombreArchivo, nombrePoblacion string) (*DatosPoblacion, error) {
	datos, err := CargarDatosDesdeArchivo[struct {
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
		return nil, fmt.Errorf("error al cargar datos desde el archivo JSON: %w", err)
	}

	datoPoblacionEspecifica, err := ValidarPoblacionExiste(datos, nombrePoblacion)
	if err != nil {
		return nil, fmt.Errorf("error al validar población: %w", err)
	}

	datosPoblacion, err := NewDatosPoblacion(
		datoPoblacionEspecifica.PoblacionTotal,
		datoPoblacionEspecifica.Hombres,
		datoPoblacionEspecifica.Mujeres,
		datoPoblacionEspecifica.EdadMedia,
		datoPoblacionEspecifica.Menor20,
		datoPoblacionEspecifica.Mayor65,
		datoPoblacionEspecifica.Nacimientos,
		datoPoblacionEspecifica.Defunciones,
		0, 
		0, 
	)
	if err != nil {
		return nil, fmt.Errorf("error al crear instancia de DatosPoblacion: %w", err)
	}

	datosPoblacion.CalcularTasas()

	return datosPoblacion, nil
}


