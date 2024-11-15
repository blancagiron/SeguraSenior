package segurasenior

type Fuente struct {
	Nombre    string
	Direccion string
}

func NewFuente(nom string, dir string) (*Fuente, error) {
	return &Fuente{
		Nombre:    nom,
		Direccion: dir,
	}, nil
}
