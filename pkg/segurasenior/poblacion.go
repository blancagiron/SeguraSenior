package segurasenior

import "errors"

type Poblacion struct {
	nombre_pueblo       string
	poblacion_total     int
	poblacion_joven     int
	poblacion_adulta    int
	poblacion_anciana   int
	tasa_natalidad      float32
	tasa_envejecimiento float32
	tasa_mortalidad     float32
}

func NewPoblacion(nom string, poblacion int, poblacionjoven int, poblacionmedia int, poblacionmayor int, tasan float32, tasae float32, tasam float32) (*Poblacion, error) {
	if poblacion < 0 {
		return nil, errors.New("valor de poblacion total debe ser mayor a 0")
	}
	if (poblacionjoven <= poblacion) || (poblacionjoven < 0) {
		return nil, errors.New("valor de poblacion_joven debe ser menor a la poblacion_total y positiva")
	}
	if (poblacionmedia <= poblacion) || (poblacionmedia < 0) {
		return nil, errors.New("valor de poblacion_adulta debe ser menor a la poblacion_total y positiva")
	}
	if (poblacionmayor <= poblacion) || (poblacionjoven < 0) {
		return nil, errors.New("valor de poblacion_anciana debe ser menor a la poblacion_total y positiva")
	}
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
		poblacion_joven:     poblacionjoven,
		poblacion_adulta:    poblacionmedia,
		poblacion_anciana:   poblacionmayor,
		tasa_natalidad:      tasan,
		tasa_envejecimiento: tasae,
		tasa_mortalidad:     tasam,
	}, nil
}
