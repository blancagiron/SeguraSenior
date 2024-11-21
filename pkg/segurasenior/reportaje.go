package segurasenior

type Reportaje struct {
	Nombre        string
	Poblacion     int
	Retransmision string
	Fuentes       []Fuente
}

func AniadirRepo(nom string, pobla int, medio string, orig []Fuente) *Reportaje {
	return &Reportaje{
		Nombre:        nom,
		Poblacion:     pobla,
		Retransmision: medio,
		Fuentes:       orig,
	}
}
