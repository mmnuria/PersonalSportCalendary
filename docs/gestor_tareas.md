# Selección de herramienta para la gestión de tareas

## **Criterio para elegir un gestor de tareas**

1. **Comunidad activa y mantenimiento**:
  La herramienta Snyk Advisor o el propio github se utilizarán para evaluar la cantidad de issues, pull requests y las versiones más recientes de cada herramienta. 

## **Análisis de opciones disponibles**

### **Make**  
- **Comunidad y mantenimiento:** Aunque la herramienta se considera inactiva, se ha recibido un [último commit](http://git.savannah.gnu.org/cgit/make.git/refs/) hace 6 semanas, lo que indica que todavía se mantiene. Dado que **Make** ha sido activamente mantenida durante décadas, ha logrado asegurar una alta fiabilidad y compatibilidad futura sin problemas, lo que ha reducido la necesidad de actualizaciones frecuentes. Además, cuenta con una comunidad amplia y consolidada que sigue respaldando su uso.

[Documentación oficial de Make](http://git.savannah.gnu.org/cgit/make.git)

### **Just**  
- **Comunidad y mantenimiento:** Es un proyecto activo y bien gestionado, con una comunidad en expansión, actualizaciones periódicas y un [historial de releases](https://github.com/casey/just/releases?page=1) constante.

[Documentación oficial de Just](https://github.com/casey/just)

### **Task**  
- **Comunidad y mantenimiento:** Es un proyecto activo que tiene una comunidad en continuo crecimiento, pero su [último commit](https://github.com/adriancooney/Taskfile/commits/master/) fue hace 8 años.

[Documentación oficial de Task](https://github.com/adriancooney/Taskfile)

### **Mage**  
- **Comunidad y mantenimiento:** Amplia comunidad y estado de mantenimiento sostenible, aunque, [no ha recibido](https://github.com/magefile/mage/releases) actualizaciones en el último año. 

[Documentación oficial de Mage](https://github.com/magefile/mage)

### **Sage**  
- **Comunidad y mantenimiento:** Comunidad limitada, al ser una herramienta relativamente reciente. Mantenimiento activo con gran cantidad de [releases publicados](https://github.com/einride/sage/pulse).

[Repositorio oficial de Sage](https://github.com/einride/sage)

## **Elección y justificación**

A partir de lo expuesto anteriormente, podríamos elegir entre las herramientas **Sage**, **Just** y **Make**. En primer lugar, descartamos **Sage** debido a su falta de madurez. En cuanto a **Make** y **Just**, ambas son excelentes opciones para el proyecto. Sin embargo, debido a la gran madurez y fiabilidad de **Make**, respaldada por décadas de mantenimiento activo, hemos decidido seleccionarla para el proyecto.