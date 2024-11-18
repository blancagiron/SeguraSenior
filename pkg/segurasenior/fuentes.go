package segurasenior

type Fuente struct {
	Formato   string //Nombre reducido de la página o especificación de la extensión para entendimiento del origen de la fuente
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
