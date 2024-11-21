package segurasenior

type Reportaje struct {
	Nombre              string
	Tendencia_reportaje Tendencia
	Fuentes             []Fuente
}

func AniadirRepo(nom string, ten Tendencia, orig []Fuente) *Reportaje {
	return &Reportaje{
		Nombre:              nom,
		Tendencia_reportaje: ten,
		Fuentes:             orig,
	}
}
