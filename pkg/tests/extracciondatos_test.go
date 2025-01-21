package segurasenior_test

import (
	//"encoding/json"
	//"os"
	"testing"
	//"time"

	"github.com/stretchr/testify/assert"
	"SeguraSenior/pkg/segurasenior"
)

const archivoPruebas = "../testdata/data.json"
const archivoInvalido = "../testdata/data_invalido.json"

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
			PoblacionTotal: 0,    // Error: población 0
			Hombres:        50,   // Error: suma incorrecta
			Mujeres:        40,
			Menor20:        -10,  // Error: porcentaje inválido
			Mayor65:        110,  // Error: porcentaje inválido
			Nacimientos:    200,  // Error: mayor que población
			Defunciones:    5,
		})
		assert.Error(test, err)
	})
}
