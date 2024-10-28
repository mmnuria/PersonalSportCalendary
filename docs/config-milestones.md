## Milestones

### [M0] Extracción y organización de información para la base de datos

+ **Descripción:** Este milestone se centra en obtener y organizar la información relevante sobre ejercicios y [rutinas](./rutina-ejercicios.md) de manera que permita crear una base de datos inicial específica para el problema (HU001). Este paso es fundamental para tener una estructura de datos clara que respalde la generación de rutinas personalizadas, de acuerdo con las necesidades de los usuarios.

+ **Producto mínimamente viable:** Una base de datos inicial que contenga una estructura organizada de ejercicios clasificados según objetivos clave (flexibilidad, resistencia, fuerza, cardiovascular, calentamiento y estiramiento). La información debe extraerse de fuentes de datos confiables y relevantes, y organizarse de manera que facilite la posterior creación de rutinas específicas.

+ **Criterio de Validación**:
  - Documento que liste las fuentes de datos revisadas y utilizadas para construir la base de datos.
  - Base de datos estructurada que clasifique los ejercicios según el objetivo principal para el cual están diseñados, con atributos como nombre, duración, instrucciones, materiales necesarios e intensidad.
  - Validación del modelo de datos mediante consultas de prueba que permitan extraer ejercicios según los objetivos definidos.

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