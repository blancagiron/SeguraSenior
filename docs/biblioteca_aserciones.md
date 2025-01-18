# Elección de la biblioteca de aserciones:

---

## Criterios:


- **Instalación de dependencias:** Se evaluará si la biblioteca de aserciones requiere instalación adicional de paquetes o bibliotecas externas, priorizando aquellos que minimicen las dependencias externas. Go incentiva a la no dependencia de librerías externas, por lo que se dará prioridad a herramientas que formen parte de su [librería estándar](https://pkg.go.dev/std)

- **Madurez:** Se considerará si la herramienta es madura y adoptada dentro de la comunidad de Go. Una herramienta madura ya ha alcanzado un nivel de estabilidad en el que los cambios y mejoras son menos frecuentes, ya que ha sido suficientemente probada y utilizada, por lo tanto, su funcionalidad es confiable y no necesita actualizaciones regulares.

- **Utilidad para comparar datos:** Se evaluará si la biblioteca ofrece funcionalidades específicas para facilitar la comparación de datos. Esto incluye capacidades para verificar la igualdad entre estructuras complejas y la equivalencia entre representaciones de datos, como objetos JSON. Contar con estas herramientas permite implementar pruebas más claras y adaptadas a las necesidades del proyecto, reduciendo la complejidad en el desarrollo y mantenimiento de las pruebas.

- **Puntuación en Snyk Advisor:** Si la herramienta tiene una puntuación en [Snyk Advisor](https://snyk.io/advisor/golang) se tendrá en cuenta, ya que proporciona una evaluación objetiva de la salud del proyecto en función de cuatro aspectos clave: seguridad, popularidad, mantenimiento y comunidad.


---

## Candidatos:

### **Testify**

Su última versión fue lanzada hace 3 meses, y es una de las bibliotecas más utilizadas en el ecosistema de Go. Esta información se puede consultar en su [repositorio](https://github.com/stretchr/testify).

- **Instalación de dependencias:**  
  Requiere la instalación de una dependencia adicional, lo que aumenta la carga inicial del proyecto. Sin embargo, esta dependencia es ampliamente aceptada en la comunidad y fácil de integrar.  

- **Utilidad para comparar datos:**  
  Testify proporciona métodos específicos para comparar datos, como `assert.Equal` o `assert.JSONeq`, incluyendo la verificación de igualdad entre estructuras complejas, comparación de colecciones y equivalencia de representaciones en diferentes formatos como JSON. Estas funcionalidades simplifican la creación de pruebas detalladas y se deben tener en cuenta.

- **Puntuación en [Snyk Advisor](https://snyk.io/advisor/golang/github.com/stretchr/testify):**  
  Testify tiene una puntuación de **94**, reflejando una excelente salud del proyecto en términos de seguridad, popularidad, mantenimiento y comunidad.  
 
---

### **GoMega:** 

Su última actualización fue hace 2 semanas pero tiene menor adopción generalizada que Testify para pruebas unitarias. Esta información se puede consultar en su [repositorio](https://github.com/onsi/gomega).

- **Instalación de dependencias:**  
  Requiere la instalación de `GoMega` como dependencia adicional, lo que incrementa la carga de dependencias del proyecto. Es comúnmente usado en conjunto con `Ginkgo`, añadiendo una capa adicional de configuración.  
  
- **Utilidad para comparar datos:**  
  GoMega ofrece opciones  para comparar datos, incluyendo:  
  - `Ω(actual).Should(Equal(expected))`: Para verificar la igualdad entre valores o estructuras.  
  - `Ω(actual).Should(MatchJSON(expected))`: Para comparar JSON, validando la equivalencia estructural independientemente del orden de las claves.  
 
- **Puntuación en [Snyk Advisor](https://snyk.io/advisor/golang/github.com/onsi/gomega):**  
  GoMega tiene una puntuación de **96**. 

---

### **Is**  

Su última actualización fue hace dos años y es menos conocida y usada que otras alternativas como Testify. Esta información se puede consultar en su [repositorio](https://github.com/matryer/is)

- **Instalación de dependencias:**  
  No requiere dependencias adicionales y es completamente compatible con el test runner estándar de Go.
  
- **Utilidad para comparar datos:**  
  Is ofrece métodos como:  
  - `is.Equal(actual, expected)`: Verifica la igualdad entre valores.  

- **Puntuación en [Snyk Advisor](https://snyk.io/advisor/golang/github.com/matryer/is):**  
  Is tiene una puntuación de **59**, su mantenimiento es evaluado como inactivo.

---

### **Ghost**

 Su repositorio es [este](https://github.com/rliebz/ghost).  Su última actualización fue hace 2 semanas, lo que refleja actividad reciente. Es una herramienta menos conocida y con menor adopción dentro de la comunidad de Go.
 
- **Instalación de dependencias:**  
  No requiere dependencias adicionales y se integra directamente con el paquete estándar de Go `testing`.  

- **Utilidad para comparar datos:**  
 No ofrece métodos directos para comparar datos. En cambio, utiliza métodos generales de aserción como Must, MustNot, NoError, Should y ShouldNot, que evalúan resultados de manera booleana. 

- **Puntuación en [Snyk Advisor](https://snyk.io/advisor/golang/github.com/rliebz/ghost):**  
  Ghost tiene una puntuación de **69**, con buenos indicadores de mantenimiento y seguridad, aunque califica su comunidad y popularidad como limitada.

---

### **Go Testing (caso especial)**  

Aunque Go Testing no es una biblioteca de aserciones, se incluye en la evaluación porque es parte de la librería estándar de Go y ofrece funcionalidades básicas para comprobar condiciones y manejar errores. Permite registrar errores y personalizar mensajes, lo que lo convierte en una opción válida en proyectos que prefieren no agregar dependencias externas. Aunque no proporciona métodos avanzados para comparar datos complejos, puede ser suficiente para pruebas más simples. 

Se puede consultar más información en [este enlace](https://pkg.go.dev/testing)

- **Instalación de dependencias:**  
  No requiere dependencias adicionales, ya que forma parte de la librería estándar de Go.  

- **Utilidad para comparar datos:**  
  Go Testing no incluye métodos predefinidos para comparar datos. La verificación de condiciones debe realizarse manualmente mediante `t.Error` o `t.Errorf`, lo que puede ser más incómodo y propenso a errores en comparación con las bibliotecas mencionadas. Para casos complejos, es necesario escribir lógica adicional.  

- **Puntuación en Snyk Advisor:**  
  Go Testing no está listado en Snyk Advisor, pero su fiabilidad está garantizada como parte del lenguaje.  

---

## Biblioteca de aserciones elegida:

Finalmente se elige Testify. Aunque requiere una dependencia adicional, es ampliamente utilizada y bien mantenida, con actualizaciones reciente. Testify también ofrece métodos avanzados para comparar datos complejos, como assert.Equal y assert.JSONEq, lo que facilita las pruebas. Además, tiene una excelente puntuación en Snyk Advisor, asegurando su fiabilidad en términos de seguridad, popularidad y mantenimiento, superando a otras bibliotecas con puntuaciones más bajas o inexistentes.



