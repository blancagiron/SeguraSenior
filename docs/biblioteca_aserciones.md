# Elección de la biblioteca de aserciones:

---

## Criterios:


- **Instalación de dependencias:** Se evaluará si la biblioteca de aserciones requiere instalación adicional de paquetes o bibliotecas externas, priorizando aquellos que minimicen las dependencias externas. Go incentiva a la no dependencia de librerías externas, por lo que se dará prioridad a herramientas que formen parte de su [librería estándar](https://pkg.go.dev/std)

- **Madurez:** Se considerará si la herramienta es madura y adoptada dentro de la comunidad de Go. Una herramienta madura ya ha alcanzado un nivel de estabilidad en el que los cambios y mejoras son menos frecuentes, ya que ha sido suficientemente probada y utilizada, por lo tanto, su funcionalidad es confiable y no necesita actualizaciones regulares.

---

## Candidatos:

### **Testify**

Su última versión fue lanzada hace 3 meses, y es una de las bibliotecas más utilizadas en el ecosistema de Go. Esta información se puede consultar en su [repositorio](https://github.com/stretchr/testify).

- **Instalación de dependencias:**  
  Requiere la instalación de una dependencia adicional, el paquete `Testify`, lo que aumenta la carga inicial del proyecto. Sin embargo, esta dependencia es ampliamente aceptada en la comunidad y fácil de integrar.  
 
---

### **GoMega:** 

Su última actualización fue hace 2 semanas pero tiene menor adopción generalizada que Testify para pruebas unitarias. Esta información se puede consultar en su [repositorio](https://github.com/onsi/gomega).

- **Instalación de dependencias:**  
  Requiere la instalación de `GoMega` como dependencia adicional, lo que incrementa la carga de dependencias del proyecto. Es comúnmente usado en conjunto con `Ginkgo`, añadiendo una capa adicional de configuración.  
   
---

### **Is**  

Su última actualización fue hace dos años y es menos conocida y usada que otras alternativas como Testify. Esta información se puede consultar en su [repositorio](https://github.com/matryer/is)

- **Instalación de dependencias:**  
  No requiere dependencias adicionales y es completamente compatible con el test runner estándar de Go.
  
- **Utilidad para comparar datos:**  
  Is ofrece métodos como:  
  - `is.Equal(actual, expected)`: Verifica la igualdad entre valores.  


---

### **Ghost**

 Su repositorio es [este](https://github.com/rliebz/ghost).  Su última actualización fue hace 2 semanas, lo que refleja actividad reciente. Es una herramienta menos conocida y con menor adopción dentro de la comunidad de Go.
 
- **Instalación de dependencias:**  
  No requiere dependencias adicionales y se integra directamente con el paquete estándar de Go `testing`.  

- **Utilidad para comparar datos:**  
 No ofrece métodos directos para comparar datos. En cambio, utiliza métodos generales de aserción como Must, MustNot, NoError, Should y ShouldNot, que evalúan resultados de manera booleana. 

---

### **Testing (caso especial)**  

Aunque Testing no es una biblioteca de aserciones, se incluye en la evaluación porque es parte de la librería estándar de Go y ofrece funcionalidades básicas para comprobar condiciones y manejar errores. Permite registrar errores y personalizar mensajes, lo que lo convierte en una opción válida en proyectos que prefieren no agregar dependencias externas. Aunque no proporciona métodos avanzados para comparar datos complejos, puede ser suficiente para pruebas más simples. 

Se puede consultar más información en [este enlace](https://pkg.go.dev/testing)

- **Instalación de dependencias:**  
  No requiere dependencias adicionales, ya que forma parte de la librería estándar de Go.  

- **Utilidad para comparar datos:**  
   Testing no incluye métodos predefinidos para comparar datos. La verificación de condiciones debe realizarse manualmente mediante `t.Error` o `t.Errorf`, lo que puede ser más incómodo y propenso a errores en comparación con las bibliotecas mencionadas. Para casos complejos, es necesario escribir lógica adicional.  
 

---

## Biblioteca de aserciones elegida:

Finalmente se elige Testify. Aunque requiere una dependencia adicional, es ampliamente utilizada y bien mantenida, con actualizaciones reciente. Testify también ofrece métodos avanzados para comparar datos complejos, como assert.Equal y assert.JSONEq, lo que facilita las pruebas. 



