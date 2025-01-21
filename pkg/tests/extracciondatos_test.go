package segurasenior_test

import (
	//"encoding/json"
	//"os"
	"testing"
	//"time"

	"github.com/stretchr/testify/assert"
	"SeguraSenior/pkg/segurasenior"
)

const (
	archivoPruebas = "../testdata/data.json"
	archivoInvalido = "../testdata/data_invalido.json"
)

func TestCargarDatosDesdeJSON(test *testing.T) {
	test.Run("Cargar datos válidos", func(test *testing.T) {
		datos, err := segurasenior.CargarDatosDesdeJSON[map[string]interface{}](archivoPruebas)
		assert.NoError(test, err)
		assert.Contains(test, datos, "Orcera")
	})	

	test.Run("Cargar datos inválidos", func(test *testing.T) {
		_, err := segurasenior.CargarDatosDesdeJSON[map[string]interface{}](archivoInvalido)
		assert.Error(test, err)
	})	
}

func TestLeerIdentificadorDesdeJSON(test *testing.T) {
	test.Run("Leer identificador válido", func(test *testing.T) {
		identificador, err := segurasenior.LeerIdentificadorDesdeJSON(archivoPruebas, "Orcera")
		assert.NoError(test, err)
		assert.Equal(test, "Orcera", identificador.NombrePoblacion)
	})

	test.Run("Leer identificador inválido", func(test *testing.T) {
		_, err := segurasenior.LeerIdentificadorDesdeJSON(archivoPruebas, "Pueblo Inexistente")
		assert.Error(test, err)
	})

	test.Run("Leer identificador con fecha inválida", func(test *testing.T) {
		_, err := segurasenior.LeerIdentificadorDesdeJSON(archivoPruebas, "Siles")
		assert.Error(test, err)
	})

}

func TestLeerDatosDesdeJSON(test *testing.T) {
	test.Run("Leer JSON no válido", func(test *testing.T) {
		_, err := segurasenior.LeerDatosDesdeJSON(archivoInvalido, "Orcera")
		assert.Error(test, err)
	})

	test.Run("Lectura correcta - Caso válido", func(test *testing.T) {
		const (
			poblacionTotal    uint32  = 1735
			hombres           uint32  = 868
			mujeres           uint32  = 867
			edadMedia         float32 = 42.2
			porcentajeMenor20 float64 = 15.6
			porcentajeMayor65 float64 = 22.1
			nacimientos       uint32  = 6
			defunciones       uint32  = 29
			tasaNatalidadEsperada float64 = 3.46
			tasaMortalidadEsperada float64 = 16.71
		)

		datos, err := segurasenior.LeerDatosDesdeJSON(archivoPruebas, "Orcera")
		assert.NoError(test, err)

		assert.Equal(test, poblacionTotal, datos.PoblacionTotal)
		assert.Equal(test, hombres, datos.Hombres)
		assert.Equal(test, mujeres, datos.Mujeres)
		assert.Equal(test, edadMedia, datos.EdadMedia)
		assert.Equal(test, porcentajeMenor20, datos.PorcentajeMenorA20)
		assert.Equal(test, porcentajeMayor65, datos.PorcentajeMayorA65)
		assert.Equal(test, nacimientos, datos.Nacimientos)
		assert.Equal(test, defunciones, datos.Defunciones)
		assert.Equal(test, tasaNatalidadEsperada, datos.TasaNatalidadSobre1000)
		assert.Equal(test, tasaMortalidadEsperada, datos.TasaMortalidadSobre1000)
	})
}


func TestValidarDatos(t *testing.T) {
	const (
		PorcentajeInvalidoMayorA60 = 110.0
		PorcentajeInvalidoMenorA20 = -5.0
		PorcentajeValidoMayorA60   = 60.0
		PorcentajeValidoMenorA20   = 20.0
		PoblacionCero           = 0
		PoblacionTotalValida    = 1000
		HombresValidos          = 500
		MujeresValidas          = 500
		NacimientosInvalidos    = 2000
		DefuncionesValidas      = 10
		NacimientosValidos      = 20

	)

	t.Run("Datos inválidos", func(t *testing.T) {
		dato := struct {
			PoblacionTotal uint32  `json:"PoblacionTotal"`
			Hombres        uint32  `json:"Hombres"`
			Mujeres        uint32  `json:"Mujeres"`
			EdadMedia      float32 `json:"EdadMedia"`
			Menor20        float64 `json:"Menor20"`
			Mayor65        float64 `json:"Mayor65"`
			Nacimientos    uint32  `json:"Nacimientos"`
			Defunciones    uint32  `json:"Defunciones"`
		}{
			PoblacionTotal: PoblacionCero,
			Hombres:        HombresValidos,
			Mujeres:        MujeresValidas,
			EdadMedia:      45.3,
			Menor20:        PorcentajeInvalidoMayorA60,
			Mayor65:        PorcentajeInvalidoMenorA20,
			Nacimientos:    NacimientosInvalidos,
			Defunciones:    DefuncionesValidas,
		}

		err := segurasenior.ValidarDatos(dato)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "la población total no puede ser 0")
		assert.NotContains(t, err.Error(), "el numero de hombres no puede ser 0")
		assert.NotContains(t, err.Error(), "el numero de mujeres no puede ser 0")
		assert.Contains(t, err.Error(), "el porcentaje de menores de 20 años debe estar entre 0 y 100")
		assert.Contains(t, err.Error(), "el porcentaje de mayores de 65 años debe estar entre 0 y 100")
		assert.Contains(t, err.Error(), "el número de nacimientos no puede ser mayor que la población total")
	})

	t.Run("Datos válidos", func(t *testing.T) {
		dato := struct {
			PoblacionTotal uint32  `json:"PoblacionTotal"`
			Hombres        uint32  `json:"Hombres"`
			Mujeres        uint32  `json:"Mujeres"`
			EdadMedia      float32 `json:"EdadMedia"`
			Menor20        float64 `json:"Menor20"`
			Mayor65        float64 `json:"Mayor65"`
			Nacimientos    uint32  `json:"Nacimientos"`
			Defunciones    uint32  `json:"Defunciones"`
		}{
			PoblacionTotal: PoblacionTotalValida,
			Hombres:        HombresValidos,
			Mujeres:        MujeresValidas,
			EdadMedia:      45.3,
			Menor20:        PorcentajeValidoMenorA20,
			Mayor65:        PorcentajeValidoMayorA60,
			Nacimientos:    NacimientosValidos,
			Defunciones:    DefuncionesValidas,
		}

		err := segurasenior.ValidarDatos(dato)
		assert.NoError(t, err)
	})
}

