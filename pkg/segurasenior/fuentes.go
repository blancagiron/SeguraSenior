package segurasenior

type Fuente struct {
	Formato   string
	Direccion string
}

func NewFuente(ext string, dir string) (*Fuente, error) {
	return &Fuente{
		Formato:   ext,
		Direccion: dir,
	}, nil
}
