package segurasenior

import "errors"

type Poblacion struct {
	Nombre_pueblo       string
	Poblacion_total     int
	Tasa_natalidad      float32
	Tasa_envejecimiento float32
	Tasa_mortalidad     float32
}

func NewPoblacion(nom string, poblacion int, tasan float32, tasae float32, tasam float32) (*Poblacion, error) {
	if (tasan < 0.0) || (tasan > 1000.0) || (tasae < 0.0) || (tasae > 1000.0) || (tasam < 0.0) || (tasam > 1000.0) {
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
