## Milestones

### [M0] Milestone 0: **Modelo del problema** (HU001)

Crear un **modelo claro y funcional del problema** que incluya los elementos clave necesarios para la generación dinámica de rutinas de ejercicio. Este modelo debe permitir la personalización de rutinas basadas en parámetros específicos como tiempo disponible, intensidad, frecuencia semanal y materiales deportivos. Utilizando **Domain-Driven Design (DDD)**, este modelo servirá como base sólida para los futuros desarrollos.

#### Entregable:
Un **documento técnico** acompañado de código que defina:
1. Los principales **componentes del dominio** (por ejemplo, entidades como "usuario", "material deportivo", ["rutina de ejercicio"](./rutina-ejercicios.md)).
2. Relaciones y comportamientos básicos entre estos componentes.
3. **Interfaces** o descripciones de cómo estos componentes interactuarán para generar la rutina.

#### Criterios de éxito:
1. **Comprensión clara del problema por parte del receptor del milestone (otro estudiante)**:
   - El documento técnico y el código deben ser comprensibles y ejecutables por un miembro del equipo que no haya estado involucrado directamente en la creación del modelo.
   
2. **Validez del modelo en base a DDD**:
   - El modelo debe representar claramente los conceptos y reglas del dominio de forma flexible, permitiendo la expansión conforme el sistema crezca.

3. **Evolutividad**:
   - El modelo debe estar diseñado para ser adaptable a posibles cambios en las historias de usuario, evitando rigidez en la lógica de negocio.

4. **Criterio de validación objetiva**:
   - El milestone se considerará completado cuando el código sea revisado y aprobado por el equipo, siempre que pase pruebas que demuestren que es capaz de generar rutinas correctamente bajo diferentes combinaciones de parámetros. Además, debe ser extensible a futuros cambios y ser ejecutable sin errores por otros miembros del equipo.

---

### [M1] Milestone 1: **Implementación de la selección de tiempo e intensidad (HU002)**
Implementar la capacidad de seleccionar el tiempo disponible para ejercitarse y la intensidad de las rutinas.
- **Entregable:** Algoritmo que ajuste las rutinas según el tiempo e intensidad elegidos por el usuario.
- **Criterios de éxito:** El usuario puede especificar el tiempo que tiene disponible y la intensidad del ejercicio, recibiendo una rutina adecuada.

---

### [M2] Milestone 2: **Implementación de rutinas basadas en equipamiento junto con estiramientos (HU003, HU004 y HU005)**
Desarrollar la funcionalidad para que el usuario indique el equipamiento disponible (si lo tiene) y generar rutinas basadas en eso.

Incluir estiramientos al inicio y al final de las rutinas para prevenir lesiones y mejorar el bienestar general.

- **Entregables:** 

    1. Algoritmo que permita seleccionar equipamiento disponible y ofrezca rutinas adaptadas.

    2. Algoritmo que sugiera estiramientos apropiados al inicio y al final de cada sesión.

- **Criterios de éxito:** 

    1. El usuario puede indicar qué equipamiento tiene, y la aplicación genera una rutina basada en ello. Si no tiene material, se le ofrece una rutina sin equipo.
            
    2. Los estiramientos son recomendados en cada rutina, y son fáciles de realizar para prevenir lesiones.