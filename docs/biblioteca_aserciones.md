# Elección de la biblioteca de aserciones:

---

## Criterios:

- **Compatibilidad con el test runner escogido:** Debe ser compatible con el test runner escogido, en este caso `Testify`, sin necesidad de configuraciones adicionales, garantizando que las aserciones funcionen sin inconvenientes al usarlas junto al test runner.

- **Instalación de dependencias:** Minimizar el número de dependencias del proyecto. Si el test runner elegido proporciona las aserciones necesarias y cubre los criterios mencionados, no se deberían añadir más bibliotecas externas.

---

## Candidatos:

- **Testify Assertions:** Al ser parte del paquete `Testify` está completamente integrado con su test runner y no requiere configuraciones adicionales, además de que tampoco añade dependencias extra.  Esta información se puede consultar en su [repositorio](https://github.com/stretchr/testify)

- **Gomega:** No es parte de `Testify`pero se puede usar junto a este test runner, es más común usarlo con `Ginkgo`. Requiere instalar `Gomega`como dependencia adicional lo que aumenta la carga de dependencias del proyecto. Se puede consultar más información en su [repositorio](https://github.com/onsi/gomega)

- **GoCheck:**  No tiene integración directa con `Testify`por lo que se requiere una configuración adicional y añadir dependencias externas. Se puede consultar más información en este [repositorio](https://github.com/go-check/check)

- **Is:** Su repositorio es [este](https://github.com/matryer/is). No requiere dependencias adicionales y es compatible con el test runner estándar de Go y Testify sin necesidad de configuraciones adicionales. 
---

## Biblioteca de aserciones elegida:

Cómo estamos utilizando `Testify` como test runner, se elige `Testify Assertions`como biblioteca de aserciones ya que ofrece compatibilidad completa con el test runners y no introduce dependencias adicionales innecesarias.










---






