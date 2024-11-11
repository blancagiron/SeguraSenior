package segurasenior

import(
	"strings"
)

func extraerinfo(texto string, localidad string)(bool){
	return strings.Contains(texto,localidad)
}

type Reportaje struct{
	nombre string
	poblacion int
	retransmision string
}
	
