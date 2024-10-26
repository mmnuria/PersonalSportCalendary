## Milestones

### [M0] Modelo del problema

+ **Descripción:** Para poder comenzar con la resolución del problema (HU001), es necesario contar con una base estructurada de datos con la que apoyarte para la extracción de los ejercicios específicos para cada [rutina.](./rutina-ejercicios.md)

+ **Producto mínimamente viable:** Obtener fuentes de datos de las que se pueda extraer la información para poder modelar (elementos claves de dominio) el problema de forma correcta.

+ **Criterio de Validación**:
  - Especificación de las diferentes bases de datos en las que se ha apoyado.
  - Construir una base de datos estructurada organizada en ejercicios que se clasifiquen en función del objetivo con el que se realiza, es decir, mejorar flexibilidad, resistencia, fuerza, cardiovasculares, calentamiento y estiramiento.

---

### [M1] Generación de rutinas sin material deportivo

+ **Descripción**: Teniendo en cuenta los resultados del M0, implementar un generador de rutinas de ejercicio que no requieran material deportivo, abordando la problemática planteada en HU003. Este generador debe crear rutinas efectivas que aprovechen ejercicios de peso corporal y movimientos básicos que se puedan realizar desde casa. El algoritmo debe ser capaz de adaptar la duración, cantidad de dias disponibles y la intensidad señalada por el usuario para la generación de rutinas personalizadas sin material deportivo.

+ **Producto mínimamente viable**: Código funcional que genere rutinas de ejercicios efectivos sin equipos especializados, permitiendo a los usuarios mantenerse activos en casa.
+ **Criterio de Validación**:
  - El generador de rutinas será validado mediante pruebas en distintos escenarios, comprobando que las rutinas se ajustan al tiempo disponible y las preferencias del usuario, mientras que utilizan solo el peso corporal.

---

### [M2] Optimización de rutinas con material deportivo limitado

+ **Descripción**: Teniendo en cuenta los resultados del M1, implementar un sistema que permita la personalización de las rutinas de ejercicio basadas en el material deportivo limitado que los usuarios tienen en casa, como se menciona en la HU002. Este sistema debe considerar equipos como pesas pequeñas, bandas elásticas, esterillas... Y ajustar las rutinas para optimizar su uso. Además, se debe asegurar que las rutinas sean variadas y adapten los ejercicios a los recursos disponibles, brindando una experiencia personalizada.
+ **Producto mínimamente viable**: Código funcional que permita generar rutinas que integren el material deportivo disponible, complementando el generador de rutinas sin material desarrollado en el M1.
+ **Criterio de Validación**:
  - El código será validado mediante pruebas en situaciones simuladas, comprobando que las rutinas utilizan de forma óptima los equipos disponibles en diferentes usuarios (que los ejercicios de la "rutina de forma óptima" incluyan cada uno de los materiales que posee el usuario) y que se mantienen adecuadas para el nivel de intensidad y especificaciones de tiempo del usuario.

---

### [M3] Prevención de lesiones con estiramientos personalizados

+ **Descripción**: Implementar un módulo que incluya estiramientos personalizados antes y después de las rutinas de ejercicio para prevenir lesiones, en respuesta a la HU004. Este módulo debe generar recomendaciones específicas de estiramientos basadas en los tipos de ejercicios de la rutina y el nivel de intensidad seleccionado. Además, el sistema debe proporcionar una secuencia de estiramientos fácil de seguir que ayude a mejorar la flexibilidad y el bienestar general del usuario.
+ **Producto mínimamente viable**: Código funcional que genere estiramientos adecuados para cada rutina de ejercicios, ajustados al nivel de actividad (intensidad) y tipo de ejercicios realizados.
+ **Criterio de Validación**:
  - El sistema será validado mediante pruebas que verifiquen la correcta integración de los estiramientos con las rutinas, asegurando que las recomendaciones son pertinentes para el tipo de ejercicio realizado y el nivel de intensidad de la rutina.