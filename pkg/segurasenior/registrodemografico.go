package segurasenior

import "fmt"

type EstadoPoblacion string

const (
	Decreciente EstadoPoblacion = "decreciente"
	Creciente   EstadoPoblacion = "creciente"
)

type RegistroDemografico struct {
	EstadisticasPoblacion map[IdentificadorDatos]DatosPoblacion
	EstadoDeLaPoblacion   EstadoPoblacion
}

func CrearRegistroDesdeDatos(identificador IdentificadorDatos, datos DatosPoblacion) (*RegistroDemografico, error) {
	estado := Creciente
	if datos.TasaMortalidadSobre1000 > datos.TasaNatalidadSobre1000 {
		estado = Decreciente
	}

	return &RegistroDemografico{
		EstadisticasPoblacion: map[IdentificadorDatos]DatosPoblacion{identificador: datos},
		EstadoDeLaPoblacion:   estado,
	}, nil
}

func (registro *RegistroDemografico) AgregarRegistro(identificador IdentificadorDatos, datos DatosPoblacion) error {
	if _, existe := registro.EstadisticasPoblacion[identificador]; existe {
		return fmt.Errorf("ya existe un registro para '%s' en la fecha %d/%d/%d",
			identificador.NombrePoblacion, identificador.FechaDeDatos.Dia,
			identificador.FechaDeDatos.Mes, identificador.FechaDeDatos.Anio)
	}

	registro.EstadisticasPoblacion[identificador] = datos
	if datos.TasaMortalidadSobre1000 > datos.TasaNatalidadSobre1000 {
		registro.EstadoDeLaPoblacion = Decreciente
	} else {
		registro.EstadoDeLaPoblacion = Creciente
	}
	return nil
}