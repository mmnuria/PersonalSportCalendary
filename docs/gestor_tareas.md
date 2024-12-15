# Selección de herramienta para la gestión de tareas

## **Criterios para elegir un gestor de tareas**

Para garantizar un flujo de trabajo eficiente y mantener la calidad del proyecto, es crucial seleccionar una herramienta que sea adecuada para el desarrollo con Go. Los criterios seleccionados son: 

- **Actualización constante**: La herramienta debe contar con soporte activo y actualizaciones regulares para evitar problemas de compatibilidad.  
- **Promoción de buenas prácticas**: Debe alinearse con las características y necesidades del ecosistema Go, facilitando la adopción de patrones estándar.  
- **Facilidad de uso**: Una sintaxis intuitiva es esencial para que los desarrolladores puedan aprender y usar la herramienta rápidamente, reduciendo curvas de aprendizaje innecesarias.   
- **Documentación y comunidad**: Es importante que la herramienta tenga una base sólida de documentación y soporte comunitario para resolver problemas y extender funcionalidades.  
- **Minimización de deuda técnica**: Cada funcionalidad debe estar diseñada para evitar configuraciones complicadas y garantizar que las tareas no introduzcan riesgos de mantenimiento a largo plazo.

## **Análisis de opciones disponibles**

### **Make**  
Make es una herramienta ampliamente utilizada para la automatización de tareas, soportada de forma robusta en muchos lenguajes y entornos. Su uso en proyectos Go puede ser más complicado debido a la necesidad de una estructura más rigurosa y la gestión explícita de dependencias.

- **Ejemplo de uso**:  
  Archivo `Makefile`:  
  ```makefile
  build:
      go build -o bin/main .
  
  clean:
      rm -rf bin

  test:
      go test ./...
  ```
  Comandos:  
  - `make build`: Compila el código.  
  - `make clean`: Limpia el directorio `bin`.  
  - `make test`: Ejecutar tests.   

