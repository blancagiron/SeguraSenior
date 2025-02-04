package segurasenior

import (
	"SeguraSenior/pkg/segurasenior"
	"fmt"
	"testing"
)

func TestCrearRegistroDesdeDatos(test *testing.T) {
	test.Run("Crear registro con población decreciente", func(test *testing.T) {
		identificador := segurasenior.IdentificadorDatos{
			NombrePoblacion: "Pueblo Feliz",
			FechaDeDatos: segurasenior.FechaObtencionDeDatos{
				Dia:  1,
				Mes:  1,
				Anio: 2023,
			},
		}

		datos, err := segurasenior.NewDatosPoblacion(PoblacionTotalOrcera, HombresOrcera, MujeresOrcera, EdadMediaOrcera, Menor20Orcera, Mayor65Orcera, NacimientosOrcera, DefuncionesOrcera, TasaNatalidadSobre1000Orcera, TasaMortalidadSobre1000Orcera)
		if err != nil {
			test.Errorf("No se esperaba error, pero se obtuvo: %v", err)
		}

		registro, err := segurasenior.CrearRegistroDesdeDatos(identificador, *datos)

		if registro.EstadoDeLaPoblacion != segurasenior.Decreciente {
			test.Errorf("Se esperaba estado Decreciente, pero se obtuvo: %v", registro.EstadoDeLaPoblacion)
		}
		if _, ok := registro.EstadisticasPoblacion[identificador]; !ok {
			test.Errorf("Se esperaba que el identificador estuviera presente en el registro")
		}
		if registro.EstadisticasPoblacion[identificador] != *datos {
			test.Errorf("Los datos de población no coinciden con los esperados")
		}

		if err != nil {
			test.Errorf("No se esperaba error, pero se obtuvo: %v", err)
		}
	})

}

func TestAgregarRegistro(test *testing.T) {
	test.Run("Intentar agregar un registro para la misma fecha", func(test *testing.T) {
		registro := segurasenior.RegistroDemografico{
			EstadisticasPoblacion: make(map[segurasenior.IdentificadorDatos]segurasenior.DatosPoblacion),
			EstadoDeLaPoblacion:   segurasenior.Creciente,
		}

		identificador := segurasenior.IdentificadorDatos{
			NombrePoblacion: "Pueblo Feliz",
			FechaDeDatos: segurasenior.FechaObtencionDeDatos{
				Dia:  1,
				Mes:  1,
				Anio: 2023,
			},
		}

		datos, _ := segurasenior.NewDatosPoblacion(PoblacionTotalOrcera, HombresOrcera, MujeresOrcera, EdadMediaOrcera, Menor20Orcera, Mayor65Orcera, NacimientosOrcera, DefuncionesOrcera, TasaNatalidadSobre1000Orcera, TasaMortalidadSobre1000Orcera)
	
		err := registro.AgregarRegistro(identificador, *datos)
		if err != nil {
			test.Errorf("No se esperaba error al agregar el registro, pero se obtuvo: %v", err)
		}

		err = registro.AgregarRegistro(identificador, *datos)
		if err == nil {
			test.Errorf("Se esperaba un error al intentar agregar un registro duplicado, pero no se obtuvo")
		}

		expectedError := fmt.Sprintf("ya existe un registro para '%s' en la fecha %d/%d/%d",
			identificador.NombrePoblacion, identificador.FechaDeDatos.Dia,
			identificador.FechaDeDatos.Mes, identificador.FechaDeDatos.Anio)

		if err.Error() != expectedError {
			test.Errorf("El mensaje de error esperado era '%s', pero se obtuvo '%s'", expectedError, err.Error())
		}

		if len(registro.EstadisticasPoblacion) != 1 {
			test.Errorf("Se esperaba 1 registro, pero se obtuvieron: %d", len(registro.EstadisticasPoblacion))
		}
	})
}
