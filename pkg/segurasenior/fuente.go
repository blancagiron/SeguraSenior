package segurasenior

type Fuente struct {
	Direccion string
}

func NewFuente(dir string) (*Fuente, error) {

	return &Fuente{
		Direccion: dir,
	}, nil
}
