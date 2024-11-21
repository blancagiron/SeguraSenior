package segurasenior

type Fuente struct {
	Formato   string
	Direccion string
	Poblacion Tendencia
}

func NewFuente(ext string, dir string, ten Tendencia) (*Fuente, error) {
	return &Fuente{
		Formato:   ext,
		Direccion: dir,
		Poblacion: ten,
	}, nil
}
