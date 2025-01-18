# Selección de herramienta para la gestión de tareas

## **Criterio para elegir un gestor de tareas**

1. **Proyecto en activo y sujeto a mantenimiento continuo**:
  Se considerará el repositorio de código oficial para evaluar la cantidad de commits y las versiones más recientes de cada herramienta. Adicionalmente, se valorará de forma positiva la antigüedad del proyecto como garantía de soporte.  

## **Análisis de opciones disponibles**

### **Make**  
- **Proyecto en activo y sujeto a mantenimiento continuo:** Aunque la herramienta se considera inactiva, se ha recibido un [último commit](http://git.savannah.gnu.org/cgit/make.git/refs/) hace 6 semanas, lo que indica que todavía se mantiene. Dado que **Make** ha sido activamente mantenida durante décadas, ha logrado asegurar una alta fiabilidad y compatibilidad futura sin problemas, lo que ha reducido la necesidad de actualizaciones frecuentes.

[Documentación oficial de Make](http://git.savannah.gnu.org/cgit/make.git)

### **Just**  
- **Proyecto en activo y sujeto a mantenimiento continuo:** Es un proyecto activo y bien gestionado con actualizaciones periódicas y un [historial de releases](https://github.com/casey/just/releases?page=1) constantes.

[Documentación oficial de Just](https://github.com/casey/just)

### **Task**  
- **Proyecto en activo y sujeto a mantenimiento continuo:** Es un proyecto activo cuyo ultimo release ha sido hace menos de un dia.

[Documentación oficial de Task](https://github.com/go-task/task)

### **Mage**  
- **Proyecto en activo y sujeto a mantenimiento continuo:** [No ha recibido](https://github.com/magefile/mage/releases) actualizaciones en el último año. 

[Documentación oficial de Mage](https://github.com/magefile/mage)

### **Sage**  
- **Proyecto en activo y sujeto a mantenimiento continuo:** Mantenimiento activo con gran cantidad de [releases publicados](https://github.com/einride/sage/pulse).

[Repositorio oficial de Sage](https://github.com/einride/sage)

## **Elección y justificación**

En primer lugar, descartamos **Mage** ya que no ha recibido actualizaciones en el último año.

A partir de lo expuesto anteriormente, podríamos elegir entre las herramientas **Sage**, **Just**, **Make** y **Task**, siendo todas ellas buenas opciones que cumplen el criterio. No obstante, **Make** está respaldada por décadas de mantenimiento activo, siendo la de mayor antigüedad y por tanto se decide seleccionarla para el proyecto.