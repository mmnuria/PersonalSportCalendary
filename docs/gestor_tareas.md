# Selección de herramienta para la gestión de tareas

## **Criterio para elegir un gestor de tareas**

1. **Proyecto en activo y sujeto a mantenimiento continuo**:
  Se considerará el repositorio de código oficial para evaluar si los commits y las versiones de cada herramienta son recientes.

## **Análisis de opciones disponibles**

### **Make**  
- **Proyecto en activo y sujeto a mantenimiento continuo:** **Make** ha sido activamente mantenida durante décadas, ha logrado asegurar una alta fiabilidad y compatibilidad futura sin problemas, lo que ha reducido la necesidad de actualizaciones frecuentes. No obstante, no ha sido actualizada en los últimos dos años.

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

En primer lugar, descartamos **Mage** y **Make** ya que no han recibido actualizaciones en el último año.

A partir de lo expuesto anteriormente, podríamos elegir entre las herramientas **Sage**, **Just** y **Task**, siendo todas ellas opciones que cumplen el criterio. Dado que las tres son válidas, elijo **Just** por criterio propio.
