# Requisitos para la selección de un test runner
1. Herramienta actualizada, para evitar el incremento de la deuda técnica.
2. Debe ajustarse al estándar del lenguaje, si este proporciona una herramienta estándar.

# Opciones que cumplen los criterios establecidos

## **Ginkgo**  
Ginkgo es un framework avanzado para pruebas basado en BDD, diseñado para manejar proyectos grandes y estructurados. Funciona de manera excelente en combinación con bibliotecas como Gomega para gestionar aserciones.  
[Documentación oficial](https://github.com/onsi/ginkgo)

## **Go test**  
`go test` es el test runner estándar incluido en el ecosistema de Go. Ofrece una integración directa con el lenguaje, siendo una opción ligera y eficiente para la mayoría de los proyectos.

# Justificación de la selección

Aunque Ginkgo representa una alternativa avanzada para proyectos complejos, se ha optado por utilizar `go test` debido a su condición de herramienta oficial del lenguaje y su alineación con los estándares de desarrollo en Go. Su simplicidad y capacidad para gestionar pruebas de manera eficiente lo convierten en una elección confiable para proyectos que no requieren un enfoque basado en BDD.