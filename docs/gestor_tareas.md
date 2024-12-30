# Selección de herramienta para la gestión de tareas

## **Criterios para elegir un gestor de tareas**

- **Minimización de deuda técnica**: La herramienta debe contar con soporte activo y actualizaciones regulares, para reducir riesgos de mantenimiento a largo plazo.

- **Compatibilidad con el ecosistema Go**: La herramienta debe estar alineada con las prácticas y necesidades comunes de los proyectos desarrollados en Go, facilitando la integración sin requerir configuraciones adicionales complejas.

## **Análisis de opciones disponibles**

### **Make**  
Make es una herramienta ampliamente utilizada para la automatización de tareas, soportada de forma robusta en muchos lenguajes y entornos. Su uso en proyectos Go puede ser más complicado debido a la necesidad de una estructura más rigurosa y la gestión explícita de dependencias. 

[Documentación oficial de Make](http://git.savannah.gnu.org/cgit/make.git)

### **Just**  
Just ofrece una alternativa enfocada e ideal para proyectos Go. Su diseño favorece la definición de tareas sin complicar la estructura ni requerir el manejo explícito de dependencias, lo que lo convierte en una opción más accesible. Además, cuenta con un soporte activo.

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

- **Minimización de deuda técnica**:  
   - **Just** destaca por contar con actualizaciones frecuentes y soporte activo, lo que lo hace ideal para reducir la deuda técnica.
   - **Make** es una herramienta establecida, pero sus actualizaciones son menos frecuentes, lo que puede incrementar la deuda técnica con el tiempo. 
   - **Task** presenta un desarrollo más modesto y, aunque funcional, el uso de archivos YAML puede aumentar los costos de mantenimiento si no se gestionan adecuadamente.  
   - **Mage** no ha recibido actualizaciones recientes, lo que supone un riesgo de deuda técnica futura, y su enfoque puede generar configuraciones complicadas.  
   - **Sage** es relativamente nueva, con menor madurez, lo que podría introducir complejidad adicional en proyectos grandes.   

- **Compatibilidad con el ecosistema Go**:
   - **Just** está diseñado con un enfoque simple y directo, alineándose bien con las necesidades de los proyectos Go. Su integración es más natural en este ecosistema.
   - **Make**, aunque válido para proyectos Go, no está específicamente orientado hacia este lenguaje y puede requerir configuraciones adicionales para tareas comunes.
   - **Task** y **Mage** ofrecen flexibilidad, pero su uso puede ser excesivo para proyectos pequeños o medianos en Go.
   - **Sage** está pensado para Go, pero su inmadurez lo hace menos interesante.

### Elección y justificación

Tras evaluar las opciones disponibles, tanto **Make** como **Just** cumplen con los criterios establecidos y son válidas para este proyecto. Sin embargo, Just es preferible porque además de minimizar la deuda técnica, está más alineado con las prácticas comunes y necesidades de los proyectos Go, lo que facilita la integración y reduce el tiempo de configuración.

Este enfoque orientado hacia Go no solo simplifica la gestión de tareas, sino que también refuerza la sostenibilidad y eficiencia del proyecto a largo plazo. Por lo tanto, **Just se selecciona como la herramienta más adecuada.**
