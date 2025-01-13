# Requisitos para la selección de un test runner
1. Herramienta actualizada, para evitar el incremento de la deuda técnica.
2. Debe ajustarse al estándar del lenguaje, si este proporciona una herramienta estándar.

# Opciones que cumplen los criterios establecidos

## **Go test**  
El test runner por defecto de Go es go test. Este se encarga de ejecutar las pruebas, pero no proporciona las funciones de aserción, solo ejecuta los tests y muestra los resultados.

## **Ginkgo**  
Ginkgo es un framework de pruebas basado en BDD, adecuado para proyectos más complejos. Sin embargo, para la mayoría de los proyectos Go, go test es suficiente y no es necesario añadir frameworks adicionales, a menos que se sigan metodologías BDD.
[Documentación oficial](https://github.com/onsi/ginkgo)

# Justificación de la selección

Se ha optado por utilizar `go test` debido a:

1. Actualizaciones regulares.
2. Su condición de herramienta oficial del lenguaje. 