# Selección de herramienta para la gestión de tareas

## **Criterios para elegir un gestor de tareas**

1. **Comunidad activa**:  
   - La herramienta debe contar con al menos 10 contribuidores activos en su repositorio oficial en los últimos 12 meses.  
   - Debe tener actividad visible, como issues cerrados o preguntas respondidas, en foros como GitHub Discussions, Reddit o Stack Overflow en los últimos 6 meses.

2. **Mantenimiento regular**:  
   - La herramienta debe haber recibido al menos un release oficial en el último año.  
   - El repositorio debe mostrar actividad de commits en los últimos 6 meses.
   - La herramienta existe desde hace más de 5 años.

3. **Dependencias y configuración**:  
   - La herramienta no debe requerir dependencias externas adicionales para funcionar en los sistemas operativos principales (Windows, macOS, Linux).  

## **Análisis de opciones disponibles**

### **Make**  
- **Comunidad activa**:  
  Cumple. Amplia comunidad, soporte extenso en [foros](https://stackoverflow.com/questions/tagged/make).  
- **Mantenimiento regular**:  
  No cumple. Las actualizaciones son menos frecuentes, y el [último release](http://git.savannah.gnu.org/cgit/make.git/refs/) es de hace más de un año.  
- **Dependencias y configuración**:  
  No cumple en Windows, ya que requiere herramientas externas como `MinGW` para funcionar correctamente.

[Documentación oficial de Make](http://git.savannah.gnu.org/cgit/make.git)

### **Just**  
- **Comunidad activa**:  
  Cumple. [Contribuidores activos](https://github.com/casey/just/graphs/contributors?from=30%2F12%2F2023) en GitHub y soporte en foros.  
- **Mantenimiento regular**:  
  Cumple. Actualizaciones regulares y un [historial de releases](https://github.com/casey/just/releases?page=1) frecuentes.  
- **Dependencias y configuración**:  
  Cumple. No requiere dependencias externas.

[Documentación oficial de Just](https://github.com/casey/just)

### **Task**  
- **Comunidad activa**:  
  No cumple. No tiene contribuidores en el ultimo año (ni desde hace 8 años).  
- **Mantenimiento regular**:  
  No cumple. [Último commit](https://github.com/adriancooney/Taskfile/commits/master/) hace 8 años. 
- **Dependencias y configuración**:  
  No cumple. Requiere la configuración de archivos YAML.

[Documentación oficial de Task](https://github.com/adriancooney/Taskfile)

### **Mage**  
- **Comunidad activa**:  
  No cumple. Solo tiene [3](https://github.com/magefile/mage/graphs/contributors) contribuidores en el ultimo año.  
- **Mantenimiento regular**:  
  No cumple. [No ha recibido](https://github.com/magefile/mage/releases) actualizaciones en el último año.  
- **Dependencias y configuración**:  
  Cumple. No requiere dependencias externas.

[Documentación oficial de Mage](https://github.com/magefile/mage)

### **Sage**  
- **Comunidad activa**:  
  Cumple. [Contribuidores activos](https://github.com/einride/sage/graphs/contributors?from=30%2F12%2F2023) en el ultimo año.
- **Mantenimiento regular**:  
  No cumple. La herramienta existe desde 2021.  
- **Dependencias y configuración**:  
  Cumple. No requiere dependencias externas.

[Repositorio oficial de Sage](https://github.com/einride/sage)

## **Evaluación de las opciones**

| Herramienta | Comunidad activa | Mantenimiento regular | Dependencias y configuración |
|-------------|------------------|-----------------------|-----------------------------|
| **Make**    | ✔️               |❌                     | ❌                          |
| **Just**    | ✔️               | ✔️                    | ✔️                          |
| **Task**    | ❌               | ❌                    | ❌                          |
| **Mage**    | ❌               | ❌                    | ✔️                          |
| **Sage**    | ✔️               | ❌                    | ✔️                          |

### Leyenda
- **✔️**: Cumple con el criterio.
- **❌**: No cumple con el criterio.


## **Elección y justificación**

Con base en los criterios establecidos, **Just** es la herramienta que mejor cumple con los requisitos objetivos. Es mantenida regularmente, tiene una comunidad activa y no introduce dependencias externas. Esto la convierte en la opción ideal para gestionar las tareas en este proyecto.