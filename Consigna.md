## Problematica

Una compañia que se dedica a organizar eventos, desea validar su negocio en Internet.
Por ello, decide contratarte para realizar un MVP que resuelva las siguientes necesidades:

- La aplicacion debe ser WEB y contar con API para la interaccion con el frontend.
  Nota: tanto el backend (modelo y reglas de negocio) como la API debera desarrollarse preferentemente en Golang. No se especifica ningun framework en particular.
- Los eventos cuientan con titulo, descripcion corta, descripcion larga, fecha y hora, organizador, lugar y un estado (borrador o publicada)
- Cuando una publicacion tiene estado borrador, solo los administradores pueden visualizarla.
- Los usuarios pueden visualizar e inscribirse en eventos publicados, siempre y cuando la fecha del evento sea futura.
- Los eventos cuya fecha y hora hayan transcurrido, pueden visualizarse pero no inscribirse.
- Los usuarios (administradores o no) pueden visualizar los eventos activos
  (publicados y con fecha-hora en el futuro) como así los eventos completados
  (publicados pero con fecha y hora en el pasado). Este listado de eventos se debe
  poder filtrar por fecha, estado, y título.
- A su vez, se contará con un endpoint en el cual se mostrarán todos los eventos
  inscriptos, filtrándose por activos o completados.

## Consideraciones Generales

1. Podes utilizar algún framework o combinación de ellos para resolver el problema,
   como así paquetes o librerías Golang que solucionen problemas puntuales
2. Se espera poder reproducir el entorno de desarrollo fácilmente (gestión de
   dependencias mediante un requirements.txt, poetry, docker, etc)
3. Estilo de código sujeto a PEP8.
4. Como entregable se espera:

- Un repositorio git con el desarrollo, que evidencia su utilizacion durante el desarrollo.
- Test automaticos (unitarios y/o integracion)
- Un alto porcentaje de los requisitos cubiertos, especificando requisitos no alcanzados por el MVP.

## Opcional

De manera opcional, puede desarrollarse el frontend que interactúe con la API
desarrollada. Esta aplicación web deberá ser desarrollada preferentemente en React.

En esta web, podra:

- Ver el listado de eventos general, con los filtros disponibles.
- Visualizar un evento de forma detallada.
- Inscribirse al evento
- Lista de eventos suscriptos

Sientase libre de añadir otras funcionalidades si lo ve necesario.

## Entrega

La fecha de entrega esta establecida para 7 dias despues de enviado el ejercicio practico

Par consultas podes comunicarete a: ......@....app
