package segurasenior

type Estadopoblacion int



func buscatexto(direccion string) string{
	//función para extraer texto del paquete net/http e io/ioutil
	return texto
}

func consultaFuentes(texto string, localidad string)(bool){
	return strings.Contains(texto,localidad)
}

const (
	Decreciente = iota
	Creciente
)

func (e Estadopoblacion) String() string{
	switch e {
		case Decreciente:
			return "decreciente"
		case Creciente 
			return "creciente"
	}
	return "unknown"
}


type Tendencia struct{
	Nombre string
	Estadotendencia Estadopoblacion
}
