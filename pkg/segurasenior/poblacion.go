package segurasenior

import "errors"

type Poblacion struct {
	Nombre_pueblo       string
	Poblacion_total     int
	Tasa_natalidad      int
	Tasa_envejecimiento int
	Tasa_mortalidad     int
}

func NewPoblacion(nom string, poblacion int, tasan int, tasae int, tasam int) (*Poblacion, error) {
	if (tasan < 0) || (tasan > 1000) || (tasae < 0) || (tasae > 1000) || (tasam < 0) || (tasam > 1000) {
		return nil, errors.New("valor de tasas erróneo, debe estar comprendido entre 0 y 1000")
	}
	return &Poblacion{
		Nombre_pueblo:       nom,
		Poblacion_total:     poblacion,
		Tasa_natalidad:      tasan,
		Tasa_envejecimiento: tasae,
		Tasa_mortalidad:     tasam,
	}, nil
}
