# Elección de la biblioteca de aserciones:

---

## Criterios:


- **Instalación de dependencias:** Se evaluará si la biblioteca de aserciones requiere instalación adicional de paquetes, priorizando aquellas que minimicen las dependencias externas. 

- **Última actualización y mantenimiento:** Se valorará que la herramienta tenga actualizaciones regulares publicadas en los últimos 12 meses. Si no hay actualizaciones regulares y recientes puede significar una falta de mantenimiento, lo que en un futuro aumentaría la deuda técnica.

---

## Candidatos:

### **Testify**

Su última versión fue lanzada hace 3 meses. Esta información se puede consultar en su [repositorio](https://github.com/stretchr/testify).

Requiere la instalación de una dependencia adicional, el paquete `Testify`, lo que aumenta la carga inicial del proyecto. 
 
---

### **GoMega:** 


Requiere instalar la biblioteca como dependencia adicional. Su última actualización fue hace 2 semanas. Esta información se puede consultar en su [repositorio](https://github.com/onsi/gomega).

  
---

### **Is**  

No requiere dependencias adicionales Go. Su última actualización fue hace dos años. Esta información se puede consultar en su [repositorio](https://github.com/matryer/is)


---

### **Ghost**

No requiere dependencias adicionales y se integra directamente con el paquete estándar de Go `testing`. Su repositorio es [este](https://github.com/rliebz/ghost). Su última actualización fue hace 2 semanas, lo que refleja actividad reciente. Ghost no ofrece métodos directos para comparar datos. En cambio, utiliza métodos generales de aserción como Must, MustNot, NoError, Should y ShouldNot, que evalúan resultados de manera booleana. 

---

### **Testing (caso especial)**  

No requiere dependencias adicionales, ya que forma parte de la librería estándar de Go. Aunque Testing no es una biblioteca de aserciones, se incluye en la evaluación porque es parte de la librería estándar de Go y ofrece funcionalidades básicas para comprobar condiciones y manejar errores. Permite registrar errores y personalizar mensajes, lo que lo convierte en una opción válida en proyectos que prefieren no agregar dependencias externas. Aunque no proporciona métodos avanzados para comparar datos complejos, puede ser suficiente para pruebas más simples. 

Testing no incluye métodos predefinidos para comparar datos. La verificación de condiciones debe realizarse manualmente mediante `t.Error` o `t.Errorf`, lo que puede ser más incómodo y propenso a errores en comparación con las bibliotecas mencionadas. Para casos complejos, es necesario escribir lógica adicional.  
 
Se puede consultar más información en [este enlace](https://pkg.go.dev/testing)

---

## Biblioteca de aserciones elegida:

Finalmente se elige Testify. Aunque requiere una dependencia adicional, es ampliamente utilizada y bien mantenida, con actualizaciones recientes. 



