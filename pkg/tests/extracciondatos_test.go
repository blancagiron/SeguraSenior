package segurasenior

import (
	"SeguraSenior/pkg/segurasenior"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLecturaCorrectaDatosDesdeJSON(t *testing.T) {

	t.Run("Lectura Correcta - Caso Válido", func(t *testing.T) {
		nombreArchivo := "../testdata/data.json"
		nombrePoblacion := "Orcera"

		const (
			poblacionTotal    uint32  = 1735
			hombres           uint32  = 868
			mujeres           uint32  = 867
			edadMedia         float32 = 42.2
			porcentajeMenor20 float64 = 15.6
			porcentajeMayor65 float64 = 22.1
			nacimientos       uint32  = 6
			defunciones       uint32  = 29
		)

		tasaNatalidadSobre1000 := math.Round((float64(nacimientos)/float64(poblacionTotal))*1000*100) / 100
		tasaMortalidadSobre1000 := math.Round((float64(defunciones)/float64(poblacionTotal))*1000*100) / 100

		expectedIdentificador := segurasenior.IdentificadorDatos{
			NombrePoblacion: nombrePoblacion,
			FechaDeDatos: segurasenior.FechaObtencionDeDatos{
				Dia:  1,
				Mes:  1,
				Anio: 2023,
			},
		}

		expectedDatos := segurasenior.DatosPoblacion{
			PoblacionTotal:          poblacionTotal,
			Hombres:                 hombres,
			Mujeres:                 mujeres,
			EdadMedia:               edadMedia,
			PorcentajeMenora20:      porcentajeMenor20,
			PorcentajeMayora65:      porcentajeMayor65,
			Nacimientos:             nacimientos,
			Defunciones:             defunciones,
			TasaNatalidadSobre1000:  tasaNatalidadSobre1000,
			TasaMortalidadSobre1000: tasaMortalidadSobre1000,
		}

		identificador, datosExtraidos, err := segurasenior.LeerDatosDesdeJSON(nombreArchivo, nombrePoblacion)

		assert.NoError(t, err)
		assert.Equal(t, expectedIdentificador, identificador)
		assert.Equal(t, expectedDatos, datosExtraidos)
	})

	t.Run("Población No Encontrada", func(t *testing.T) {
		nombreArchivo := "../testdata/data.json"
		nombrePoblacion := "Granada"

		_, _, err := segurasenior.LeerDatosDesdeJSON(nombreArchivo, nombrePoblacion)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "población '"+nombrePoblacion+"' no encontrada en el archivo")
	})

	t.Run("JSON inválido", func(t *testing.T) {
		nombreArchivo := "../testdata/data_invalido.json"
		nombrePoblacion := "Orcera"

		_, _, err := segurasenior.LeerDatosDesdeJSON(nombreArchivo, nombrePoblacion)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid character")
	})

}

func TestValidarDatos(t *testing.T) {
	t.Run("Datos inválidos ", func(t *testing.T) {
		dato := struct {
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
		}{
			NombrePueblo:   "",
			FechaDatos:     "01/01/2023",
			PoblacionTotal: 0,
			Hombres:        500,
			Mujeres:        400,
			EdadMedia:      45.3,
			Menor20:        110,
			Mayor65:        -5,
			Nacimientos:    2000,
			Defunciones:    0,
		}

		err := segurasenior.ValidarDatos(dato)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "nombre de la población está vacío")
		assert.NotContains(t, err.Error(), "la fecha de datos está vacía")
		assert.NotContains(t, err.Error(), "el número de defunciones no puede ser mayor que la población total")
		assert.Contains(t, err.Error(), "el porcentaje de menores de 20 años debe estar entre 0 y 100")
		assert.Contains(t, err.Error(), "el porcentaje de mayores de 65 años debe estar entre 0 y 100")
		assert.Contains(t, err.Error(), "el número de nacimientos no puede ser mayor que la población total")
		assert.Contains(t, err.Error(), "la población total no puede ser 0")
	})
}
