package segurasenior

type Reportaje struct {
	nombre        string
	poblacion     int
	retransmision string
}

func AniadirRepo(nom string, pobla int, medio string) *Reportaje {
	return &Reportaje{
		nombre:        nom,
		poblacion:     pobla,
		retransmision: medio,
	}
}
