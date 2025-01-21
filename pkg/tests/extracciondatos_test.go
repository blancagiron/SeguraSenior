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



func TestValidarDatos(test *testing.T) {
	test.Run("Datos válidos", func(test *testing.T) {
		err := segurasenior.ValidarDatos(struct {
			PoblacionTotal uint32  `json:"PoblacionTotal"`
			Hombres        uint32  `json:"Hombres"`
			Mujeres        uint32  `json:"Mujeres"`
			EdadMedia      float32 `json:"EdadMedia"`
			Menor20        float64 `json:"Menor20"`
			Mayor65        float64 `json:"Mayor65"`
			Nacimientos    uint32  `json:"Nacimientos"`
			Defunciones    uint32  `json:"Defunciones"`
		}{
			PoblacionTotal: 100,
			Hombres:        50,
			Mujeres:        50,
			EdadMedia:      30.5,
			Menor20:        15.0,
			Mayor65:        20.0,
			Nacimientos:    10,
			Defunciones:    5,
		})
		assert.NoError(test, err)
	})

	test.Run("Errores múltiples", func(test *testing.T) {
		err := segurasenior.ValidarDatos(struct {
			PoblacionTotal uint32  `json:"PoblacionTotal"`
			Hombres        uint32  `json:"Hombres"`
			Mujeres        uint32  `json:"Mujeres"`
			EdadMedia      float32 `json:"EdadMedia"`
			Menor20        float64 `json:"Menor20"`
			Mayor65        float64 `json:"Mayor65"`
			Nacimientos    uint32  `json:"Nacimientos"`
			Defunciones    uint32  `json:"Defunciones"`
		}{
			PoblacionTotal: 0,    
			Hombres:        50,  
			Mujeres:        40,
			Menor20:        -10,  
			Mayor65:        110,  
			Nacimientos:    200,  
			Defunciones:    5,
		})
		assert.Error(test, err)
	})
}
