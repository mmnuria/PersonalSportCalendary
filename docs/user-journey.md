## User Journeys

### Laura
Laura es una madre trabajadora con tres hijos. Su trabajo le ocupa la mayor parte del día, y al regresar a casa, tiene poco tiempo para dedicarse a su bienestar. A menudo, se siente fatigada y preocupada por su falta de actividad física, pero sus obligaciones familiares y laborales no le permiten hacer ejercicio regularmente. Le gustaría encontrar una forma de mantenerse activa en los pocos minutos que tiene disponibles entre sus tareas.

---

### David
David es un ejecutivo de una empresa tecnológica que trabaja muchas horas frente al ordenador. Aunque es consciente de la importancia de hacer ejercicio, rara vez tiene tiempo para planificarlo o desplazarse a un gimnasio. Además, tiene algunas pesas en casa, pero no está seguro de cómo usarlas de manera efectiva. Se siente frustrado porque no quiere comprometer su salud, pero su vida laboral no le deja tiempo para buscar una solución.

---

### Teresa
Teresa es una mujer que, debido a la pandemia, se ha quedado sin acceso a su gimnasio habitual. Tiene algunas bandas elásticas en casa, pero no sabe cómo organizarse para entrenar con ellas. Siente que está perdiendo forma física y le gustaría encontrar una manera de aprovechar el material que tiene sin necesidad de adquirir más equipos o asistir a un gimnasio.

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

2. **Pantalla de la rutina de ejercicios:** El usuario ve una lista de rutinas sugeridas basadas en su configuración (se generará semanalmente para cada dia que haya mencionado):
    - Ejercicio de calentamiento: 5 minutos.
    - Ejercicios principales: Sentadillas, flexiones, plancha (con número de repeticiones y series).
    - Enfriamiento: Estiramientos básicos.

3. **Descripción detallada de cada ejercicio:** La aplicación muestra para cada ejercicio:
    - **Nombre del ejercicio.**
    - **Duración:** Ejemplo, 3 series de 10 sentadillas.
    - **Instrucciones:** Instrucciones paso a paso con imágenes o videos.
    - **Materiales necesarios:** Detalle de cada uno de los materiales a necesitar.

