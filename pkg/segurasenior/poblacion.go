package segurasenior

import "errors"

type Datos_poblacion struct {
	Nombre_pueblo       string
	Poblacion_total     int
	Poblacion_joven     int
	Poblacion_adulta    int
	Poblacion_anciana   int
	Tasa_natalidad      float64
	Tasa_envejecimiento float64
	Tasa_mortalidad     float64
}

func NewPoblacion(nom string, poblacion int, poblacionjoven int, poblacionmedia int, poblacionmayor int, tasan float64, tasae float64, tasam float64) (*Datos_poblacion, error) {
	if nom == "" {
		return nil, errors.New("nombre no puede estar vacío")
	}
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
	return &Datos_poblacion{
		Nombre_pueblo:       nom,
		Poblacion_total:     poblacion,
		Poblacion_joven:     poblacionjoven,
		Poblacion_adulta:    poblacionmedia,
		Poblacion_anciana:   poblacionmayor,
		Tasa_natalidad:      tasan,
		Tasa_envejecimiento: tasae,
		Tasa_mortalidad:     tasam,
	}, nil
}
