package segurasenior

import (
	"SeguraSenior/pkg/segurasenior"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDatosPoblacion(t *testing.T) {
	t.Run("Datos válidos", func(t *testing.T) {
		datos, err := segurasenior.NewDatosPoblacion(1000, 500, 500, 30.5, 25.0, 15.0, 50, 20, 50.0, 20.0)
		assert.NoError(t, err)
		assert.NotNil(t, datos)
		assert.Equal(t, uint32(1000), datos.PoblacionTotal)
		assert.Equal(t, uint32(500), datos.Hombres)
		assert.Equal(t, uint32(500), datos.Mujeres)
		assert.Equal(t, float32(30.5), datos.EdadMedia)
		assert.Equal(t, float64(25.0), datos.PorcentajeMenora20)
		assert.Equal(t, float64(15.0), datos.PorcentajeMayora65)
		assert.Equal(t, uint32(50), datos.Nacimientos)
		assert.Equal(t, uint32(20), datos.Defunciones)
		assert.Equal(t, float64(50.0), datos.TasaNatalidadSobre1000)
		assert.Equal(t, float64(20.0), datos.TasaMortalidadSobre1000)
	})

	t.Run("Población total incorrecta", func(t *testing.T) {
		_, err := segurasenior.NewDatosPoblacion(1000, 500, 501, 30.5, 25.0, 15.0, 50, 20, 50.0, 20.0)
		assert.Error(t, err)
	})

	t.Run("Porcentaje de menores de 20 incorrecto", func(t *testing.T) {
		_, err := segurasenior.NewDatosPoblacion(1000, 500, 500, 30.5, 101.0, 15.0, 50, 20, 50.0, 20.0)
		assert.Error(t, err)
	})
	
	t.Run("Porcentaje de mayores de 65 incorrecto", func(t *testing.T) {
		_, err := segurasenior.NewDatosPoblacion(1000, 500, 500, 30.5, 25.0, 101.0, 50, 20, 50.0, 20.0)
		assert.Error(t, err)
	})

	t.Run("Nacimientos mayor que población total", func(t *testing.T) {
		_, err := segurasenior.NewDatosPoblacion(1000, 500, 500, 30.5, 25.0, 15.0, 1001, 20, 50.0, 20.0)
		assert.Error(t, err)
	})

	t.Run("Defunciones mayor que población total", func(t *testing.T) {
		_, err := segurasenior.NewDatosPoblacion(1000, 500, 500, 30.5, 25.0, 15.0, 50, 1001, 50.0, 20.0)
		assert.Error(t, err)
	})

}