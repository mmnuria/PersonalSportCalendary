# Requisitos para la selección de un test runner
1. Herramienta actualizada, para evitar el incremento de la deuda técnica.
2. Debe considerarse (en caso de que lo hubiera) las herramientas de la biblioteca estándar.

# Posibles opciones

## **Go testing framework**   
El paquete `testing` es el marco de pruebas por defecto en Go, adecuado para proyectos pequeños o medianos. Todas las pruebas en Go, independientemente de la biblioteca utilizada, se ejecutan con el comando `go test`, lo que garantiza integración y compatibilidad con el ecosistema del lenguaje.  

## **Ginkgo**  
Ginkgo es un test runner basado en BDD (Behavior Driven Development) que ofrece una mayor flexibilidad y funcionalidad avanzada para proyectos más complejos. Es adecuado para aquellos que necesitan un enfoque más estructurado y detallado en sus pruebas. Sin embargo, para proyectos sencillos, la complejidad adicional de Ginkgo puede no ser necesaria.
[Documentación oficial](https://github.com/onsi/ginkgo)


# Justificación de la selección

Se ha optado por utilizar el framework de pruebas de Go (`testing`) debido a:

1. Es de la librería estándar del lenguaje.
2. Actualizaciones regulares y mantenimiento constante.

**Ginkgo** se considera una opción válida para proyectos que requieran un enfoque BDD, pero no es necesario para la mayoría de los proyectos Go que se benefician del marco de pruebas básico (como es el caso de este proyecto) proporcionado por `testing`.