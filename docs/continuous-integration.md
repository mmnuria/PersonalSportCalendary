## Requisitos de elección 

1. La herramienta debe de ser gratis o al menos porporcionar algunas opciones gratuitas.
2. Debe considerarse (en caso de que lo hubiera) que la herramienta sea nativa con GitHub. 

## Opciones disponibles 
**Circle CI**:    
    Plan gratuito con un límite de 6000 minutos mensuales. Se conecta al repositorio de GitHub mediante una cuenta en su plataforma.
    [Documentación oficial](https://circleci.com/)

**Semaphore CI**:    
    Plan gratuito con un límite de 7000 minutos mensuales. Integración disponible mediante inicio de sesión con GitHub.
    [Documentación oficial](https://semaphoreci.com/)   
    
**GitHub Actions**:    
    Gratuita para repositorios públicos con límite gratuito de 2000 minutos mensuales para repositorios privados.
    [Documentación oficial](https://github.com/features/actions)

**Travis CI**:
    Gratuito para las primeras 100 builds, luego requiere un pago mínimo de $13.75 al mes. 
    [Documentación oficial](https://www.travis-ci.com/)  

**Jenkins**:    
    Software de automatización de código abierto, completamente gratuito. Aunque no es nativo de GitHub, ofrece integraciones mediante *webhooks* y plugins para conectar con GitHub. Requiere configuración en un servidor propio o en un servicio de alojamiento en la nube.  
    [Documentación oficial](https://www.jenkins.io/doc/)

**Appveyor**:    
    Gratuito para proyectos de código abierto con posibilidad de pago $29 al mes para repositorios privados. Integración con GitHub disponible mediante configuración en su plataforma. 
    [Documentación oficial](https://www.appveyor.com/)

## Posibles opciones 

Las herramientas que cumplen con los requisitos establecidos son: **GitHub Actions**, **CircleCI**, **Semaphore CI** y **Appveyor**.  

Para configurar algunas pruebas de las herramientas anteriores, se han tenido que realizar las siguientes configuraciones:

1. [**GitHub Actions**](./images/prueba-github-actions.png)  
   - Crear un archivo en el directorio `.github/workflows/` llamado `continuous-integration.yml`.  
   - Configurar pruebas para Go 1.23 o versiones superiores, acorde a los requisitos del proyecto.  

2. [**CircleCI**](./images/prueba-circleci.png)  
   - Crear un archivo `config.yml` en un nuevo directorio llamado `.circleci`.  

3. [**Semaphore CI**](./images/prueba-semaphore.png)  
   - Crear un archivo `semaphore.yml` en un nuevo directorio llamado `.semaphore`.

4. **Appveyor**  
   -  Crear un archivo `appveyor.yml` en la raíz del proyecto.  

### Conclusión

Se elige **GitHub Actions** como la herramienta principal, ya que:  
- Es la solución nativa de GitHub.  
- Cumple con el criterio inicial de ser gratuita.