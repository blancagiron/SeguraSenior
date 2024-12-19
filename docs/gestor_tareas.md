# Gestor de tareas:

---

## Criterios:

Los criterios que se han seguido para escoger el gestor de tareas más adecuado son:

- Última actualización: 

Debe haber actualizaciones regulares publicadas en los últimos 12 meses para considerar que la herramienta está activa y mantenida. El intervalo entre actualizaciones no debe ser superior a 3 meses ya que un intervalo mayor podría indicar falta de mantenimiento, lo que aumentaría la deuda técnica. Esto no es necesario para herramientas que ya sean maduras y estables, como es el caso de Make

- Comunidad : 

Debe tener una comunidad amplia y activa ya que una herramienta utilizada por un número significativo de usuarios suele ser un indicativo de su utilidad y estabilidad. Esto lo podemos medir mediante en número de estrellas y forks en las plataformas de su repositorio, ya que indican el interés general y cuantos usuarios están contribuyendo o adaptando la herramienta. 

Debe haber discusiones relacionadas con la herramienta en foros relevantes (por ejemplo, StackOverflow o Reddit). Esto se debe cumplir para que la resolución de problemas sea más rápida, ya que cuando se tenga que hacer frente a un inconveniente es probable que existan usuarios que ya hayan tenido experiencias similares y puedan ofrecer respuestas. Además, la presencia de discusiones en foros  también ayuda a agilizar la resolución de inconvenientes.

---

## Candidatos:

- **Mage:** Si bien su comunidad es grande, las actualizaciones de este gestor cuentan con largos períodos de inactividad en su desarrollo, esto causa dudas sobre si podría incrementar la deuda técnica en proyectos que dependan de la herramienta. Estos datos, las estrellas y los forks junto a su documentación se puede consultar en su [Repositorio oficial](https://github.com/magefile/mage)

- **Gotaskr:** Esta diseñado para facilitar la ejecución de tareas automatizadas desde la línea de comandos, permite definir tareas usando código Go de forma directa lo que hace el proceso más accesible y sencillo para los desarrolladores familiarizados con el lenguaje. Tiene una extensión en VSCode, pero su comunidad es muy pequeña y su última actualización fue hace 2 meses. Esto junto a sus estrellas y número de forks se puede consultar en su [Repositorio oficial](https://github.com/Roemer/gotaskr)

- **Make:** Herramienta muy conocida y consolidada. No recibe actualizaciones desde hace un tiempo pero ya está establecida y cuenta con una gran comunidad. Su documentación se puede consultar en [este enlace](https://www.gnu.org/software/make/manual/make.html) y [aquí](https://lists.gnu.org/archive/html/info-gnu/2023-02/msg00011.html) podemos ver cuando fue su última actualización.

- **Task:**  Su comunidad es grande y recibe actualizaciones constantes, se utiliza un archivo `Taskfile.yml` en el que se definen las tareas. Se puede ver en el [Repositorio oficial](https://github.com/go-task/task)

- **Gilbert:** Gilbert es un sistema de construcción y gestor de tareas para proyectos en Go. Permite definir y ejecutar tareas de forma declarativa utilizando un archivo gilbert.yaml. No se actualiza desde 2019 y su comunidad no es grande, esto se pueden consultar en el [Repositorio oficial](https://github.com/go-gilbert/gilbert) 

- **Just:** Ofrece una manera fácil y flexible de automatizar tareas. Tiene una gran comunidad, con más de 22.7k estrellas y 495 forks, además de actualizaciones constantes. [Repositorio oficial](https://github.com/casey/just)


Para más información sobre estos gestores, se pueden consultar las siguientes fuentes además de sus repositorios:


- [Reddit](https://www.reddit.com/r/golang/search/?q=task+runner&cId=f6cebf5a-fcad-4cee-bd76-bc6493023d7e&iId=a1d557cc-da56-4106-ac2d-7d9178d8a798)
- [AwesomeGo](https://awesome-go.com/)

---

## Gestor elegido:

Después de haber descartado Mage, Gilbert y Gotaskr por no cumplir con ninguno de los criterios, quedan como opciones Make, ya que si bien no recibe actualizaciones regulares es una herramienta madura y estable, Just y Task. Descarto Make frente a Just y Task ya que estas dos opciones me parecen menos complejas de usar. Entre Just y Task, decido elegir Task ya que está dirigido a proyectos escritos en Go, además cómo ya he mencionado previamente cuenta con una gran comunidad y se pueden consultar discusiones sobre él en sitios como [Stack Overflow](https://stackoverflow.com/questions/tagged/taskfile).
