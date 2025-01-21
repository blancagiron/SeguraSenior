package segurasenior_test

import (
	"testing"

	"SeguraSenior/pkg/segurasenior"
	"github.com/stretchr/testify/assert"
)

const (
	archivoPruebas  = "../testdata/data.json"
	archivoInvalido = "../testdata/data_invalido.json"
)

func TestCargarDatosDesdeJSON(test *testing.T) {
	datos, err := segurasenior.CargarDatosDesdeJSON[map[string]interface{}](archivoPruebas)
	assert.NoError(test, err)
	assert.NotEmpty(test, datos)

	_, err = segurasenior.CargarDatosDesdeJSON[map[string]interface{}](archivoInvalido)
	assert.Error(test, err)
}

func TestLeerIdentificadorDesdeJSON(test *testing.T) {
	casosDeTest := []struct {
		descripcion    string
		poblacion      string
		esValido       bool
		nombreEsperado string
	}{
		{"Identificador válido", "Orcera", true, "Orcera"},
		{"Identificador inexistente", "Pueblo Inexistente", false, ""},
		{"Fecha inválida en datos", "Siles", false, ""},
	}

	for _, casos := range casosDeTest {
		test.Run(casos.descripcion, func(test *testing.T) {
			identificador, err := segurasenior.LeerIdentificadorDesdeJSON(archivoPruebas, casos.poblacion)
			if casos.esValido {
				assert.NoError(test, err)
				assert.Equal(test, casos.nombreEsperado, identificador.NombrePoblacion)
			} else {
				assert.Error(test, err)
			}
		})
	}
}
func TestLeerDatosDesdeJSON(test *testing.T) {
	const (
		poblacionTotal         uint32  = 1735
		hombres                uint32  = 868
		mujeres                uint32  = 867
		edadMedia              float32 = 42.2
		porcentajeMenor20      float64 = 15.6
		porcentajeMayor65      float64 = 22.1
		nacimientos            uint32  = 6
		defunciones            uint32  = 29
		tasaNatalidadEsperada  float64 = 3.46
		tasaMortalidadEsperada float64 = 16.71
	)

	test.Run("Lectura correcta", func(test *testing.T) {
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

	validarErrores := func(test *testing.T, err error, mensajeEsperado []string, mensajeNoEsperado []string) {
		assert.Error(test, err)
		for _, mensaje := range mensajeEsperado {
			assert.Contains(test, err.Error(), mensaje)
		}
		for _, mensaje := range mensajeNoEsperado {
			assert.NotContains(test, err.Error(), mensaje)
		}
	}

	test.Run("Datos inválidos", func(test *testing.T) {
		_, err := segurasenior.LeerDatosDesdeJSON(archivoPruebas, "Siles")
		assert.Error(test, err)
		validarErrores(test, err, []string{
			"la población total no puede ser 0",
			"la población total no coincide con la suma de hombres y mujeres",
			"el porcentaje de menores de 20 años debe estar entre 0 y 100",
			"el porcentaje de mayores de 65 años debe estar entre 0 y 100",
			"el número de nacimientos no puede ser mayor que la población total",
			"el número de defunciones no puede ser mayor que la población total",
		},
			[]string{
				"el número de hombres no puede ser 0",
				"el número de mujeres no puede ser 0",
				"la edad media no puede ser 0",
			},
		)

	})

}
