## Milestones

### [M0] Modelo del Problema

+ **Descripción**: Realizar un análisis exhaustivo de la HU001 para identificar y definir los elementos clave relacionados con la falta de tiempo para el ejercicio y la dificultad de planificación de éstos. Se debe establecer una metodología que permita modelar y representar las preferencias de los usuarios, los momentos disponibles para entrenar y cómo se relacionan con la creación de rutinas personalizadas. El objetivo es crear un modelo que represente de manera precisa el contexto y las necesidades de usuarios con limitaciones en el tiempo y cantidad de días que disponen para mantener un nivel adecuado de actividad física.

+ **Producto mínimamente viable**: Código que encapsule los elementos fundamentales del dominio (componentes clave que definen como funciona realmente el problema como por ejemplo definiendo "Usuario" con atributos de tiempo disponible, preferencias de intensidad...) y generando una estructura básica para rutinas de ejercicios (colección de ejercicios que se ajustan al tiempo y días que dispone a la semana el usuario para realizar la rutina).

+ **Requisitos para la Validación**:
  - El diseño debe seguir buenas prácticas de modelado, permitiendo que en futuras iteraciones se puedan añadir nuevas variables como materiales disponibles y estiramientos.
  - El modelo será validado mediante pruebas que simulen diferentes situaciones de usuario, asegurando que las rutinas propuestas se ajustan a la disponibilidad de tiempo y cantidad de días de la semana a realizar la rutina descrito por el usuario.

---

### [M1] Generación de Rutinas sin Material Deportivo

+ **Descripción**: Teniendo en cuenta lo resultados del M0, implementar un generador de rutinas de ejercicio que no requieran material deportivo, abordando la problemática planteada en HU003. Este generador debe crear rutinas efectivas que aprovechen ejercicios de peso corporal y movimientos básicos que se puedan realizar desde casa. El algoritmo debe ser capaz de adaptar la duración (recogida en el M0) y la intensidad (opción para el usuario a incorporar en este hito) para la generación de rutinas personalizadas sin material deportivo.

+ **Producto mínimamente viable**: Código funcional que genere rutinas de ejercicios efectivos sin equipos especializados, permitiendo a los usuarios mantenerse activos en casa.
+ **Criterio de Validación**:
  - El generador de rutinas será validado mediante pruebas en distintos escenarios, comprobando que las rutinas se ajustan al tiempo disponible y las preferencias de intensidad del usuario, mientras que utilizan solo el peso corporal.

---

### [M2] Optimización de Rutinas con Material Deportivo Limitado

+ **Descripción**: Teniendo en cuenta los resultados del M1, implementar un sistema que permita la personalización de las rutinas de ejercicio basadas en el material deportivo limitado que los usuarios tienen en casa, como se menciona en la HU002. Este sistema debe considerar equipos como pesas pequeñas, bandas elásticas, esterillas... Y ajustar las rutinas para optimizar su uso. Además, se debe asegurar que las rutinas sean variadas y adapten los ejercicios a los recursos disponibles, brindando una experiencia personalizada.
+ **Producto mínimamente viable**: Código funcional que permita generar rutinas que integren el material deportivo disponible, complementando el generador de rutinas sin material desarrollado en el M1.
+ **Criterio de Validación**:
  - El código será validado mediante pruebas en situaciones simuladas, comprobando que las rutinas utilizan de forma óptima los equipos disponibles en diferentes usuarios (que los ejercicios de la "rutina de forma óptima" incluyan cada uno de los materiales que posee el usuario) y que se mantienen adecuadas para el nivel de intensidad y especificaciones de tiempo del usuario.

---

### [M3] Prevención de Lesiones con Estiramientos Personalizados

+ **Descripción**: Implementar un módulo que incluya estiramientos personalizados antes y después de las rutinas de ejercicio para prevenir lesiones, en respuesta a la HU004. Este módulo debe generar recomendaciones específicas de estiramientos basadas en los tipos de ejercicios de la rutina y el nivel de intensidad seleccionado. Además, el sistema debe proporcionar una secuencia de estiramientos fácil de seguir que ayude a mejorar la flexibilidad y el bienestar general del usuario.
+ **Producto mínimamente viable**: Código funcional que genere estiramientos adecuados para cada rutina de ejercicios, ajustados al nivel de actividad (intensidad) y tipo de ejercicios realizados.
+ **Criterio de Validación**:
  - El sistema será validado mediante pruebas que verifiquen la correcta integración de los estiramientos con las rutinas, asegurando que las recomendaciones son pertinentes para el tipo de ejercicio realizado y el nivel de intensidad de la rutina.