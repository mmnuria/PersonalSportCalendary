# Requisitos para la selección de un test runner
1. Herramienta actualizada, para evitar el incremento de la deuda técnica.
2. Debe ajustarse al estándar del lenguaje, si este proporciona una herramienta oficial o ampliamente adoptada.

# Opciones que cumplen los criterios establecidos

## **Go testing framework**  
El framework de pruebas por defecto en Go es el paquete `testing`, este se encarga de definir las pruebas. El test runner de Go es el comando `go test`, que invoca el marco de pruebas y ejecuta las pruebas. Para proyectos pequeños o medianos, este test runner es suficiente.

## **Ginkgo**  
Ginkgo es un test runner basado en BDD (Behavior Driven Development) que ofrece una mayor flexibilidad y funcionalidad avanzada para proyectos más complejos. Es adecuado para aquellos que necesitan un enfoque más estructurado y detallado en sus pruebas. Sin embargo, para proyectos sencillos, la complejidad adicional de Ginkgo puede no ser necesaria.
[Documentación oficial](https://github.com/onsi/ginkgo)

# Justificación de la selección

Se ha optado por utilizar el framework de pruebas de Go (`testing`) debido a:

1. Su integración nativa con Go y la ejecución con el comando `go test`.
2. Actualizaciones regulares y mantenimiento constante.
3. Es la herramienta oficial, lo que asegura su compatibilidad y fiabilidad a largo plazo.

**Ginkgo** se considera una opción válida para proyectos que requieran un enfoque BDD, pero no es necesario para la mayoría de los proyectos Go que se benefician del marco de pruebas básico (como es el caso de este proyecto) proporcionado por `testing`.