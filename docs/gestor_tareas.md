# Gestor de tareas:

---

## Criterios:

Los criterios que se han seguido para escoger el gestor de tareas más adecuado son:

- Frecuencia de versiones: debe estar actualizado y mantenido con regularidad, para minimizar la deuda técnica. Aun así, en casos en los que la herramienta ya sea estable y no requiera actualizaciones constantes porque ya cumple su propósito y tiene soporte activo por la comunidad y disponibilidad de documentación, también se puede elegir.

- Facilidad de uso y familiaridad con el conocimiento. Se escogerá un gestor con el que se esté familiarizado con sus reglas y la sintaxis usada frente a uno del que se conozca menos.


---

## Candidatos:

- **Mage:** Si bien su comunidad es grande, las actualizaciones de este gestor cuentan con largos períodos de inactividad en su desarrollo, esto causa dudas sobre si podría incrementar la deuda técnica en proyectos que dependan de la herramienta. Se puede consultar en su [Repositorio oficial](https://github.com/magefile/mage)

- **Gotaskr:** Esta diseñado para facilitar la ejecución de tareas automatizadas desde la línea de comandos, permite definir tareas usando código Go de forma directa lo que hace el proceso más accesible y sencillo para los desarrolladores familiarizados con el lenguaje. Tiene una extensión en VSCode, pero su comunidad es muy pequeña y su última actualización fue hace 2 meses. Esto se puede consultar en su [Repositorio oficial](https://github.com/Roemer/gotaskr)

- **Make:** Herramienta muy conocida y consolidada. No recibe actualizaciones desde hace un tiempo pero ya está establecida y cuenta con una gran comunidad.


- **Task:**  Su comunidad es grande y recibe actualizaciones constantes, se utiliza un archivo `Taskfile.yml` en el que se definen las tareas. Se puede ver en el [Repositorio oficial](https://github.com/go-task/task)

- **Gilbert:** Gilbert es un sistema de construcción y gestor de tareas para proyectos en Go. Permite definir y ejecutar tareas de forma declarativa utilizando un archivo gilbert.yaml. No se actualiza desde 2019 y su comunidad no es grande, esto se pueden consultar en el [Repositorio oficial](https://github.com/go-gilbert/gilbert) 

- **Just:** Ofrece una manera fácil y flexible de automatizar tareas. Tiene una gran comunidad y actualizaciones constantes. [Repositorio oficial](https://github.com/casey/just)


Para más información sobre estos gestores, se pueden consultar las siguientes fuentes además de sus repositorios:


- [Reddit](https://www.reddit.com/r/golang/comments/10s1gja/task_runner_like_gotasktask_but_in_pure_go_no/?tl=es-es&rdt=58809)
- [AwesomeGo](https://awesome-go.com/)

---

## Gestor elegido:

Después de haber descartado Mage, Gilbert y Gotaskr por no cumplir con ninguno de los criterios, las opciones restantes son Make (ya que si bien no recibe actualizaciones constantes es una herramienta establecida), Just y Task. Make queda descartado en esta ocasión porque si bien tengo experiencia con este gestor, las opciones de Just y Task me parecen menos complejas de usar. Finalmente, entre Just y Task, escojo como gestor de tareas Task ya que está dirigido a proyectos escritos en Go.
