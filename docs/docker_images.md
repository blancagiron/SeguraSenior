# Elección de la imagen de docker:

---

## Criterios:

- **Tamaño de la imagen:** Se valorará minimizar el tamaño del contenedor sin sacrificar funcionalidad, facilitando despliegues más rápidos y reduciendo el consumo de recursos.

- **Imágenes Oficiales y Verificadas:**  Las imágenes oficiales y de usuarios verificados ofrecen mayor confianza en términos de seguridad y mantenimiento. En Docker Hub esto se puede consultar viendo cuáles tienen las insignias de `Docker Official Image` o de `Verified Publisher`.

- **Soporte y actualizaciones regulares:**  Se valorará que las imágenes reciban actualizaciones frecuentes, ya que esto es importante para corregir posibles vulnerabilidades y mejorar funcionalidades.

- **Compatibilidad con Golang:** Se considerará que la imagen elegida venga con Go ya instalado, para ahorrarnos este paso y evitar posibles problemas con las versiones en el futuro.

---

## Candidatos:

### [Golang](https://hub.docker.com/_/golang):

Son las imágenes de Docker de Go, mantenidas de forma regular por la comunidad de Go. Todas tienen la insignia de `Docker Official Image` y Go preinstalado. Existen distintas variantes:

- **Basadas en Debian:**
 - `golang:<latest>`: Versión estándar basada en Debian. Su útima actualización fue hace una semana. Su tamaño es en torno a 300 MB.
 - `golang:<version>-bookworm`: Versión basada en Debian Bookworm, la cúal es la versión estable más reciente (Debian 12). Su última actualización fue hace dos semanas. Su tamaño es en torno a 284 MB.

- **Basadas en Alpine:**
  - `golang:<version>-alpine`: Imagen basada en Alpine Linux. Reduce el tamaño del contenedor. Su última actualización fue hace dos semanas. Su tamaño es en torno a 67.82 MB

### [Alpine](https://hub.docker.com/_/alpine)

Cuenta con la insignia de `Docker Official Image`. No tiene Go preinstalado. Es útil para construir contenedores ligeros.

- `alpine:<latest>`: Su última actualización fue hace 20 días. Su tamaño es en torno a 3-4 MB, pero hay que tener en cuenta que tendríamos que instalar Go.


### [Bitnami/Golang](https://hub.docker.com/r/bitnami/golang):  
- `[bitnami/golang:<version>](https://hub.docker.com/r/bitnami/golang)`: Cuenta con la insignia de `Verified Publisher`, su última actualización fue hace días. Tiene Go preinstalado y su tamaño suele ir en torno a los 230 MB.



