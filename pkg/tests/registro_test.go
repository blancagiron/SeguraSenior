package segurasenior

import (
	"SeguraSenior/pkg/segurasenior"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
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

		datos, err := segurasenior.NewDatosPoblacion(1000,500,500,30.5,25.0,15.0,50,20,50.0,20.0)
		assert.NoError(t, err)

		registro, err := segurasenior.CrearRegistroDesdeDatos(identificador, *datos)

		assert.NoError(t, err)
		assert.Equal(t, segurasenior.Creciente, registro.EstadoDeLaPoblacion)
		assert.Contains(t, registro.EstadisticasPoblacion, identificador)
		assert.Equal(t, *datos, registro.EstadisticasPoblacion[identificador])
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

		datos, err := segurasenior.NewDatosPoblacion(1000,500,500,45.2,15.0,30.0,10,50,10.0,50.0)
		assert.NoError(t, err)

		registro, err := segurasenior.CrearRegistroDesdeDatos(identificador, *datos)

		assert.NoError(t, err)
		assert.Equal(t, segurasenior.Decreciente, registro.EstadoDeLaPoblacion)
		assert.Contains(t, registro.EstadisticasPoblacion, identificador)
		assert.Equal(t, *datos, registro.EstadisticasPoblacion[identificador])
	})
}

func TestAgregarRegistro(t *testing.T) {
	t.Run("Agregar un registro nuevo", func(t *testing.T) {
		registro := segurasenior.RegistroDemografico{
			EstadisticasPoblacion: make(map[segurasenior.IdentificadorDatos]segurasenior.DatosPoblacion),
			EstadoDeLaPoblacion:   segurasenior.Creciente,
		}

		identificador := segurasenior.IdentificadorDatos{
			NombrePoblacion: "Pueblo Ejemplo",
			FechaDeDatos: segurasenior.FechaObtencionDeDatos{
				Dia:  1,
				Mes:  1,
				Anio: 2025,
			},
		}

		datos, err := segurasenior.NewDatosPoblacion(1000, 500, 500, 30.0, 25.0, 15.0, 50, 20, 50.0, 20.0)
		assert.NoError(t, err)

		err = registro.AgregarRegistro(identificador, *datos)

		assert.NoError(t, err)
		assert.Len(t, registro.EstadisticasPoblacion, 1)
		assert.Equal(t, segurasenior.Creciente, registro.EstadoDeLaPoblacion)
	})

	t.Run("Intentar agregar un registro para la misma fecha", func(t *testing.T) {

		registro := segurasenior.RegistroDemografico{
			EstadisticasPoblacion: make(map[segurasenior.IdentificadorDatos]segurasenior.DatosPoblacion),
			EstadoDeLaPoblacion:   segurasenior.Creciente,
		}

		identificador := segurasenior.IdentificadorDatos{
			NombrePoblacion: "Pueblo Ejemplo",
			FechaDeDatos: segurasenior.FechaObtencionDeDatos{
				Dia:  1,
				Mes:  1,
				Anio: 2025,
			},
		}
		datos, err := segurasenior.NewDatosPoblacion(1600, 700, 900, 30.0, 25.0, 15.0, 10, 40, 10.0, 40.0)
		assert.NoError(t, err)

		err = registro.AgregarRegistro(identificador, *datos)
		assert.NoError(t, err)
		assert.Len(t, registro.EstadisticasPoblacion, 1)

		err = registro.AgregarRegistro(identificador, *datos)
		assert.Error(t, err)

		expectedError := fmt.Sprintf("ya existe un registro para '%s' en la fecha %d/%d/%d",
			identificador.NombrePoblacion, identificador.FechaDeDatos.Dia,
			identificador.FechaDeDatos.Mes, identificador.FechaDeDatos.Anio)
		assert.Contains(t, err.Error(), expectedError)

		assert.Len(t, registro.EstadisticasPoblacion, 1)
	})

}
