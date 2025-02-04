# Elección de la biblioteca de aserciones:

---

## Criterios:


- **Evitar dependencias adicionales al lenguaje base:** La filosofía de Go incentiva a la no dependencia de librerías externas. Se evaluará si la biblioteca de aserciones necesita instalarse por separado al no formar parte de la biblioteca estándar de Go.

- **Última actualización y mantenimiento:** Se valorará que la herramienta tenga actualizaciones regulares publicadas en los últimos 12 meses. Si no hay actualizaciones regulares y recientes puede significar una falta de mantenimiento, lo que en un futuro aumentaría la deuda técnica.

---

## Candidatos:

### [**Testify**](https://github.com/stretchr/testify)

Su última versión fue lanzada hace 3 meses y se debe instalar al tratarse de una biblioteca externa.

### [**GoMega:**](https://github.com/onsi/gomega)

Se debe instalar la biblioteca por separado. Usualmente depende del framework Ginkgo, por lo que habría que instalar otra dependencia más (Ginkgo) en vez de sólo una.
  
---

### [**Is**](https://github.com/matryer/is) 

Es una biblioteca externa que hay que instalar. Además, su última actualización fue hace dos años, lo que podría indicar una falta de mantenimiento activo. 


---

### [**Ghost**](https://github.com/rliebz/ghost)

Es una biblioteca externa que es necesario instalar. Su última actualización fue hace 2 semanas, lo que refleja actividad reciente.

---

### [**Testing**](https://pkg.go.dev/testing)

Paquete que forma parte de la biblioteca estándar de Go que proporciona las herramientas básicas para escribir y ejecutar pruebas.

No tiene funciones de aserción predefinidas. Las comprobaciones se realizan con condicionales `if` y permite reportar fallos con métodos como `t.Errorf()`, `t.Fatal()` o `t.Fail()`, entre otros, los cuales indican si una prueba ha fallado. Esto permite validar el comportamiento del código y detectar errores.

 En [este enlace](https://pkg.go.dev/testing?tab=versions) se puede ver que se mantiene con actualizaciones regulares.

---

## Elección final:

Se elige Testing, ya que si ya existe una herramienta útil que proporcione lo necesario para escribir pruebas, Go incentiva a no instalar bibliotecas externas. Además al ser nativo, recibe actualizaciones regulares.




