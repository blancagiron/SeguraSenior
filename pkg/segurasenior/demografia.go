package segurasenior

type Estado string

const (
	Decreciente Estado = "decreciente"
	Creciente   Estado = "creciente"
)

type Registro_demografico struct {
	Estadisticas_poblacion Datos_poblacion
	Tendencias_poblacion   Estado
}

func NewRegistro(poblacion Datos_poblacion, tendencia Estado) (*Registro_demografico, error) {
	return &Registro_demografico{
		Estadisticas_poblacion: poblacion,
		Tendencias_poblacion:   tendencia,
	}, nil
}

type Registros_demografia map[*Registro_demografico][]string
