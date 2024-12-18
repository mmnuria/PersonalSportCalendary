# Selección de herramienta para la gestión de tareas

## **Criterios para elegir un gestor de tareas**

Para garantizar un flujo de trabajo eficiente y mantener la calidad del proyecto, es crucial seleccionar una herramienta que sea adecuada para el desarrollo con Go. Los criterios seleccionados son: 

- **Actualización constante y minimización de deuda técnica**: La herramienta debe contar con soporte activo y actualizaciones regulares, para reducir riesgos de mantenimiento a largo plazo.    
- **Documentación y soporte comunitario**: Se evaluará por la cantidad de issues resueltos en el repositorio oficial durante el último mes y la disponibilidad de ejemplos claros en la documentación oficial.

## **Análisis de opciones disponibles**

### **Make**  
Make es una herramienta ampliamente utilizada para la automatización de tareas, soportada de forma robusta en muchos lenguajes y entornos. Su uso en proyectos Go puede ser más complicado debido a la necesidad de una estructura más rigurosa y la gestión explícita de dependencias. 

[Documentación oficial de Make](https://github.com/wkusnierczyk/make?tab=readme-ov-file)

### **Just**  
Just ofrece una alternativa más sencilla y enfocada, ideal para proyectos Go. Su diseño simple y directo favorece la definición de tareas sin complicar la estructura ni requerir el manejo explícito de dependencias, lo que lo convierte en una opción más accesible.

[Documentación oficial de Just](https://github.com/casey/just)

### **Task**  
Task utiliza un enfoque basado en archivos YAML para definir tareas. Aunque es potente, esta configuración puede ser innecesariamente compleja para proyectos más simples.  

[Documentación oficial de Task](https://github.com/adriancooney/Taskfile) 

### **Mage**  
Mage destaca por permitir la definición de tareas directamente en Go, lo que elimina la necesidad de aprender una sintaxis diferente. Sin embargo, su falta de actualizaciones recientes puede ser un inconveniente.   

[Documentación oficial de Mage](https://github.com/magefile/mage) 

### **Sage**  
Inspirado en Mage, Sage permite automatizar tareas dentro de proyectos en Go. Se posiciona como un reemplazo para herramientas tradicionales como Make y Just, pero al ser relativamente nueva, no mantiene un soporte sólido.

[Repositorio oficial de Sage](https://github.com/einride/sage)

## **Evaluación de los gestores**

Para seleccionar la herramienta más adecuada en la gestión de tareas en este proyecto, evaluaremos y compararemos cada una de las opciones disponibles con base en los criterios establecidos previamente:

- **Evaluación de actualización constante y minimización de deuda técnica**:  
   - **Just** es la opción más sólida, ya que cuenta con actualizaciones constantes.
   - **Make** cumple con la minimización de deuda técnica gracias a su flujo eficiente, aunque las pocas actualizaciones aumentarán la deuda técnica.  
   - **Task** presenta un desarrollo más modesto y, aunque funcional, el uso de archivos YAML puede aumentar los costos de mantenimiento si no se gestionan adecuadamente.  
   - **Mage** no ha recibido actualizaciones recientes, lo que supone un riesgo de deuda técnica futura, y su enfoque puede generar configuraciones complicadas.  
   - **Sage** es relativamente nueva, con menor madurez, lo que podría introducir complejidad adicional en proyectos grandes.   

- **Documentación y soporte comunitario**
  - **Just y Mage**: La comunidad es más activa, con un número considerable de *issues* cerrados. La documentación es extensa y proporciona ejemplos claros que ayudan en la implementación y solución de problemas.  
  - **Make**: El repositorio no muestra actividad reciente en la resolución de *issues*, lo que sugiere una comunidad inactiva. Además, la documentación es limitada y no incluye ejemplos claros, dificultando su uso.  
  - **Task**: La comunidad es pequeña, con pocas *issues* cerradas en el último mes. La documentación también es limitada y carece de ejemplos claros, lo que puede dificultar la resolución de problemas. 
  - **Sage**: Al ser relativamente nuevo, el repositorio tiene una comunidad en crecimiento, con pocos *issues* resueltos hasta ahora. La documentación aún se está desarrollando, lo que limita la disponibilidad de ejemplos claros y soporte inmediato.  

### Elección y justificación

Después de evaluar las opciones, **Just** se integra perfectamente con el ecosistema Go y proporciona una solución confiable y fácil de mantener. Por lo tanto, esta herramienta es la más adecuada para el proyecto. Al cumplir con todos los criterios establecidos, permite mantener el proyecto organizado, reducir la deuda técnica y automatizar tareas esenciales de manera efectiva. 
