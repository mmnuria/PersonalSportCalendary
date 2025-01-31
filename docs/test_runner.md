# Requisitos para la selección de un test runner
1. Herramienta actualizada, para evitar el incremento de la deuda técnica.
2. Debe considerarse (en caso de que lo hubiera) las herramientas de la biblioteca estándar.

# Posibles opciones

## **Go testing framework**   
El paquete `testing` es el marco de pruebas por defecto en Go, adecuado para proyectos pequeños o medianos. Todas las pruebas en Go, independientemente de la biblioteca utilizada, se ejecutan con el comando `go test`, lo que garantiza integración y compatibilidad con el ecosistema del lenguaje.  

## **Ginkgo**  
Ginkgo es un test runner basado en BDD (Behavior Driven Development) que ofrece una mayor flexibilidad y funcionalidad avanzada para proyectos más complejos. Es adecuado para aquellos que necesitan un enfoque más estructurado y detallado en sus pruebas. Sin embargo, para proyectos sencillos, la complejidad adicional de Ginkgo puede no ser necesaria.
[Documentación oficial](https://github.com/onsi/ginkgo)

## **goblin**
Es un framework inspirado en Mocha (JavaScript) que permite escribir pruebas con una sintaxis concisa y basada en descripciones. No obstante, no se actualiza desde hace cuatro años.
[Documentación oficial](https://github.com/franela/goblin)

## **gotestsum**
Un test runner que proporciona resúmenes detallados y legibles de los resultados de las pruebas. No se ha actualizado desde hace ocho meses.
[Documentación oficial](https://github.com/gotestyourself/gotestsum) 

# Justificación de la selección

Se ha optado por utilizar el framework de pruebas de Go (`testing`) debido a:

1. Es el estándar del lenguaje.
2. Actualizaciones regulares y mantenimiento constante.

**goblin** y **gotestsum** no han sido recientemente actualizadas y respecto a **Ginkgo**, se considera una opción válida para proyectos que requieran un enfoque BDD, pero no es necesario para la mayoría de los proyectos Go que se benefician del marco de pruebas básico (como es el caso de este proyecto) proporcionado por `testing`.