package segurasenior

type Fuente struct {
	Origen    string //Nombre reducido de la página o especificación de la extensión para entendimiento del origen de la fuente
	Direccion string
	Poblacion Tendencia
}

func NewFuente(nom string, dir string, ten Tendencia) (*Fuente, error) {
	return &Fuente{
		Origen:    nom,
		Direccion: dir,
		Poblacion: ten,
	}, nil
}
