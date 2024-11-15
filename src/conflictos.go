package segurasenior

import "fmt"

func ContrastarFuentes(reporte *Reportaje) error {
	for i := 0; i < len(reporte.Fuente); i++ {
		fuenteog := reporte.Fuente[i]

		for k := 0; k < len(reporte.Fuente); k++ {
			otrafuente := reporte.Fuente[k]

			if fuenteog.Poblacion.Nombre == otrafuente.Poblacion.Nombre && fuenteog.Poblacion.Estadotendencia == otrafuente.Poblacion.Estadotendencia {
				continue
			}

			if fuenteog.Poblacion.ConflictoInfo(&otrafuente.Poblacion) {
				return fmt.Errorf("informacion conflictiva entre fuentes: %s en %s ",
					fuenteog.Direccion, otrafuente.Direccion)
			}
		}
	}
	return nil
}
