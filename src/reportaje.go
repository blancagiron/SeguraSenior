package segurasenior

type Reportaje struct {
	nombre        string
	poblacion     int
	retransmision string
	Fuente        []Fuente
}

func AniadirRepo(nom string, pobla int, medio string, orig []Fuente) *Reportaje {
	return &Reportaje{
		nombre:        nom,
		poblacion:     pobla,
		retransmision: medio,
	}
}
