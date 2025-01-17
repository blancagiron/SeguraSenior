# Elección de la biblioteca de aserciones:

---

## Criterios:


- **Instalación de dependencias:** Se evaluará si la biblioteca de aserciones requiere instalación adicional de paquetes o bibliotecas externas, priorizando aquellos que minimicen las dependencias externas. Go incentiva a la no dependencia de librerías externas, por lo que se dará prioridad a herramientas que formen parte de su [librería estándar](https://pkg.go.dev/std)

- **Estabilidad:** Debe tener una versión estable lanzada al menos hace un año. Esto asegura que la biblioteca está mantenida y que sigue recibiendo soporte.

- **Madurez:** Se considerará si la herramienta es madura y adoptada dentro de la comunidad de Go. Una herramienta madura ya ha alcanzado un nivel de estabilidad en el que los cambios y mejoras son menos frecuentes, ya que ha sido suficientemente probada y utilizada, por lo tanto, su funcionalidad es confiable y no necesita actualizaciones regulares.

- **Utilidad para comparar datos:** Se considerará que la biblioteca ofrezca funcionalidades específicas como métodos para la comparación de datos. Por ejemplo, métodos como `assert.Equal` permiten verificar la igualdad entre estructuras de datos, mientras que funciones como `assert.JSONEq` facilita la comparación de objetos JSON. La elección de una biblioteca con estas capacidades garantiza que podamos implementar pruebas adaptadas a las necesidades de mi proyecto.

- **Compatibilidad:** La biblioteca debe ser totalmente compatible con el test runner escogido, en este caso, el test runner estándar de Go `testing`.

---

## Candidatos:

- **Testify:** Es ampliamente adoptada en la comunidad de Go y compatible con el test runner estánadr `testing`. Requiere instalar una dependencia adicional. Su última versión salió hace 3 meses. Esta información se puede consultar en su [repositorio](https://github.com/stretchr/testify)

- **GoMega:** Es más común usarlo con `Ginkgo`. Requiere instalar `GoMega`como dependencia adicional lo que aumenta la carga de dependencias del proyecto. Su última actualización fue hace 2 semanas. Se puede consultar más información en su [repositorio](https://github.com/onsi/gomega)

- **GoCheck:**  Requiere una configuración adicional y añadir dependencias externas. Por lo que se puede consultar en su repositorio no se actualiza desde hace más de 5 años. Se puede consultar más información en este [repositorio](https://github.com/go-check/check)

- **Is:** Su repositorio es [este](https://github.com/matryer/is). No requiere dependencias adicionales y es compatible con el test runner estándar de Go sin necesidad de configuraciones adicionales. Su última actualización fue hace 2 años.

- **Ghost:** Su repositorio es [este](https://github.com/rliebz/ghost). Su última actualización fue hace 2 semanas. No requiere instalaciones adicionales ya que se integra de forma directa con el paquete estándar de Go `testing`. 

- **Caso especial. Go Testing:** No se trata de una biblioteca de aserciones pero permite la comprobación directa de condiciones y el manejo de errores. Se puede registrar un error usando `t.Error` personalizando el mensaje.  No requiere dependencias externas y es mantenida junto al lenguaje. Se puede consultar más información [aquí](https://pkg.go.dev/testing).


---

## Biblioteca de aserciones elegida:

La herramienta elegida finalmente es `Testify`, ya que aunque añada una dependencia externa al proyecto, proporciona métodos como `assert.Equal` o `assert.JSONEq`que permiten hacer comparaciones entre estructuras de datos, lo cual es particularmente relevante para mi proyecto. 






