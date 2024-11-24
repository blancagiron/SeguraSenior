package segurasenior

type Reportaje struct {
	nombre          string
	tendencias      []Tendencia
	datos_poblacion []Poblacion
}

func NewRepo(nom string, ten []Tendencia, datospob []Poblacion) (*Reportaje, error) {
	return &Reportaje{
		nombre:          nom,
		tendencias:      ten,
		datos_poblacion: datospob,
	}, nil
}
