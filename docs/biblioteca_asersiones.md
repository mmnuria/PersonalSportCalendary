# Requisitos para la elección de la biblioteca de aserciones
1. Herramienta actualizada.
2. La biblioteca de aserciones seleccionada no debe basarse en un enfoque BDD (Behavior-Driven Development), para la realización de aserciones simples con la capacidad de proporcionar mensajes personalizados.

# Opciones de la biblioteca de aserciones

**GoCheck**: Aunque GoCheck ofrece funcionalidades útiles, su falta de actualizaciones durante más de cuatro años genera una alta deuda técnica. Por esta razón, no se considera una opción viable para el proyecto.
[Documentación oficial](https://github.com/go-check/check)

**Gomega**: Gomega es una biblioteca robusta y actualizada que proporciona amplias funcionalidades para pruebas, incluyendo soporte para mensajes detallados y pruebas basadas en BDD. Sin embargo, su enfoque principal en BDD podría ser más de lo necesario para este proyecto.
[Documentación oficial](https://github.com/onsi/gomega)

**Testify**: Esta biblioteca es una de las más populares en el ecosistema de Go. Ofrece aserciones simples y efectivas, con soporte para personalizar mensajes de error. Además, Testify es ampliamente utilizada y recibe actualizaciones frecuentes, lo que garantiza su sostenibilidad a largo plazo.
[Documentación oficial](https://github.com/stretchr/testify) 

### Justificación.
Se ha seleccionado **Testify** como la herramienta de aserciones para el proyecto, debido a su amplia aceptación en la comunidad de Go y su mantenimiento constante. Su última actualización fue hace menos de un mes, lo que asegura soporte continuo. Además, Testify ofrece un conjunto de aserciones simples pero poderosas que se integran fácilmente con la herramienta de pruebas estándar `go test`. Su fiabilidad y versatilidad la convierten en una opción óptima para satisfacer los requisitos definidos.