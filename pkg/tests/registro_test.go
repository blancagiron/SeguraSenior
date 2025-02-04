package segurasenior

import (
	"SeguraSenior/pkg/segurasenior"
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

		datos, err := segurasenior.NewDatosPoblacion( PoblacionTotalOrcera,HombresOrcera,MujeresOrcera,EdadMediaOrcera,Menor20Orcera,Mayor65Orcera,NacimientosOrcera,DefuncionesOrcera,TasaNatalidadSobre1000Orcera,TasaMortalidadSobre1000Orcera)
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
