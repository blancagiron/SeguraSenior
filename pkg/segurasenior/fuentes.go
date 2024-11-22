package segurasenior

import "errors"

type formato int

const (
	html formato = iota
	pdf
	excel
)

func (f formato) String() string {
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
	formato_fuente formato
	direccion      string
}

func NewFuente(ext formato, dir string) (*Fuente, error) {
	if ext != html && ext != pdf && ext != excel {
		return nil, errors.New("formato indefinido")
	}
	return &Fuente{
		formato_fuente: ext,
		direccion:      dir,
	}, nil
}
