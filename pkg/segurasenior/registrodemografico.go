package segurasenior

import (
	"fmt"
)

type EstadoPoblacion string


const (
	Decreciente EstadoPoblacion = "decreciente"
	Creciente   EstadoPoblacion = "creciente"
)

type RegistroDemografico struct {
	EstadisticasPoblacion map[IdentificadorDatos]DatosPoblacion
	EstadoDeLaPoblacion   EstadoPoblacion
}

func ObtenerEstadoPoblacion(datos DatosPoblacion) EstadoPoblacion {
	if datos.TasaMortalidadSobre1000 > datos.TasaNatalidadSobre1000 {
		return Decreciente
	}
	return Creciente
}

func CrearRegistroDesdeDatos(identificador IdentificadorDatos, datos DatosPoblacion) (*RegistroDemografico, error) {
	estado := ObtenerEstadoPoblacion(datos)

	return &RegistroDemografico{
		EstadisticasPoblacion: map[IdentificadorDatos]DatosPoblacion{identificador: datos},
		EstadoDeLaPoblacion:   estado,
	}, nil
}

func ValidarRegistroExiste(estadisticas map[IdentificadorDatos]DatosPoblacion, identificador IdentificadorDatos) error {
	if _, existe := estadisticas[identificador]; existe {
		return fmt.Errorf("ya existe un registro para '%s' en la fecha %d/%d/%d",
			identificador.NombrePoblacion, identificador.FechaDeDatos.Dia,
			identificador.FechaDeDatos.Mes, identificador.FechaDeDatos.Anio)
	}
	return nil
}