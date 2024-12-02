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

func NewRegistroDemografico(datospoblacion DatosPoblacion, tendencia EstadoPoblacion) (*RegistroDemografico, error) {
	return &RegistroDemografico{
		EstadisticasPoblacion: datospoblacion,
		EstadoDeLaPoblacion:   tendencia,
	}, nil
}

type Registros_demografia map[string][]RegistroDemografico
