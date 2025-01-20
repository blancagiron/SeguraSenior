# Elección del test runner:

---

## Criterios:

Los criterios seguidos para escoger el test runner más adecuado son:

- **Instalación de dependencias:** Se evaluará si el test runner requiere instalación adicional de paquetes o bibliotecas externas, priorizando aquellos que minimicen las dependencias externas. Go incentiva a la no dependencia de librerías externas, por lo que se dará prioridad a herramientas que formen parte de su [librería estándar](https://pkg.go.dev/std)

- **Última actualización:** Se valorará que la herramienta tenga actualizaciones regulares publicadas en los últimos 12 meses, preferiblemente con intervalos no superiores a 3 meses. Dado que el ciclo de lanzamiento de Go es semestral [(Go Release Cycle)](https://go.dev/wiki/Go-Release-Cycle#:~:text=Go%20is%20released%20every%20six,polishing%20called%20the%20release%20freeze.), actualizaciones más frecuentes permiten a las herramientas mantenerse en sintonía con los cambios en el lenguaje. Un intervalo mayor podría sugerir una menor atención al mantenimiento, lo que incrementa el riesgo de acumulación de deuda técnica y posibles incompatibilidades futuras. Un ciclo de 3 meses garantiza agilidad en la respuesta a nuevas versiones y un soporte constante para los desarrolladores que dependen de la herramienta.

- **Madurez y estabilidad:** Se considerará si la herramienta es madura y adoptada dentro de la comunidad de Go. En el caso de herramientas que ya sean maduras, estables e integradas no es necesario que tengan actualizaciones regulares. Una herramienta madura ya ha alcanzado un nivel de estabilidad en el que los cambios y mejoras son menos frecuentes, ya que ha sido suficientemente probada y utilizada, por lo tanto, su funcionalidad es confiable.



---

## Candidatos:

- **Testing:** 
	- `Testing`forma parte de la biblioteca estándar de Go, por lo que no requiere ninguna instalación adicional y es mantenida junto al lenguaje. Su última actualización fue hace un mes.

	Podemos consultar más información [aquí](https://pkg.go.dev/testing) y en su [repositorio](https://github.com/golang/go/blob/master/src/testing/testing.go)
	
 
- **Ginkgo**
	- Requiere la instalación de varias dependencias, como `github.com/onsi/ginkgo` y `github.com/onsi/gomega` para su funcionalidad completa, lo que puede hacer que el proyecto dependa de más bibliotecas. 
	  
	Podemos consultar más información en su [repositorio](https://github.com/onsi/ginkgo), como que su última actualización fue hace menos de un mes.


- **GoConvey**
	
	Requiere la instalación de la librería `GoConvey` (`go get github.com/smartystreets/goconvey`), lo que aumenta las dependencias externas en el proyecto. Además, su última actualización fue hace casi un año. 
	  
	Podemos consultar más información en su [repositorio](https://github.com/smartystreets/goconvey)
	
---

## Test runner elegido:

El test runner seleccionado es `Testing` ya que el propio lenguaje incentiva el uso de la biblioteca estándar para reducir las dependencias de bibliotecas externas, además de que se trata de una herramienta madura. 


