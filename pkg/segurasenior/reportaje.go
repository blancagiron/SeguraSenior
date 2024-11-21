package segurasenior

type Reportaje struct {
	Nombre              string
	Tendencia_reportaje Tendencia
	Fuentes             []Fuente
}

func NewRepo(nom string, ten Tendencia, orig []Fuente) (*Reportaje, error) {
	return &Reportaje{
		Nombre:              nom,
		Tendencia_reportaje: ten,
		Fuentes:             orig,
	}, nil
}
