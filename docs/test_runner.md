# Elección del test runner:

---

## Criterios:

Los criterios que se han seguido para escoger el test runner más adecuado son:

- **Instalación de dependencias:** Se evaluará si el test runner requiere instalación adicional de paquetes o bibliotecas externas, priorizando aquellos que minimicen las dependencias externas.

- **Cobertura de código:** Verificaremos si el test runner facilita la generación de métricas de cobertura.

- **Soporte para mocks:** Evaluaremos si el test runner incluye soporte para la simulación de servicios externos mediante mocks, lo cuál es importante para facilitar la simulación de servicios externos como solicitudes HTTP.

---

## Candidatos:

- **Go Testing:** 
	- **Instalación de dependencias:**  
	  - `Go Testing` forma parte de la biblioteca estándar de Go, por lo que no se necesita ninguna instalación adicional.
	  
	- **Cobertura de código:**  
	  - Compatible con herramientas estándar de Go para cobertura, como el comando `go test -cover`, que genera métricas de cobertura.
	  
	- **Soporte para mocks:**  
	  - No incluye soporte nativo para mocks. Sin embargo, se puede usar con bibliotecas externas como `Testify/mock` o `gomock`.
	 
	Podemos consultar más información [aquí](https://pkg.go.dev/testing)
	
	
- **Testify**
	- **Instalación de dependencias:**  
	  - Hay que instalar el paquete `testify` (`go get github.com/stretchr/testify`), lo cual es una dependencia adicional.
	  
	- **Cobertura de código:**  
	  - No genera métricas de cobertura directamente, pero es compatible con herramientas de cobertura de Go, como `go test -cover`.
	  
	- **Soporte para mocks:**  
	  - Incluye soporte nativo para mocks a través del paquete `mock`.

	Podemos consultar más información en su [repositorio](https://github.com/stretchr/testify)
	
 
- **Ginkgo**
	- **Instalación de dependencias:**  
	  - Requiere la instalación de varias dependencias, como `github.com/onsi/ginkgo` y `github.com/onsi/gomega` para su funcionalidad completa, lo que puede hacer que el proyecto dependa de más bibliotecas.
	  
	- **Cobertura de código:**  
	  - Compatible con herramientas de cobertura estándar de Go, como `go test -cover`. 
	  
	- **Soporte para mocks:**   
	  - Requiere configuración adicional para mocks. Se puede usar con bibliotecas externas como `Testify/mock` o `gomock`
	  
	Podemos consultar más información en su [repositorio](https://github.com/onsi/ginkgo)
  
- **Gomega**
	- **Instalación de dependencias:**  
	  - Se instala como dependencia de Ginkgo, pero también puede usarse de forma independiente (`go get github.com/onsi/gomega`).
	  
	- **Cobertura de código:**  
	  - Compatible con herramientas de cobertura estándar de Go, como `go test -cover`. 
	  
	- **Soporte para mocks:**  
	  - No proporciona soporte directo para mocks, pero se puede combinar con herramientas como `gomock` o `Testify/mock` para realizar pruebas de simulación de dependencias externas.
	  
	  Podemos consultar más información en su [repositorio](https://github.com/onsi/gomega)
	  


- **GoConvey**
	- **Instalación de dependencias:**  
	  - Requiere la instalación de la librería `GoConvey` (`go get github.com/smartystreets/goconvey`), lo que aumenta las dependencias externas en el proyecto.
	  
	- **Cobertura de código:**  
	  - Compatible con herramientas de cobertura estándar de Go, como `go test -cover`. 
	  
	- **Soporte para mocks:**  
	  - No proporciona soporte directo para mocks, pero se puede combinar con herramientas como `gomock` o `Testify/mock` para realizar pruebas de simulación de dependencias externas.

	Podemos consultar más información en su [repositorio](https://github.com/smartystreets/goconvey)


 

