package segurasenior

import "errors"

type Datos_poblacion struct {
	Nombre_pueblo       string
	Poblacion_total     uint32
	Poblacion_joven     uint32
	Poblacion_adulta    uint32
	Poblacion_anciana   uint32
	Tasa_natalidad      float64
	Tasa_envejecimiento float64
	Tasa_mortalidad     float64
}

func NewPoblacion(nom string, poblacion uint32, poblacion_menor_18 uint32, poblacion_entre_18_y_64 uint32, poblacion_mayor_64 uint32,
	tasa_natalidad_sobre_1000 float64, tasa_envejecimiento_sobre_1000 float64, tasa_mortalidad float64) (*Datos_poblacion, error) {
	if nom == "" {
		return nil, errors.New("nombre no puede estar vacío")
	}

	if (tasa_natalidad_sobre_1000 < 0.0) || (tasa_natalidad_sobre_1000 > 1000.0) {
		return nil, errors.New("valor de tasa de natalidada erróneo, debe estar comprendido entre 0 y 1000")
	}
	if (tasa_envejecimiento_sobre_1000 < 0.0) || (tasa_envejecimiento_sobre_1000 > 1000.0) {
		return nil, errors.New("valor de tasa de envejecimiento erróneo, debe estar comprendido entre 0 y 1000")
	}
	if (tasa_mortalidad < 0.0) || (tasa_mortalidad > 1000.0) {
		return nil, errors.New("valor de tasa de mortalidad erróneo, debe estar comprendido entre 0 y 1000")
	}
	return &Datos_poblacion{
		Nombre_pueblo:       nom,
		Poblacion_total:     poblacion,
		Poblacion_joven:     poblacion_menor_18,
		Poblacion_adulta:    poblacion_entre_18_y_64,
		Poblacion_anciana:   poblacion_mayor_64,
		Tasa_natalidad:      tasa_natalidad_sobre_1000,
		Tasa_envejecimiento: tasa_envejecimiento_sobre_1000,
		Tasa_mortalidad:     tasa_mortalidad,
	}, nil
}
