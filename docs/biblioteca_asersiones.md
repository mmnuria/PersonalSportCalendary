# Requisitos para la elección de la biblioteca de aserciones
1. Herramienta actualizada para evitar deuda técnica.
2. Debe permitir realizar aserciones simples y claras, con la capacidad de proporcionar mensajes personalizados.
3. La herramienta no debe basarse en un enfoque BDD, ya que este proyecto no requiere esa complejidad adicional.

# Opciones de la biblioteca de aserciones

**GoCheck**: Aunque GoCheck ofrece funcionalidades útiles, su falta de actualizaciones durante cinco años genera una alta deuda técnica. Por esta razón, no se considera una opción viable para el proyecto.
[Documentación oficial](https://github.com/go-check/check)

**Gomega**: Gomega es una biblioteca robusta y actualizada que proporciona amplias funcionalidades para pruebas, incluyendo soporte para mensajes detallados. Sin embargo, su enfoque principal en BDD introduce una complejidad innecesaria para este proyecto.
[Documentación oficial](https://github.com/onsi/gomega)

**Testify**: Testify es una biblioteca de aserciones ampliamente utilizada en Go. Facilita la realización de pruebas unitarias con un conjunto de funciones de aserción personalizables, que se integran perfectamente con `go test`. Recibe actualizaciones frecuentes, lo que garantiza su sostenibilidad a largo plazo.
[Documentación oficial](https://github.com/stretchr/testify) 

### Justificación.
Se ha seleccionado **Testify** como la herramienta de aserciones para el proyecto debido a:

1. Su mantenimiento constante (última actualización hace menos de un mes).
3. Su capacidad para realizar aserciones simples con mensajes personalizados.
4. Su integración nativa con go test, evitando configuraciones adicionales.