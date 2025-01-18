package segurasenior

import (
	"testing"
	"SeguraSenior/pkg/segurasenior"
	"github.com/stretchr/testify/assert"
)

func TestCrearRegistroDesdeDatos(t *testing.T) {
	t.Run("Crear registro con población creciente", func(t *testing.T) {
		identificador := segurasenior.IdentificadorDatos{
			NombrePoblacion: "Pueblo Feliz",
			FechaDeDatos: segurasenior.FechaObtencionDeDatos{
				Dia:  1,
				Mes:  1,
				Anio: 2023,
			},
		}

		datos := segurasenior.DatosPoblacion{
			PoblacionTotal:          1000,
			Hombres:                 500,
			Mujeres:                 500,
			EdadMedia:               30.5,
			PorcentajeMenora20:      25.0,
			PorcentajeMayora65:      15.0,
			Nacimientos:             50,
			Defunciones:             20,
			TasaNatalidadSobre1000:  50.0,
			TasaMortalidadSobre1000: 20.0,
		}
		

		registro, err := segurasenior.CrearRegistroDesdeDatos(identificador, datos)


		assert.NoError(t, err)
		assert.Equal(t, segurasenior.Creciente, registro.EstadoDeLaPoblacion)
		assert.Contains(t, registro.EstadisticasPoblacion, identificador)
		assert.Equal(t, datos, registro.EstadisticasPoblacion[identificador])
	})

	t.Run("Crear registro con población decreciente", func(t *testing.T) {
		identificador := segurasenior.IdentificadorDatos{
			NombrePoblacion: "Pueblo Triste",
			FechaDeDatos: segurasenior.FechaObtencionDeDatos{
				Dia:  1,
				Mes:  1,
				Anio: 2023,
			},
		}

		datos := segurasenior.DatosPoblacion{
			PoblacionTotal:          1000,
			Hombres:                 500,
			Mujeres:                 500,
			EdadMedia:               45.2,
			PorcentajeMenora20:      15.0,
			PorcentajeMayora65:      30.0,
			Nacimientos:             10,
			Defunciones:             50,
			TasaNatalidadSobre1000:  10.0,
			TasaMortalidadSobre1000: 50.0,
		}

		registro, err := segurasenior.CrearRegistroDesdeDatos(identificador, datos)

		assert.NoError(t, err)
		assert.Equal(t, segurasenior.Decreciente, registro.EstadoDeLaPoblacion)
		assert.Contains(t, registro.EstadisticasPoblacion, identificador)
		assert.Equal(t, datos, registro.EstadisticasPoblacion[identificador])
	})
}