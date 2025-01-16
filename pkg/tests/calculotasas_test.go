package segurasenior

import (
	"SeguraSenior/pkg/segurasenior"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalcularTasas(t *testing.T) {

	datos := segurasenior.DatosPoblacion{
		PoblacionTotal: 350,
		Nacimientos:    20,
		Defunciones:    10,
	}

	expectedTasaNatalidad := 57.14
	expectedTasaMortalidad := 28.57

	datos.CalcularTasas()

	assert.Equal(t, expectedTasaNatalidad, datos.TasaNatalidadSobre1000, "Tasa de natalidad incorrecta")
	assert.Equal(t, expectedTasaMortalidad, datos.TasaMortalidadSobre1000, "Tasa de mortalidad incorrecta")
}
