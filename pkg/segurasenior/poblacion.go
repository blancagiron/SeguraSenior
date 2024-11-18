package segurasenior

import "errors"

type Poblacion struct {
	Demografia     string
	PoblacionTotal int
	Tasanat        int //medido en tantos por mil
	Tasaenv        int //medido en tantos por mil
	Tasamort       int //medido en tantos por mil
}

func NewPoblacion(nom string, poblacion int, tasan int, tasae int, tasam int) (*Poblacion, error) {
	if (tasan < 0) || (tasan > 1000) || (tasae < 0) || (tasae > 1000) || (tasam < 0) || (tasam > 1000) {
		return nil, errors.New("valor de tasas erróneo, debe estar comprendido entre 0 y 1000")
	}
	return &Poblacion{
		Demografia:     nom,
		PoblacionTotal: poblacion,
		Tasanat:        tasan,
		Tasaenv:        tasae,
		Tasamort:       tasam,
	}, nil
}
