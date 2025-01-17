package segurasenior

import (
	"SeguraSenior/pkg/segurasenior"
	"math"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLecturaCorrectaDatosDesdeJSON(t *testing.T) {

	t.Run("Lectura Correcta - Caso Válido", func(t *testing.T) {
		nombreArchivo := "../testdata/data.json"
		nombrePoblacion := "Orcera"

		const (
			poblacionTotal    uint32  = 1735
			hombres           uint32  = 868
			mujeres           uint32  = 867
			edadMedia         float32 = 42.2
			porcentajeMenor20 float64 = 15.6
			porcentajeMayor65 float64 = 22.1
			nacimientos       uint32  = 6
			defunciones       uint32  = 29
		)

		tasaNatalidadSobre1000 := math.Round((float64(nacimientos)/float64(poblacionTotal))*1000*100) / 100
		tasaMortalidadSobre1000 := math.Round((float64(defunciones)/float64(poblacionTotal))*1000*100) / 100

		expectedIdentificador := segurasenior.IdentificadorDatos{
			NombrePoblacion: nombrePoblacion,
			FechaDeDatos: segurasenior.FechaObtencionDeDatos{
				Dia:  1,
				Mes:  1,
				Anio: 2023,
			},
		}

		expectedDatos := segurasenior.DatosPoblacion{
			PoblacionTotal:          poblacionTotal,
			Hombres:                 hombres,
			Mujeres:                 mujeres,
			EdadMedia:               edadMedia,
			PorcentajeMenora20:      porcentajeMenor20,
			PorcentajeMayora65:      porcentajeMayor65,
			Nacimientos:             nacimientos,
			Defunciones:             defunciones,
			TasaNatalidadSobre1000:  tasaNatalidadSobre1000,
			TasaMortalidadSobre1000: tasaMortalidadSobre1000,
		}

		identificador, datosExtraidos, err := segurasenior.LeerDatosDesdeJSON(nombreArchivo, nombrePoblacion)

		assert.NoError(t, err)
		assert.Equal(t, expectedIdentificador, identificador)
		assert.Equal(t, expectedDatos, datosExtraidos)
	})
	t.Run("Archivo No Encontrado", func(t *testing.T) {
		nombreArchivo := "../testdata/no_existe.json"
		nombrePoblacion := "Orcera"

		_, _, err := segurasenior.LeerDatosDesdeJSON(nombreArchivo, nombrePoblacion)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "no such file or directory")
	})

	t.Run("Población No Encontrada", func(t *testing.T) {
		nombreArchivo := "../testdata/data.json"
		nombrePoblacion := "Granada"

		_, _, err := segurasenior.LeerDatosDesdeJSON(nombreArchivo, nombrePoblacion)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "población '"+nombrePoblacion+"' no encontrada en el archivo")
	})

	t.Run("JSON inválido", func(t *testing.T) {
		nombreArchivo := "../testdata/data_invalido.json"
		nombrePoblacion := "Orcera"

		_, _, err := segurasenior.LeerDatosDesdeJSON(nombreArchivo, nombrePoblacion)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid character")
	})

	

}
