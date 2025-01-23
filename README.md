# PersonalSportCalendary

## Problema principal

Madres, padres y personas adultas con una alta carga de trabajo diario se enfrentan al desafío de equilibrar sus múltiples responsabilidades, como el trabajo, el cuidado de sus hijos y los desplazamientos. Conciliar todo con una buena salud es un verdadero reto al que se enfrentan, es por eso que, muchas de estas personas no logran conseguirlo.

Los largos desplazamientos hasta un gimnasio y las altas esperas para el uso de las maquinas especificas ocasionan una perdida de tiempo enorme del cual no disponen.

Estos adultos necesitan una solucion a este problema, es decir, como conciliar sus vidas personales con un mantenimiento fisico adecuado y suficiente para mantener buen estado de salud, es por ello que, necesitan emplear el menor tiempo posible (segun el que dispongan cada dia) sin emplear tiempos ociosos de desplazamientos innecesarios, **llevándolo a cabo desde casa**, con o sin materiales deportivos, pero adaptado a su disponibilidad diaria.

Es importante entender que, el problema es fundamentalmente el **mantenerse activos fisicamente para una correcta salud**, sin requerimientos especificos en grupos musculares, pero siempre adaptado al tiempo que disponen.

> 💡 **Información adicional:**
>
> Puedes encontrar más información en el archivo [noticias](./docs/news.md) sobre las diferentes consecuencias y riesgos para la salud que conlleva el no realizar suficiente actividad fisica.

## Configuración del repositorio

En [configuración](./docs/config-repo.md) se detalla paso a paso lo que se ha requerido realizar para la correcta configuracion de las claves públicas y privadas y su posterior subida a GitHub (la pública). Además de los comandos necesarios para establecer correctamente mi nombre de usuario y correo.

Finalmente la [clave publica](./docs/images/ClavePublica.png), el [nombre y correo](./docs/images/NombreyCorreo.png) se puede ver que se ha configurado correctamente en GitHub.

## Definición de HUs, milestones y user-journey:

- [Historias de usuario](./docs/HUs.md)
- [Milestones](./docs/config-milestones.md)
- [user-journey](./docs/user-journey.md)

## Selección del gestor de dependencias

El gestor de dependencias escogido es `go.mod` para una información mucho más completa y detallada puede leer el siguiente [documento](./docs/gestor_dependencias.md).

## Selección del gestor de tareas.

En el siguiente [archivo](./docs/gestor_tareas.md) puede encontrar una amplia explicación de cada uno de los gestores de tareas para un proyecto en go, además de un riguroso estudio para la elección de la más apropiada para este proyecto.

De las diferentes tareas una de la más necesarias es la comprobación de la sintaxis:

```
just check
```

## [Biblioteca de aserciones](./docs/biblioteca_asersiones.md)

## [Test runner](./docs/test_runner.md)

Para ejecutar los tets:

```
just test
```

## Docker
La imagen a usar se puede encontrar explicada en el [archivo](/docs/imagen-docker.md).

1. Ejecución y montaje local:
```
docker build -t [nombre] .
```
```
docker run -t -u 1001 -v `pwd`:/app/test [nombre]
```

2. Ejecución Docker:
```
docker run -t -u 1001 -v `pwd`:/app/test mmnuria/personalsportcalendary:latest

```

## Información de la licencia:

[Licencia](./LICENSE)