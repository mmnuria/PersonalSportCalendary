# Criterios de selección para la herramienta para configuración

1. Se priorizarán herramientas que permitan el uso de múltiples formatos de configuración (JSON, YAML, TOML, ENV).
2. Se usará [Snyk Advisor](https://snyk.io/advisor/) y [pkg.go.dev](https://pkg.go.dev/), que proporcionan evaluaciones de seguridad, popularidad y mantenimiento de las librerías.
3. Herramienta mantenida y actualizada, a través de Snyk Advisor y sus repositorios de GitHub.  


## Comparación de herramientas

1. **Viper**:  
   [Snyk Advisor](https://snyk.io/advisor/golang/github.com/spf13/viper)  
   [GitHub](https://github.com/spf13/viper)  
   Como podemos ver en su página de Snyk Advisor, tiene una alta puntuación (91/100). Es compatible con JSON, YAML, TOML y variables de entorno. Además, permite la recarga dinámica de configuración. Su último commit fue hace 3 semanas, lo que indica un mantenimiento activo.

2. **Godotenv**:  
   [Snyk Advisor](https://snyk.io/advisor/golang/github.com/joho/godotenv)  
   [GitHub](https://github.com/joho/godotenv)  
   Como podemos ver en su página de Snyk Advisor, tiene una puntuación media-alta (69/100). Su principal ventaja es la facilidad de uso para cargar variables desde archivos `.env`. Sin embargo, carece de soporte para otros formatos de configuración como YAML o JSON. Es una opción adecuada para proyectos pequeños o cuando solo se usan variables de entorno. Su último commit fue hace más de 1 mes, lo que indica un mantenimiento activo.

3. **Koanf**:  
   [Snyk Advisor](https://snyk.io/advisor/golang/github.com/knadh/koanf/v2)  
   [GitHub](https://github.com/knadh/koanf)  
   Como podemos ver en su página de Snyk Advisor, tiene una puntuación de 89/100. Soporta múltiples formatos como JSON, YAML, TOML y variables de entorno. Es similar a Viper, pero con un diseño más modular. Su último commit fue el mes pasado, lo que indica un mantenimiento activo.

4. **Afero** (para configuraciones avanzadas con sistemas de archivos virtuales):  
   [Snyk Advisor](https://snyk.io/advisor/golang/github.com/spf13/afero)  
   [GitHub](https://github.com/spf13/afero)  
   Tiene una puntuación de 95/100 y es ideal cuando se requiere gestionar configuraciones almacenadas en distintos sistemas de archivos. Se suele combinar con Viper para mayor versatilidad. Su último commit fue hace 3 semanas, lo que indica un mantenimiento activo.

## Conclusión

Basándonos en los criterios establecidos, **Viper** es la herramienta más adecuada.

1. Soporte de configuración en múltiples formatos.
2. Puntuación elevada en Snyk Advisor 91/100.
3. Mantenimiento activo.

**Afero** tiene una puntuación mayor pero, no es una opción adecuada para la configuración debido a que, está enfocada en la gestión de archivos y sistemas de archivos virtuales, mientras que **Viper** y **Koanf** están mucho más orientadas a la gestión de configuraciones en distintos formatos.

Por lo tanto, otra opcion factible puede ser **Koanf** que también cumple los criterios establecidos. Si la configuración solo usa variables de entorno, **Godotenv** es una alternativa ligera y sencilla.  

Por tanto, **Viper** es la herramienta de configuración elegida para este proyecto.

