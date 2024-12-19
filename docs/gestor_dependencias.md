### Gestión de dependencias en Go

Con la llegada de los módulos en **Go 1.11**, el manejo de bibliotecas externas dio un giro significativo hacia un sistema más eficiente y cohesionado. Antes de esta innovación, los desarrolladores confiaban en herramientas de terceros como **Dep**, **Glide**, **Godep** y **Govendor** para gestionar dependencias. Estas soluciones ofrecían cierta funcionalidad, pero eran propensas a problemas de compatibilidad, duplicación de esfuerzos y complejidad en la configuración.

Los módulos de Go eliminaron estas limitaciones al introducir un enfoque nativo y centralizado. Este sistema permite organizar el código y sus dependencias como una unidad versionada, facilitando un control más preciso sobre las bibliotecas utilizadas. A través de un diseño integrado, los módulos no solo simplificaron el flujo de trabajo, sino que también contribuyeron a estandarizar la manera en que los proyectos en Go gestionan sus dependencias.

En la práctica, los módulos funcionan mediante dos archivos esenciales que permiten estructurar proyectos y gestionar sus dependencias de manera eficiente. El archivo `go.mod` define las versiones exactas de las bibliotecas utilizadas y su procedencia, mientras que el archivo `go.sum` registra las sumas de verificación que aseguran la integridad y consistencia del proyecto. Gracias a esta estructura, Go elimina la necesidad de configuraciones manuales y garantiza un entorno de trabajo reproducible.

[Documentación oficial](https://go.dev/wiki/Modules) 