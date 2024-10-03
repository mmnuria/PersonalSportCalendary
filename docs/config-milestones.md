# Creación de los milestones en Github

En el momento en el que ya tengamos los issues creados podemos proceder a la creación de los llamados milestones. Estos son objetivos que nos vamos a ir marcando poco a poco para ir resolviendo las historias de usuario de uno a uno.

1. Accedemos al apartado de los [milestones](./images/milestones.png) y seleccionamos ["New milestone"](./images/new-milestone.png).

2. Añadimos título y descripción necesaria para explicar como vamos a llevar acabo la resolución de cada uno de nuestros issues creados, de forma que quede algo así:
    - [Milestone 0: Planificación inicial](./images/milestone0.png) 
        
        >
        >Analizar las historias de usuario para identificar y definir los **elementos clave del dominio** que serán utilizados en el desarrollo del MVP. Este hito se centrará en estructurar los conceptos del dominio de acuerdo a las necesidades del negocio, usando **Domain-Driven Design (DDD)** como metodología.
        >
        >**Entregables:**
        >
        >- **Definición de las entidades principales**, objetos de valor y agregados, basados en el análisis de las historias de usuario.
        >- **Modelo inicial del dominio** que represente las relaciones entre las entidades y otros elementos clave.
        >- Revisión de las **historias de usuario** con criterios de aceptación claros, alineados con el modelo de dominio.
        >
        >**Criterios de éxito:**
        >
        >- El **modelo del dominio** está definido y comprende las entidades clave, objetos de valor, agregados y sus relaciones.
        >- Todos los elementos del dominio han sido **validados por el revisor** y son comprensibles para el equipo.
        >- Las historias de usuario están alineadas con los elementos del dominio y los criterios de aceptación son claros.
        >- El milestone se dará por completado cuando el **modelo del dominio** sea validado y esté listo para guiar el desarrollo del MVP.


    - [Milestone 1: Implementación de la selección de tiempo (HU001)](./images/milestone1.png)

        >En este hito se implementará la funcionalidad que permite a los usuarios **seleccionar el tiempo que tienen disponible para ejercitarse, tanto el tiempo y los días** y generar una rutina adaptada a esa duración.
        >
        >Entregables:
        >
        >- Interfaz de usuario que permite seleccionar el tiempo disponible (10, 20, 30 minutos, etc). IMPORTANTE: solo aceptar minutos para evitar problemas con las unidades.
        >- Interfaz de usuario que permita seleccionar los días que dispone en la semana.
        >- Algoritmo que genera una rutina de ejercicios basada en el tiempo y días seleccionados.
        >- Base de datos de ejercicios.
        >
        >Criterios de éxito:
        >
        >El usuario puede seleccionar un rango de tiempo en minutos, la cantidad de días que puede realizar ejercicio y recibir una rutina adecuada a la duración elegida.
        Las rutinas generadas son funcionales y se adaptan a la disponibilidad del usuario.

    - [Milestone 2: Implementación de rutinas basadas en equipamiento (HU002)](./images/milestone2.png)

        >Este hito se enfocará en **permitir a los usuarios indicar qué material deportivo tienen disponible** o en caso de no tenerlo, haya una opción que lo recoja y generar rutinas que incorporen esos equipos o sin ellos.
        >
        >Entregables:
        >
        >- Interfaz para que el usuario seleccione los tipos de material deportivo (pesas, bandas elásticas, esterilla, etc.).
        >- Algoritmo que genera rutinas incorporando el material seleccionado o sin equipamiento dependiendo de lo seleccionado.
        >- Base de datos de ejercicios que se pueden realizar con los equipos materiales disponibles o sin ellos.
        >
        >Criterios de éxito:
        >
        >El usuario puede indicar el material disponible y recibe rutinas que lo utilizan. Si no hay material, la aplicación ofrece ejercicios con peso corporal.

    - [Milestone 3: Inclusión de ejercicios de estiramiento (HU003)](./images/milestone3.png)

        >Este hito incluirá la implementación de la funcionalidad para **sugerir ejercicios de estiramiento al inicio y al final** de cada rutina de ejercicios.
        >
        >Entregables:
        >
        >- Algoritmo que sugiere una serie de estiramientos al inicio de cada rutina.
        >- Algoritmo que sugiere estiramientos al finalizar la rutina.
        >- Base de datos de estiramientos que son de bajo impacto.
        >
        >Criterios de éxito:
        >
        >La aplicación ofrece estiramientos antes y después de cada rutina.
        Los estiramientos son relevantes para los ejercicios realizados y son fáciles de ejecutar en casa.

3. Por último debemos asegurarnos que, tras haber creado cada uno de los milestones anteriores deben de estar relacionados con el [issue](./config-issues.md) correspondiente (HU001, HU002 o HU003).

