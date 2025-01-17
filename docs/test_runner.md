# Elección del test runner:

---

## Criterios:

Los criterios seguidos para escoger el test runner más adecuado son:

- **Instalación de dependencias:** Se evaluará si el test runner requiere instalación adicional de paquetes o bibliotecas externas, priorizando aquellos que minimicen las dependencias externas. Go incentiva a la no dependencia de librerías externas, por lo que se dará prioridad a herramientas que formen parte de su [librería estándar](https://pkg.go.dev/std)

- **Última actualización:** Se valorará que haya actualizaciones regulares publicadas en los últimos 12 meses y que el intervalo entre ellas no sea superior a 3 meses, ya que un intervalo mayor puede indicar falta de mantenimiento, lo que aumentaría la deuda técnica.

- **Madurez y estabilidad:** Se considerará si la herramienta es madura y adoptada dentro de la comunidad de Go. En el caso de herramientas que ya sean maduras, estables e integradas no es necesario que tengan actualizaciones regulares. Una herramienta madura ya ha alcanzado un nivel de estabilidad en el que los cambios y mejoras son menos frecuentes, ya que ha sido suficientemente probada y utilizada, por lo tanto, su funcionalidad es confiable.

- **Cobertura de código:** Verificaremos si el test runner facilita la generación de métricas de cobertura.

---

## Candidatos:

- **Go Testing:** 
	- `Go Testing`forma parte de la biblioteca estándar de Go, por lo que no requiere ninguna instalación adicional y es mantenida junto al lenguaje. Su última actualización fue hace un mes y es compatible con herramientas estándar de Go para cobertura, como el comando `go test -cover`, el cuál genera métricas de cobertura.

	Podemos consultar más información [aquí](https://pkg.go.dev/testing) y en su [repositorio](https://github.com/golang/go/blob/master/src/testing/testing.go)
	
 
- **Ginkgo**
	- Requiere la instalación de varias dependencias, como `github.com/onsi/ginkgo` y `github.com/onsi/gomega` para su funcionalidad completa, lo que puede hacer que el proyecto dependa de más bibliotecas. Es compatible con herramientas de cobertura estándar de Go, como `go test -cover`. 
	  
	Podemos consultar más información en su [repositorio](https://github.com/onsi/ginkgo), como que su última actualización fue hace menos de un mes.


- **GoConvey**
	
	Requiere la instalación de la librería `GoConvey` (`go get github.com/smartystreets/goconvey`), lo que aumenta las dependencias externas en el proyecto. Además, su última actualización fue hace casi un año. Al igual que los otros candidatos, es compatible con herramientas de cobertura estándar de Go, como `go test -cover`. 
	  
	Podemos consultar más información en su [repositorio](https://github.com/smartystreets/goconvey)
	
---

## Test runner elegido:

El test runner seleccionado es `Go Testing` ya que el propio lenguaje incentiva el uso de la biblioteca estándar para reducir las dependencias de bibliotecas externas, además de que se trata de una herramienta madura. 


