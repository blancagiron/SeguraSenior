package segurasenior

type Estado int

const (
	Decreciente Estado = iota
	Creciente
)

type Registro_demografico struct {
	Estadisticas_poblacion Datos_poblacion
	Tendencias_poblacion   Estado
	Direcciones_fuentes    []string
}

func NewRegistro(poblacion Datos_poblacion, fuentes []string, tendencia Estado) (*Registro_demografico, error) {
	return &Registro_demografico{
		Estadisticas_poblacion: poblacion,
		Tendencias_poblacion:   tendencia,
		Direcciones_fuentes:    fuentes,
	}, nil
}

type Registros_demografia map[*Registro_demografico][]string
