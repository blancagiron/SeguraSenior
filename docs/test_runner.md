# Elección del test runner:

---

## Criterios:

Los criterios a seguir para escoger el test runner más adecuado son:

- **Evitar dependencias adicionales:** Se evaluará si el test runner necesita la instalación de paquetes o bibliotecas externas. Se dará prioridad a herramientas que reduzcan al mínimo las dependencias adicionales, alineándose con la filosofía de Go de evitar el uso de librerías externas siempre que sea posible. Por ello, se preferirán herramientas incluidas en su  [librería estándar](https://pkg.go.dev/std). 

- **Última actualización:** Se valorará que la herramienta tenga actualizaciones regulares publicadas en los últimos 12 meses, preferiblemente con intervalos no superiores a 3 meses. Dado que el ciclo de lanzamiento de Go es semestral [(Go Release Cycle)](https://go.dev/wiki/Go-Release-Cycle#:~:text=Go%20is%20released%20every%20six,polishing%20called%20the%20release%20freeze.), actualizaciones más frecuentes permiten a las herramientas mantenerse en sintonía con los cambios en el lenguaje. Un intervalo mayor podría sugerir una menor atención al mantenimiento, lo que incrementa el riesgo de acumulación de deuda técnica y posibles incompatibilidades futuras. Un ciclo de 3 meses garantiza agilidad en la respuesta a nuevas versiones y un soporte constante para los desarrolladores que dependen de la herramienta.

---

## Candidatos:

- **Testing:** 
	- `Testing`forma parte de la biblioteca estándar de Go, por lo que no requiere ninguna instalación adicional y es mantenida junto al lenguaje. Su última actualización fue hace un mes.

	Podemos consultar más información [aquí](https://pkg.go.dev/testing) y en su [repositorio](https://github.com/golang/go/blob/master/src/testing/testing.go)
	
 
- **Ginkgo**
	- No minimiza las dependencias externas. Requiere su instalación y la de `github.com/onsi/gomega` para su funcionalidad completa.
	  
	Podemos consultar más información en su [repositorio](https://github.com/onsi/ginkgo), como que su última actualización fue hace menos de un mes.


- **GoConvey**
	
	Requiere la instalación de dependencias adicionales. Además, su última actualización fue hace casi un año. 
	  
	Podemos consultar más información en su [repositorio](https://github.com/smartystreets/goconvey)
	
---

## Test runner elegido:

El test runner seleccionado es `Testing`, forma parte de la biblioteca estándar de Go.

--- 

# Herramienta CLI:

---

# Criterios:

Los criterios a seguir para escoger la herramienta CLI más adecuada son:

- **Evitar dependencias adicionales:** Se evaluará si la herramienta necesita la instalación de paquetes o bibliotecas externas.
  
- **Soporte y mantenimiento:**  Se valorará que la herramienta tenga actualizaciones regulares publicadas en los últimos 12 meses.

---

## Candidatos:


- **[go test](https://pkg.go.dev/testing)**: forma parte de la biblioteca estándar de Go, al ser una herramienta nativa, recibe actualizaciones constantes junto con el lenguaje Go.
  
- **[gotestsum](https://github.com/gotestyourself/gotestsum):** Al ser una utilidad externa necesita instalación adicional. Su última actualización fue hace 3 meses.
  
- **[richgo](https://github.com/kyoh86/richgo):** Es una utilidad externa, necesita instalación adicional. Su última actualización fue en 2023.

## Herramienta CLI elegida:

La herramienta CLI elegida es `go test` ya que  es es suficiente para manejar los distintos escenarios de pruebas unitarias y forma parte del estándar de Go, por lo que no requiere instalación adicional.