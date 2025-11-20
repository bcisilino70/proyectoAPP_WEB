# proyectoAPP_WEB

## Ejecución con Docker Compose

Ahora el proyecto está dockerizado completo (App + Base de datos), así que es mucho más fácil de levantar, `docker-compose` se encarga de todo.

### Comandos Principales

1. **`make up`**:
   Este es el comando principal. Levanta todo el entorno (la base de datos Postgres y la app en Go) usando Docker Compose.
   *   Si es la primera vez, va a construir la imagen y configurar la base de datos automáticamente con el `schema.sql`.
   *   **Resultado:** El servidor queda corriendo en segundo plano y puedes entrar a la web en `http://localhost:8080`.

2. **`make down`**:
   Baja y apaga todos los contenedores. IMPORTANTE : Usalo cuando termines de trabajar para liberar recursos.

3. **`make test`**:
   

4. **`make db-data`**:
   Comando útil para ver rápido qué datos hay cargados en las tablas `cliente` y `resena` sin tener que entrar a la base manualmente.


### Cómo probarlo
Simplemente ejecutar `make up`, esperar unos segundos a que levante todo, luego entrar al navegador a probar la app (registrarse, crear reseñas, etc). Si se requiere ver que se guardaron bien en la base de datos, ejecutar un `make db-data`.

