# Ejercicio Practico - Proceso de seleccion Backend

### Todos las rutas probadas en postman estan en el json

## Dependencias

- go get github.com/gofiber/fiber/v2
- go get github.com/google/uuid
- go get go.mongodb.org/mongo-driver/mongo

## Requisitos no alcanzados

- Autenticacion y Autorizacion con JWT
- Test

## Rutas

- Register:

  - Metodo http post
  - Se pasa por body: name, lastname, email y password

- Login:

  - Metodo http post
  - Se pasa por body: email y password

- Own Event:

  - Metodo http get
  - Se pasa por query un userId y se puede pasar una fecha (con formato dd/mm/aa) para filtrar por fecha (se pasa por query con el nombre time)

- Inscribe event:

  - Metodo http put
  - Se pasa por body un userId y un eventId (Se podria hacer que el title sea unico y utilizar este).

- Get user/s:

  - Metodo http get
  - Si queremos un user especifico podemos pasar por query un name, sino pasamos nada nos devuelve todos los users.

- Get event:

  - Metodo http get
  - Consta con 3 filtros que se pasan por query, si no se pasa ninguno se devuelven todos los eventos existentes.
  - Los filtros disponibles son:
    - Por title, date y state. Puede ser cualquier combinacion entre estos 3.

- Create event:

  - Metodo http post
  - Se pasa por body title, short description, large description, organizer, date, hour, place y state (Por defecto se setea en true pero la linea esta comentada para test con postman).

- Edit event:

  - Metodo http put
  - Se pasa por params un id de evento y por body los datos que quieras cambiar, los datos que no sean enviados se setean por defecto.
