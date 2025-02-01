# Requisitos para la elección de la biblioteca de aserciones
1. Herramienta actualizada para evitar deuda técnica. Para medir el estado de mantenimiento de un proyecto se usará ["Snyk Advisor"](https://snyk.io/advisor/golang) que aporta una métrica _Maintenance_ que indica si el proyecto está mantenido actualmente

# Opciones de la biblioteca de aserciones

**GoCheck**: La falta de actualizaciones en GoCheck desde hace cinco años genera una alta deuda técnica. Por esta razón, no se considera una opción viable para el proyecto. En Snyk Advisor no aparece un análisis al respecto.
[Documentación oficial](https://github.com/go-check/check)

**Gomega**: Gomega es una biblioteca robusta y actualizada que proporciona amplias funcionalidades para pruebas. Aparece como [_Healthy_](https://snyk.io/advisor/golang/github.com/onsi/gomega) en el apartado de mantenimiento de Snyk Advisor.
[Documentación oficial](https://github.com/onsi/gomega)

**Assert**: Assert es un paquete que forma parte de la biblioteca testify y permite realizar aserciones en las pruebas unitarias. Es ampliamente utilizado en Go, facilita la integración con `go test` y se actualiza regularmente. Aunque no existe el análisis del repositorio original por parte de Snyk, si que figuran algunos [análisis](https://snyk.io/advisor/golang/github.com/01ne/testify) de forks de hace tres meses de este mismo proyecto que tienen un rating de _Sustainable_ en el apartado de mantenimiento. 
[Documentación oficial](https://pkg.go.dev/github.com/stretchr/testify/assert) 

**GoConvey**: GoConvey es otro framework de pruebas que ofrece una integración fácil con Ginkgo, no obstante lleva sin actualizarse desde hace casi dos años.  
[Documentación oficial](https://github.com/smartystreets/goconvey)

### Justificación.

Aunque **Assert** es una opción viable, se ha seleccionado el paquete **Gomega** como la herramienta de aserciones para el proyecto debido a contar con el mejor mantenimiento de las cuatro opciones disponibles, lo que se ve reflejado en tener la mejor puntuación en Snyk.