package segurasenior

type Estadopoblacion int

const (
	Decreciente = iota
	Creciente
)

func buscatexto(direccion string)(string){
	//función para extraer texto del paquete net/http e io/ioutil
	return texto
}

func consultaFuentes(texto string, localidad string)(bool){
	return strings.Contains(texto,localidad)
}

func (e Estadopoblacion) String() (string){
	return [...]string{"Decreciente","Creciente"}
}

func (e Estadopoblacion) EnumIndex () (int){
	return int(e)
}

type Tendencia struct{
	Nombre string
	Estadotendencia Estadopoblacion
}
