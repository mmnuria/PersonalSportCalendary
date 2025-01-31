# Criterios de selección de la herramienta para registro de actividad (LOGS) 

1. **Mantenimiento y actualizaciones**: La herramienta debe estar bien mantenida y contar con actualizaciones frecuentes, lo que podemos verificar a través de su repositorio en GitHub.
2. **Reducción de dependencias**: Se considerará la herramienta que agregue el menor número de dependencias al proyecto, evitando sobrecargar el proyecto con bibliotecas innecesarias.
3. **Logs estructurados**: Se priorizarán herramientas que permitan el uso de múltiples formatos de configuración (JSON, YAML, TOML, ENV).

## Comparación de herramientas

1. **Log estándar de Go (`log`)**:  
   - [Documentación](https://golang.org/pkg/log/)
   - **Mantenimiento**: Parte del paquete estándar de Go, lo que garantiza que se mantenga actualizado con cada nueva versión de Go.
   - **Dependencias**: No requiere dependencias adicionales ya que es parte del estándar de Go.
   - **Características**: Es últil para proyectos pequeños y ofrece opciones de logging básicas. Sin embargo, carece de funcionalidades avanzadas como soporte de logs estructurados, rotación de logs, o integración con servicios externos.

2. **Zap** (de Uber):  
   - [GitHub](https://github.com/uber-go/zap)
   - **Mantenimiento**: Es una biblioteca muy activa y actualizaciones frecuentes. Último commit fue hace 2 semanas.
   - **Dependencias**: Zap tiene pocas dependencias, lo que lo hace liviano. Se recomienda usar `zapcore` para la personalización del logging.
   - **Características**: Es muy rápido y eficiente en términos de rendimiento. Ofrece soporte para logs estructurados en JSON, niveles de logs, y configuraciones personalizadas, lo que lo hace adecuado para proyectos más grandes y complejos.

3. **Logrus**:  
   - [GitHub](https://github.com/sirupsen/logrus)
   - **Mantenimiento**: Logrus tiene un buen historial de actualizaciones, aunque la actividad en el repositorio ha disminuido ligeramente en los últimos meses. Último commit hace 2 meses.
   - **Dependencias**: Requiere algunas dependencias adicionales, como `logrus` mismo y otros paquetes para formatos de salida avanzados.
   - **Características**: Ofrece un sistema de niveles de logs, y permite crear logs estructurados (JSON, texto).

4. **ZeroLog**:  
   - [GitHub](https://github.com/rs/zerolog)
   - **Mantenimiento**: ZeroLog es otra herramienta activa, con actualizaciones regulares y un enfoque en el rendimiento. Último commit hace menos de 1 mes.
   - **Dependencias**: Al igual que Zap, ZeroLog es muy liviano y no requiere muchas dependencias adicionales.
   - **Características**: Ofrece logging estructurado en JSON y se enfoca en ser extremadamente rápido y eficiente, con un menor consumo de memoria. Ideal para aplicaciones de alto rendimiento que necesiten registrar grandes volúmenes de datos.

## Conclusión

Tras evaluar las diferentes opciones en función de los criterios de mantenimiento y actualización, así como de la reducción de dependencias, la herramienta que mejor se adapta a las necesidades de nuestro proyecto es **Zap**.

1. Mantenimiento activo.
2. No requiere de dependencias externas.
3. Manejo de logs estructurados.

Si el proyecto requiere una solución de logging más simple, sin necesidad de características avanzadas, **Log estándar de Go** podría ser una buena opción.

Por tanto, **Zap** es la herramienta de logging elegida para este proyecto.