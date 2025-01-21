package segurasenior

import (
	"SeguraSenior/pkg/segurasenior"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCrearRegistroDesdeDatos(test *testing.T) {
	test.Run("Crear registro con población creciente", func(test *testing.T) {
		identificador := segurasenior.IdentificadorDatos{
			NombrePoblacion: "Pueblo Feliz",
			FechaDeDatos: segurasenior.FechaObtencionDeDatos{
				Dia:  1,
				Mes:  1,
				Anio: 2023,
			},
		}

		datos, err := segurasenior.NewDatosPoblacion(1000,500,500,30.5,25.0,15.0,50,20,50.0,20.0)
		assert.NoError(test, err)

		registro, err := segurasenior.CrearRegistroDesdeDatos(identificador, *datos)

		assert.NoError(test, err)
		assert.Equal(test, segurasenior.Creciente, registro.EstadoDeLaPoblacion)
		assert.Contains(test, registro.EstadisticasPoblacion, identificador)
		assert.Equal(test, *datos, registro.EstadisticasPoblacion[identificador])
	})

	test.Run("Crear registro con población decreciente", func(test *testing.T) {
		identificador := segurasenior.IdentificadorDatos{
			NombrePoblacion: "Pueblo Triste",
			FechaDeDatos: segurasenior.FechaObtencionDeDatos{
				Dia:  1,
				Mes:  1,
				Anio: 2023,
			},
		}

		datos, err := segurasenior.NewDatosPoblacion(1000,500,500,45.2,15.0,30.0,10,50,10.0,50.0)
		assert.NoError(test, err)

		registro, err := segurasenior.CrearRegistroDesdeDatos(identificador, *datos)

		assert.NoError(test, err)
		assert.Equal(test, segurasenior.Decreciente, registro.EstadoDeLaPoblacion)
		assert.Contains(test, registro.EstadisticasPoblacion, identificador)
		assert.Equal(test, *datos, registro.EstadisticasPoblacion[identificador])
	})
}

func TestAgregarRegistro(test *testing.T) {
	test.Run("Agregar un registro nuevo", func(test *testing.T) {
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
		assert.NoError(test, err)

		err = registro.AgregarRegistro(identificador, *datos)

		assert.NoError(test, err)
		assert.Len(test, registro.EstadisticasPoblacion, 1)
		assert.Equal(test, segurasenior.Creciente, registro.EstadoDeLaPoblacion)
	})

	test.Run("Intentar agregar un registro para la misma fecha", func(test *testing.T) {

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
		assert.NoError(test, err)

		err = registro.AgregarRegistro(identificador, *datos)
		assert.NoError(test, err)
		assert.Len(test, registro.EstadisticasPoblacion, 1)

		err = registro.AgregarRegistro(identificador, *datos)
		assert.Error(test, err)

		expectedError := fmt.Sprintf("ya existe un registro para '%s' en la fecha %d/%d/%d",
			identificador.NombrePoblacion, identificador.FechaDeDatos.Dia,
			identificador.FechaDeDatos.Mes, identificador.FechaDeDatos.Anio)
		assert.Contains(test, err.Error(), expectedError)

		assert.Len(test, registro.EstadisticasPoblacion, 1)
	})

}
