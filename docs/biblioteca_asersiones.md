# Requisitos para la elección de la biblioteca de aserciones
1. Herramienta actualizada para evitar deuda técnica.
2. Debe permitir realizar aserciones que permitan proporcionar mensajes personalizados.

# Opciones de la biblioteca de aserciones

**GoCheck**: Aunque GoCheck ofrece funcionalidades útiles, su falta de actualizaciones durante cinco años genera una alta deuda técnica. Por esta razón, no se considera una opción viable para el proyecto.
[Documentación oficial](https://github.com/go-check/check)

**Gomega**: Gomega es una biblioteca robusta y actualizada que proporciona amplias funcionalidades para pruebas, incluyendo soporte para mensajes detallados. 
[Documentación oficial](https://github.com/onsi/gomega)

**Testify**: Testify es una biblioteca de pruebas que incluye, entre otras funcionalidades, un paquete llamado **assert** que permite realizar aserciones en las pruebas unitarias. Es ampliamente utilizada en Go, facilita la integración con `go test` y se actualiza regularmente. Testify ofrece funciones de aserción personalizables y soporta mensajes detallados. Aunque es una herramienta más completa (que incluye también un paquete de mocking), su componente de aserciones es perfectamente adecuado para el proyecto.

[Documentación oficial](https://github.com/stretchr/testify) 

**GoConvey**: GoConvey es otro framework de pruebas que ofrece una integración fácil con Ginkgo, y se mantiene actualizado. Su sintaxis para aserciones es bastante legible, pero podría introducir una sobrecarga innecesaria. Ofrece soporte para mensajes personalizados en las aserciones.  
[Documentación oficial](https://github.com/smartystreets/goconvey)

**Assert**: Assert es una biblioteca sencilla, de fácil integración y mantenimiento, que permite realizar aserciones personalizadas. Aunque es más básica que las opciones anteriores, ofrece una forma rápida y directa de realizar pruebas.  
[Documentación oficial](https://github.com/go-playground/assert)

### Justificación.

Se ha seleccionado **Testify** como la herramienta de aserciones para el proyecto debido a:

1. Su mantenimiento constante (última actualización hace menos de un mes).
2. Su capacidad para realizar aserciones con mensajes personalizados.
3. Su integración nativa con `go test`, evitando configuraciones adicionales.
4. Es una biblioteca muy popular y ampliamente utilizada, lo que garantiza un amplio soporte comunitario.

Alternativamente, **GoConvey** y **Assert** son opciones viables, aunque **Testify** sobresale por su frecuencia de actualizaciones, lo que la hace la opción preferida para este proyecto.