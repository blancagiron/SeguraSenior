package segurasenior

type FormatoFuente string

const (
	html  FormatoFuente = "html"
	pdf   FormatoFuente = "pdf"
	excel FormatoFuente = "excel"
)

type Fuente struct {
	nombre_fuente  string
	formato_fuente FormatoFuente
	direccion      string
}