[Documentación oficial de Make](https://github.com/wkusnierczyk/make?tab=readme-ov-file)

### **Just**  
Just ofrece una alternativa más sencilla y enfocada, ideal para proyectos Go. Su diseño simple y directo favorece la definición de tareas sin complicar la estructura ni requerir el manejo explícito de dependencias, lo que lo convierte en una opción más accesible.

- **Ejemplo de uso**:  
  Archivo `justfile`:  
  ```just
  build:
      go build -o bin/main .
  
  clean:
      rm -rf bin
  
  test:
      go test ./...
  ```
  Comandos:  
  - `just build`: Compila el código.  
  - `just clean`: Elimina archivos generados.  
  - `just test`: Ejecuta los tests.   

[Documentación oficial de Just](https://github.com/casey/just)

### **Task**  
Task utiliza un enfoque basado en archivos YAML para definir tareas. Aunque es potente, esta configuración puede ser innecesariamente compleja para proyectos más simples.  

- **Ejemplo de uso**:  
  Archivo `Taskfile.yml`:  
  ```yaml
  version: '3'

  tasks:
    build:
      cmds:
        - go build -o bin/main .
    clean:
      cmds:
        - rm -rf bin
    test:
      cmds:
        - go test ./...
  ```
  Comandos:  
  - `task build`: Compila el código.  
  - `task clean`: Elimina archivos generados.  
  - `task test`: Ejecuta los tests.  

[Documentación oficial de Task](https://taskfile.dev/) 

### **Mage**  
Mage destaca por permitir la definición de tareas directamente en Go, lo que elimina la necesidad de aprender una sintaxis diferente. Sin embargo, su falta de actualizaciones recientes puede ser un inconveniente.  

- **Ejemplo de uso**:  
  Archivo `magefile.go`:  
  ```go
  //go:build mage
  package main

  import (
      "fmt"
      "os"
  )

  func Build() {
      fmt.Println("Building project...")
      os.MkdirAll("bin", os.ModePerm)
  }

  func Clean() {
      fmt.Println("Cleaning project...")
      os.RemoveAll("bin")
  }
  ```
  Comandos:  
  - `mage build`: Compila el proyecto.  
  - `mage clean`: Elimina los archivos generados.  

[Documentación oficial de Mage](https://github.com/magefile/mage) 

### **Sage**  
Inspirado en Mage, Sage permite automatizar tareas dentro de proyectos en Go. Se posiciona como un reemplazo para herramientas tradicionales como Make y Just, pero al ser relativamente nueva, no mantiene un soporte sólido.

- **Ejemplo de uso**:  
  Archivo `sage.go`:  
  ```go
  package main

  import (
      "fmt"
  )

  func Build() {
      fmt.Println("Building...")
  }

  func Clean() {
      fmt.Println("Cleaning...")
  }
  ```
  Comandos:  
  - `sage build`: Construye el proyecto.  
  - `sage clean`: Limpia los archivos generados.  

[Repositorio oficial de Sage](https://github.com/einride/sage)


## **Evaluación de los gestores**

Para seleccionar la herramienta más adecuada en la gestión de tareas en este proyecto, evaluaremos y compararemos cada una de las opciones disponibles con base en los criterios establecidos previamente:

- Para obtener **actualización constante y soporte activo**  
   -  **Just** cumple con este criterio, ya que está en constante actualización y cuenta con una comunidad activa que contribuye a su desarrollo y mantenimiento. Esto garantiza que siempre tendrá soporte y que se puede confiar en su estabilidad. 
   - **Make:** no ofrece soporte activo, lo que puede complicar la integración a largo plazo.  
   - **Task:** tiene un desarrollo más modesto en comparación, lo que puede generar incertidumbre sobre su futuro a largo plazo.  
   - **Mage:** no ha recibido actualizaciones recientes, lo que representa un riesgo de deuda técnica a futuro.  
   - **Sage:** es relativamente nueva y puede no tener el mismo nivel de madurez y soporte que el resto.

- **Promoción de buenas prácticas**

    - **Just:** promueve buenas prácticas al ser sencilla, directa y adaptada a Go, lo que permite a los desarrolladores enfocarse en el código sin complicaciones adicionales.
    - **Make:** no está especialmente optimizado para Go, lo que puede generar dificultades al intentar adaptarlo a proyectos Go, haciendo que no fomente las mejores prácticas de este ecosistema.
    - **Task:** el uso de archivos YAML no está tan alineado con las prácticas comunes en Go.
    - **Mage:** permite definir tareas directamente en código Go, lo que tiene la ventaja de integrarse perfectamente con el lenguaje y aprovechar las herramientas de su ecosistema. Sin embargo, este enfoque puede mezclar la lógica del negocio con las tareas de automatización, lo que no es ideal desde el punto de vista de las buenas prácticas de configuración.
    - **Sage:** utiliza Go para las tareas, pero la implementación es relativamente nueva y puede no ser tan sólida en cuanto a buenas prácticas específicas de Go.

- Respecto a la **facilidad de uso e intuición en la sintaxis**  
   - **Just y Make:** se distingue por su simplicidad. Utilizan un archivo que es claro y fácil de entender, lo que facilita la adopción de las herramientas sin necesidad de aprender una nueva sintaxis compleja.   
   - **Task:** utiliza YAML, lo que agrega complejidad.  
   - **Mage:** requiere escribir código en Go, lo que, si bien es ventajoso para algunos, podría resultar excesivo para quienes prefieren soluciones de configuración más simples.  
   - **Sage:** también usa Go para definir tareas, lo que puede ser una ventaja para proyectos Go, pero añade complejidad innecesaria.
   
- Respecto a la **documentación y comunidad**  
   - **Just:** tiene una buena documentación y una comunidad activa, lo que facilita la solución de problemas y la extensión de la herramienta.  
   - **Make:** tiene una documentación más extensa, pero su comunidad es menos dinámica en términos de soporte para Go.  
   - **Task** y **Mage** tienen comunidades más pequeñas, lo que puede dificultar la resolución rápida de problemas.  
   - **Sage:** al ser más nueva, tiene una comunidad en crecimiento.

- Respecto a la **minimización de la deuda técnica**  
   - **Just y Make** cumplen con este criterio al evitar configuraciones complejas y promover un flujo de trabajo eficiente con una sintaxis simple. Esto ayuda a mantener el código y las configuraciones limpios y fáciles de mantener.  
   - **Sage** y **Mage** no se alinean completamente con este criterio, ya que su sintaxis o enfoque pueden generar configuraciones complicadas que se suman a la deuda técnica, especialmente en proyectos grandes.  
   - **Task:** el uso de archivos YAML puede hacer que el mantenimiento sea más costoso a largo plazo si no se gestionan adecuadamente.

### Elección y justificación

Después de evaluar las opciones, **Just** se integra perfectamente con el ecosistema Go y proporciona una solución confiable y fácil de mantener. Por lo tanto, esta herramienta es la más adecuada para el proyecto. Al cumplir con todos los criterios establecidos, permite mantener el proyecto organizado, reducir la deuda técnica y automatizar tareas esenciales de manera efectiva. 
