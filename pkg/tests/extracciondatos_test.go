package segurasenior_test

import (
	
	"testing"
	
	"github.com/stretchr/testify/assert"
	"SeguraSenior/pkg/segurasenior"
)


const (
	archivoPruebas  = "../testdata/data.json"
	archivoInvalido = "../testdata/data_invalido.json"
)

func TestCargarDatosDesdeJSON(test *testing.T) {
	datos, err := segurasenior.CargarDatosDesdeJSON[map[string]interface{}](archivoPruebas)
	assert.NoError(test, err)
	assert.Contains(test, datos, "Orcera")

	_, err = segurasenior.CargarDatosDesdeJSON[map[string]interface{}](archivoInvalido)
	assert.Error(test, err)
}

func TestLeerIdentificadorDesdeJSON(test *testing.T) {
	casosDeTest := []struct {
		descripcion      string
		poblacion        string
		esValido         bool
		nombreEsperado   string
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
