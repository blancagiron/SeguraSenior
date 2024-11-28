package segurasenior

import "errors"

type Estado int

const (
	Decreciente Estado = iota
	Creciente
)

func (e Estado) String() string {
	switch e {
	case Decreciente:
		return "decreciente"
	case Creciente:
		return "creciente"
	}
	return "unknown"
}

type Registro_demografico struct {
	Estadisticas_poblacion []Datos_poblacion
	Tendencias_poblacion   Estado
	Fuentes                []Fuente
}

func NewRegistro(poblaciones []Datos_poblacion, fuentes []Fuente, tendencia Tendencia) (*Registro_demografico, error) {
	if len(poblaciones) == 0 {
		return nil, errors.New("el registro debe incluir algún dato sobre la población")
	}
	return &Registro_demografico{
		Estadisticas_poblacion: poblaciones,
		Tendencias_poblacion:   tendencia,
		Fuentes:                fuentes,
	}, nil
}
