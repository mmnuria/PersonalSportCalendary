## Milestones

### [M0] Modelo del problema: Organización inicial de ejercicios

+ **Descripción**: Definir una estructura de datos básica para almacenar ejercicios en categorías funcionales y atributos clave. Esto servirá como base para que las rutinas se generen de acuerdo con las preferencias del usuario.

+ **Objetivo del Milestone**: 
   - Definir cómo los datos de ejercicios se organizan en una estructura accesible y comprensible para las futuras rutinas.
   - Establecer categorías de ejercicios y asegurar la inclusión de los atributos necesarios para la generación de rutinas efectivas.

+ **Detalles de la Estructura**:
   - **Categorías de Ejercicio**: Los ejercicios se agruparán en seis categorías:
     1. **Calentamiento**
     2. **Cardiovasculares**
     3. **Fuerza**
     4. **Flexibilidad**
     5. **Resistencia**
     6. **Estiramiento**

   - **Atributos de cada ejercicio**:
     - **Nombre**
     - **Duración/Series**
     - **Intensidad**: Baja, media o alta, para definir la demanda de esfuerzo.
     - **Instrucciones detalladas**: Explicación clara del ejercicio.
     - **Materiales necesarios**: Sólo en caso de necesitar equipo.
     
+ **Criterios de Validación**:
   - Documento que detalle la estructura de datos en formato JSON.
   - Ejemplos básicos de cada categoría con los atributos correspondientes.
   - La estructura debe permitir, sin modificación, añadir nuevas categorías y atributos específicos si se requieren en futuros milestones. 

### [M1] Implementación de lógica de negocio (interno)

+ **Descripción:** Desarrollo del **módulo de generación de rutinas** utilizando la estructura de datos definida en el primer milestone. Este módulo debe contener la lógica de negocio para crear rutinas personalizadas basadas en diferentes criterios como el tiempo y los días disponibles. Hay que tener en cuenta que estas rutinas se realizarán en casa por lo que, inicialmente no deben requerir de materiales deportivos (HU003).

+ **Criterio de Validación**:
  - El sistema debe ser capaz de generar una rutina adaptada a la información ingresada por el usuario.
  - Los ejercicios deben ser seleccionados y organizados según el tiempo disponible definido en el input.
  - Pruebas unitarias para asegurar que la lógica de negocio selecciona ejercicios adecuados en función del tiempo del usuario.