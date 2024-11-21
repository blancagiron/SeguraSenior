package segurasenior

type Reportaje struct {
	nombre              string
	tendencia_reportaje Tendencia
	fuentes             []Fuente
}

func NewRepo(nom string, ten Tendencia, orig []Fuente) (*Reportaje, error) {
	return &Reportaje{
		nombre:              nom,
		tendencia_reportaje: ten,
		fuentes:             orig,
	}, nil
}
