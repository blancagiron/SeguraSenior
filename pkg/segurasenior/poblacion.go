package segurasenior

import "errors"

type Datos_poblacion struct {
	Nombre_pueblo       string
	Poblacion_total     uint16
	Poblacion_joven     uint16
	Poblacion_adulta    uint16
	Poblacion_anciana   int
	Tasa_natalidad      float64
	Tasa_envejecimiento float64
	Tasa_mortalidad     float64
}

func NewPoblacion(nom string, poblacion uint16, poblacionjoven uint16, poblacionmedia uint16, poblacionmayor int, tasan float64, tasae float64, tasam float64) (*Datos_poblacion, error) {
	if nom == "" {
		return nil, errors.New("nombre no puede estar vacío")
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
