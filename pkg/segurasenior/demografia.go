package segurasenior

type Estado string

const (
	Decreciente Estado = "decreciente"
	Creciente   Estado = "creciente"
)

type RegistroDemografico struct {
	EstadisticasPoblacion DatosPoblacion
	TendenciasPoblacion   Estado
}

func NewRegistro(poblacion DatosPoblacion, tendencia Estado) (*RegistroDemografico, error) {
	return &RegistroDemografico{
		EstadisticasPoblacion: poblacion,
		TendenciasPoblacion:   tendencia,
	}, nil
}

type Registros_demografia map[string][]RegistroDemografico
