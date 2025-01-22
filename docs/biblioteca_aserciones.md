# Elección de la biblioteca de aserciones:

---

## Criterios:


- **Evitar dependencias adicionales:** Se evaluará si la biblioteca de aserciones requiere instalación adicional de paquetes, priorizando aquellas que minimicen las dependencias externas. 

- **Última actualización y mantenimiento:** Se valorará que la herramienta tenga actualizaciones regulares publicadas en los últimos 12 meses. Si no hay actualizaciones regulares y recientes puede significar una falta de mantenimiento, lo que en un futuro aumentaría la deuda técnica.

- **Soporte para comparación de datos:** Debe permitir realizar comparaciones directas entre estructuras de datos más complejos como mapas, slices o structs sin requerir construcciones adicionales (como sentencias condicionales o bloques personalizados), lo que sería igual a tener que añadir más código.

---

## Candidatos:

### [**Testify**](https://github.com/stretchr/testify)

Su última versión fue lanzada hace 3 meses. Esta información se puede consultar en su [repositorio]
Requiere la instalación de una dependencia adicional. Permite realizar comparaciones sin necesidad de sentencias condicionales adicionales con funciones como `assert.Equal`.
 
---

### [**GoMega:**](https://github.com/onsi/gomega)


Requiere instalar la biblioteca como dependencia adicional y depende usualmente del framework Ginkgo, por lo que habría que instalar otra dependencia. Su última actualización fue hace 2 semanas y proporciona funciones para comparación directa como `Expect().To(Equal(...))`. 

  
---

### [**Is**](https://github.com/matryer/is) 

No requiere dependencias adicionales Go. Su última actualización fue hace dos años. Proporciona métodos para comparaciones de datos como `is.Equal`. 


---

### [**Ghost**](https://github.com/rliebz/ghost)

No requiere dependencias adicionales. Su última actualización fue hace 2 semanas, lo que refleja actividad reciente. Ghost no ofrece métodos directos para comparar datos. En cambio, utiliza métodos generales que evalúan resultados de manera booleana. 


---

## Biblioteca de aserciones elegida:

Finalmente se elige Testify. Aunque requiere una dependencia adicional, está bien mantenida con actualizaciones recientes y permite realizar comparaciones de datos de manera directa sin necesitar construcciones adicionales.



