
package segurasenior

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

