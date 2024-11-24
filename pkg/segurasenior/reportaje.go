package segurasenior

type Reportaje struct {
	nombre          string
	tendencias      []Tendencia
	datos_poblacion []Poblacion
	fuentes         []Fuente
}

func NewRepo(nom string, ten []Tendencia, datospob []Poblacion, origen []Fuente) (*Reportaje, error) {
	return &Reportaje{
		nombre:          nom,
		tendencias:      ten,
		datos_poblacion: datospob,
		fuentes:         origen,
	}, nil
}
