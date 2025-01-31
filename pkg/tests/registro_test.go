package segurasenior

import (
	"SeguraSenior/pkg/segurasenior"
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
		if err != nil {
			test.Errorf("No se esperaba error, pero se obtuvo: %v", err)
		}

		registro, err := segurasenior.CrearRegistroDesdeDatos(identificador, *datos)

		if err != nil {
			test.Errorf("No se esperaba error, pero se obtuvo: %v", err)
		}
		if registro.EstadoDeLaPoblacion != segurasenior.Creciente {
			test.Errorf("Se esperaba estado Creciente, pero se obtuvo: %v", registro.EstadoDeLaPoblacion)
		}
		if _, ok := registro.EstadisticasPoblacion[identificador]; !ok {
			test.Errorf("Se esperaba que el identificador estuviera presente en el registro")
		}
		if registro.EstadisticasPoblacion[identificador] != *datos {
			test.Errorf("Los datos de población no coinciden con los esperados")
		}
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
		if err != nil {
			test.Errorf("No se esperaba error, pero se obtuvo: %v", err)
		}

		registro, err := segurasenior.CrearRegistroDesdeDatos(identificador, *datos)

		if err != nil {
			test.Errorf("No se esperaba error, pero se obtuvo: %v", err)
		}
		if registro.EstadoDeLaPoblacion != segurasenior.Decreciente {
			test.Errorf("Se esperaba estado Decreciente, pero se obtuvo: %v", registro.EstadoDeLaPoblacion)
		}
		if _, ok := registro.EstadisticasPoblacion[identificador]; !ok {
			test.Errorf("Se esperaba que el identificador estuviera presente en el registro")
		}
		if registro.EstadisticasPoblacion[identificador] != *datos {
			test.Errorf("Los datos de población no coinciden con los esperados")
		}
	})
}
