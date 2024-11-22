package segurasenior

type Reportaje struct {
	nombre              string
	tendencia_reportaje Tendencia
}

func NewRepo(nom string, ten Tendencia) (*Reportaje, error) {
	return &Reportaje{
		nombre:              nom,
		tendencia_reportaje: ten,
	}, nil
}
