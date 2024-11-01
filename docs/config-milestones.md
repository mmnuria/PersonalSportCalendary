## Milestones

### [M0] Estructura del problema (interno)

+ **Descripción:** **Módulo que defina la estructura de datos** que va almacenar los ejercicios para la obtención de rutinas y la interfaz del mismo, sin implementación.

+ **Hus asociadas** HU001.

+ **Criterio de Validación**:
  - Módulo que defina la estructura de datos del problema (modelo lógico de cómo se organizarán los ejercicios y las [rutinas](./rutina-ejercicios.md) dentro del sistema).

---

### [M1] Implementación de lógica de negocio (interno)

+ **Descripción:** Desarrollo del **módulo de generación de rutinas** utilizando la estructura de datos definida en el primer milestone. Este módulo debe contener la lógica de negocio para crear rutinas personalizadas basadas en diferentes criterios como el tiempo y los días disponibles. Hay que tener en cuenta que estas rutinas se realizarán en casa por lo que, no deben requerir de materiales deportivos.

+ **HUs asociadas**: HU003.

+ **Criterio de Validación**:
  - El sistema debe ser capaz de generar una rutina adaptada a la información ingresada por el usuario.
  - Los ejercicios deben ser seleccionados y organizados según el tiempo disponible definido en el input.
  - Pruebas unitarias para asegurar que la lógica de negocio selecciona ejercicios adecuados en función del tiempo del usuario.