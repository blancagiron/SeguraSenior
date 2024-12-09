package segurasenior

type EstadoPoblacion string

const (
	Decreciente EstadoPoblacion = "decreciente"
	Creciente   EstadoPoblacion = "creciente"
)

type RegistroDemografico struct {
	EstadisticasPoblacion DatosPoblacion
	EstadoDeLaPoblacion   EstadoPoblacion
}

func NewRegistroDemografico(datosPoblacion DatosPoblacion, estadoPoblacion EstadoPoblacion) (*RegistroDemografico, error) {
	return &RegistroDemografico{
		EstadisticasPoblacion: datosPoblacion,
		EstadoDeLaPoblacion:   estadoPoblacion,
	}, nil
}

type RegistrosDemografia map[string][]RegistroDemografico
