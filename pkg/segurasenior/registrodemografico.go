package segurasenior

type EstadoPoblacion string

const (
	Decreciente EstadoPoblacion = "decreciente"
	Creciente   EstadoPoblacion = "creciente"
)

type IdentificadorDatos struct {
	NombrePoblacion string
	FechaDeDatos    FechaObtencionDeDatos
}

type RegistroDemografico struct {
	EstadisticasPoblacion map[IdentificadorDatos]DatosPoblacion
	EstadoDeLaPoblacion   EstadoPoblacion
}

func NewRegistroDemografico(datosPoblacion DatosPoblacion, estadoPoblacion EstadoPoblacion) (*RegistroDemografico, error) {
	return &RegistroDemografico{
		EstadisticasPoblacion: datosPoblacion,
		EstadoDeLaPoblacion:   estadoPoblacion,
	}, nil
}

type RegistrosDemografia map[string][]RegistroDemografico
