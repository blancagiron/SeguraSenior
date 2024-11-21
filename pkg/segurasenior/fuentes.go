package segurasenior

import "errors"

type Formato int

const (
	html = iota
	pdf
	excel
)

func (f Formato) String() string {
	switch f {
	case html:
		return "html"
	case pdf:
		return "pdf"
	case excel:
		return "excel"
	}
	return "unknown"
}

type Fuente struct {
	formato_fuente Formato
	direccion      string
}

func NewFuente(ext Formato, dir string) (*Fuente, error) {
	if ext != html && ext != pdf && ext != excel {
		return nil, errors.New("Formato indefinido")
	}
	return &Fuente{
		formato_fuente: ext,
		direccion:      dir,
	}, nil
}
