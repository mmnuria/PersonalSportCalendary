### **Definición de las rutinas de ejercicios**

Cada **rutina** es una combinación de ejercicios individuales, previamente definidos y estructurados en el sistema.

Cada **ejercicio** debe tener los siguientes atributos en la base de datos:
    - **Nombre del ejercicio**
    - **Duración/Series**
    - **Instrucciones detalladas**
    - **Materiales necesarios**
    
Las **rutinas** se generan de forma dinámica en función de las preferencias del usuario, permitiendo flexibilidad en tiempo y equipo. A continuación desarrollo el paso a paso desde la aplicación:

1. **Pantalla de inicio:** Al abrir la aplicación, el usuario sin necesidad de iniciar sesión responde a algunas preguntas iniciales:
    - Tiempo disponible para entrenar cada día (5-30 minutos).
    - Cuantos días de la semana dispone.
    - Materiales disponibles (sin equipo o señalar que materiales poseé de las opciones que le ofrece la app).

2. **Pantalla de la rutina de ejercicios:** El usuario ve una lista de rutinas sugeridas basadas en su configuración (se generará una lista con las rutinas para una semana segun los dias disponibles que se haya seleccionado):
    - Ejercicio de calentamiento: 5 minutos.
    - Ejercicios principales: Sentadillas, flexiones, plancha (con número de repeticiones y series).
    - Enfriamiento: Estiramientos básicos.

3. **Descripción detallada de cada ejercicio:** Se debe de muestrar para cada ejercicio:
    - **Nombre del ejercicio.**
    - **Duración:** Ejemplo, 3 series de 10 sentadillas.
    - **Instrucciones:** Instrucciones paso a paso con imágenes o videos.
    - **Materiales necesarios:** Detalle de cada uno de los materiales a necesitar.