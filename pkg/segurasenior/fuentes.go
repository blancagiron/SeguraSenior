package segurasenior

import "errors"

type Formato int

const (
	Html = iota
	Pdf
	Excel
)

func (f Formato) String() string {
	switch f {
	case Html:
		return "html"
	case Pdf:
		return "pdf"
	case Excel:
		return "excel"
	}
	return "unknown"
}

type Fuente struct {
	Formato_fuente Formato
	Direccion      string
}

func NewFuente(ext Formato, dir string) (*Fuente, error) {
	if ext != Html && ext != Pdf && ext != Excel {
		return nil, errors.New("Formato indefinido")
	}
	return &Fuente{
		Formato_fuente: ext,
		Direccion:      dir,
	}, nil
}
