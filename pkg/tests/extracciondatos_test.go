package segurasenior

import (
	"SeguraSenior/pkg/segurasenior"
	"testing"
)

const (
	archivoValido   = "../testdata/data.json"
	archivoInvalido = "../testdata/data_invalido.json"
)

func TestLeerDatosPoblacionDesdeArchivo(test *testing.T) {
	test.Run("Creación exitosa de DatosPoblacion a partir de archivo JSON válido", func(test *testing.T) {
		nombrePoblacion := "Orcera"

		datos, err := segurasenior.LeerDatosPoblacionDesdeArchivo(archivoValido, nombrePoblacion)
		if err != nil {
			test.Fatalf("No se esperaba error, pero se obtuvo: %v", err)
		}
		if datos.PoblacionTotal != PoblacionTotalOrcera || datos.Hombres != HombresOrcera || datos.Mujeres != MujeresOrcera {
			test.Errorf("Los valores no coinciden con los datos esperados")
		}

		if datos.TasaNatalidadSobre1000 != TasaNatalidadSobre1000Orcera {
			test.Errorf("La tasa de natalidad calculada es incorrecta: se esperaba 50.0, se obtuvo %v", datos.TasaNatalidadSobre1000)
		}

		if datos.TasaMortalidadSobre1000 != TasaMortalidadSobre1000Orcera {
			test.Errorf("La tasa de mortalidad calculada es incorrecta: se esperaba 25.0, se obtuvo %v", datos.TasaMortalidadSobre1000)
		}
	})

	test.Run("Error al procesar datos JSON inválidos", func(test *testing.T) {
		nombreArchivo := "data_invalido.json"
		nombrePoblacion := "PoblacionInvalida"

		_, err := segurasenior.LeerDatosPoblacionDesdeArchivo(nombreArchivo, nombrePoblacion)
		if err == nil {
			test.Fatal("Se esperaba un error debido a datos inconsistentes, pero no se obtuvo ninguno")
		}
	})
}

func TestLeerIdentificadorDatosDesdeJSON(test *testing.T) {
	test.Run("Creación exitosa de IdentificadorDatos desde archivo JSON válido", func(test *testing.T) {
		nombrePoblacion := "Orcera"
		identificador, err := segurasenior.LeerIdentificadorDatosDesdeJSON(archivoValido, nombrePoblacion)
		if err != nil {
			test.Fatalf("No se esperaba error, pero se obtuvo: %v", err)
		}
		if identificador.NombrePoblacion != "Orcera" {
			test.Errorf("El nombre de la población no coincide: se esperaba 'PoblacionEjemplo', se obtuvo '%s'", identificador.NombrePoblacion)
		}

		if identificador.FechaDeDatos.Anio != 2023 || identificador.FechaDeDatos.Mes != 1 || identificador.FechaDeDatos.Dia != 1 {
			test.Errorf("La fecha no coincide con los datos esperados")
		}
	})

	test.Run("Error al procesar archivo JSON inválido", func(test *testing.T) {
		nombreArchivo := "data_invalido.json"
		nombrePoblacion := "Orcera"

		_, err := segurasenior.LeerIdentificadorDatosDesdeJSON(nombreArchivo, nombrePoblacion)
		if err == nil {
			test.Fatal("Se esperaba un error debido a datos inconsistentes, pero no se obtuvo ninguno")
		}
	})
}