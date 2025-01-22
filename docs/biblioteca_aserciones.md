# Elección de la biblioteca de aserciones:

---

## Criterios:


- **Evitar dependencias adicionales al lenguaje base:** La filosofía de Go incentiva a la no dependencia de librerías externas. Se evaluará si la biblioteca de aserciones necesita instalarse por separado al no formar parte de la biblioteca estándar de Go.

- **Última actualización y mantenimiento:** Se valorará que la herramienta tenga actualizaciones regulares publicadas en los últimos 12 meses. Si no hay actualizaciones regulares y recientes puede significar una falta de mantenimiento, lo que en un futuro aumentaría la deuda técnica.

- **Comparación de datos sin código adicional:** La biblioteca debe permitir comparar estructuras complejas como mapas, slices o structs sin necesidad de añadir más código (como estructuras condicionales) para validaciones. Más código añadido es más código que hay que revisar.

---

## Candidatos:

### [**Testify**](https://github.com/stretchr/testify)

Su última versión fue lanzada hace 3 meses. 
Se debe instalar, es una dependencia adicional que no forma parte de la biblioteca estándar de Go . Permite realizar comparaciones sin necesidad de sentencias condicionales adicionales con funciones como `assert.Equal`.
 
---

### [**GoMega:**](https://github.com/onsi/gomega)


Se debe instalar la biblioteca por separado. Usualmente depende del framework Ginkgo, por lo que habría que instalar otra dependencia más en vez de sólo una. Su última actualización fue hace 2 semanas y proporciona funciones para comparación directa como `Expect().To(Equal(...))`. 

  
---

### [**Is**](https://github.com/matryer/is) 

No requiere instalar dependencias externas adicionales a la biblioteca estándar de Go. Sin embargo, su última actualización fue hace dos años, lo que podría indicar una falta de mantenimiento activo. Ofrece métodos básicos para realizar comparaciones, como is.Equal.


---

### [**Ghost**](https://github.com/rliebz/ghost)

No requiere instalar dependencias externas adicionales a la biblioteca estándar de Go. Su última actualización fue hace 2 semanas, lo que refleja actividad reciente. Sin embargo, Ghost no ofrece métodos específicos para comparar estructuras de datos.


---

## Biblioteca de aserciones elegida:

Se elige Testify. Es la opción elegida ya que Is no cumple el criterio de actualizaciones regulares, Ghost no proporciona métodos directos para comparar datos y GoMega tiene una dependencia extra en el framework Ginkgo para su uso completo. Aunque no forma parte de la biblioteca estándar de Go y se tiene que instalar por separado, Testify es la opción que cumple un mayor número de criterios.




