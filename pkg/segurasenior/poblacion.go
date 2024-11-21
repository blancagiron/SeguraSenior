package segurasenior

import "errors"

type Poblacion struct {
	nombre_pueblo       string
	poblacion_total     int
	tasa_natalidad      float32
	tasa_envejecimiento float32
	tasa_mortalidad     float32
}

func NewPoblacion(nom string, poblacion int, tasan float32, tasae float32, tasam float32) (*Poblacion, error) {
	if (tasan < 0.0) || (tasan > 1000.0) {
		return nil, errors.New("valor de tasa de natalidada erróneo, debe estar comprendido entre 0 y 1000")
	}
	if (tasae < 0.0) || (tasae > 1000.0) {
		return nil, errors.New("valor de tasa de envejecimiento erróneo, debe estar comprendido entre 0 y 1000")
	}
	if (tasam < 0.0) || (tasam > 1000.0) {
		return nil, errors.New("valor de tasa de mortalidad erróneo, debe estar comprendido entre 0 y 1000")
	}
	return &Poblacion{
		nombre_pueblo:       nom,
		poblacion_total:     poblacion,
		tasa_natalidad:      tasan,
		tasa_envejecimiento: tasae,
		tasa_mortalidad:     tasam,
	}, nil
}
