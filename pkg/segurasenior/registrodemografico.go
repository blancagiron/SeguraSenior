package segurasenior

type EstadoPoblacion string

const (
	Decreciente EstadoPoblacion = "decreciente"
	Creciente   EstadoPoblacion = "creciente"
)

type RegistroDemografico struct {
	EstadisticasPoblacion DatosPoblacion
	TendenciasPoblacion   EstadoPoblacion
}

func NewRegistroDemografico(poblacion DatosPoblacion, tendencia EstadoPoblacion) (*RegistroDemografico, error) {
	return &RegistroDemografico{
		EstadisticasPoblacion: poblacion,
		TendenciasPoblacion:   tendencia,
	}, nil
}

type Registros_demografia map[string][]RegistroDemografico
