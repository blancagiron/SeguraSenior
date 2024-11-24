package segurasenior

type Reportaje struct {
	nombre     string
	tendencias []Tendencia
}

func NewRepo(nom string, ten []Tendencia) (*Reportaje, error) {
	return &Reportaje{
		nombre:     nom,
		tendencias: ten,
	}, nil
}
