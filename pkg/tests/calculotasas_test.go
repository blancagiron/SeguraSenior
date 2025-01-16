package segurasenior

import (
	"SeguraSenior/pkg/segurasenior"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalcularTasas(t *testing.T) {

	tests := []struct {
		name                      string
		datosEntrada              segurasenior.DatosPoblacion
		expectedTasaNatalidad     float64
		expectedTasaMortalidad    float64
	}{
		{
			name: "Caso 1: Datos normales",
			datosEntrada: segurasenior.DatosPoblacion{
				PoblacionTotal: 350,
				Nacimientos:    20,
				Defunciones:    10,
			},
			expectedTasaNatalidad: 57.14,
			expectedTasaMortalidad: 28.57,
		},
		{
			name: "Caso 2: Población de 1000, sin nacimientos ni defunciones",
			datosEntrada: segurasenior.DatosPoblacion{
				PoblacionTotal: 1000,
				Nacimientos:    0,
				Defunciones:    0,
			},
			expectedTasaNatalidad: 0.00,
			expectedTasaMortalidad: 0.00,
		},
		
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			
			datos := tt.datosEntrada
			
			datos.CalcularTasas()
	
			assert.InDelta(t, tt.expectedTasaNatalidad, datos.TasaNatalidadSobre1000, 0.01, "Tasa de natalidad incorrecta")
			assert.InDelta(t, tt.expectedTasaMortalidad, datos.TasaMortalidadSobre1000, 0.01, "Tasa de mortalidad incorrecta")
		})
	}
}