# proyectoAPP_WEB

## Ejecución con Docker Compose

El proyecto está completamente dockerizado (App + Base de datos), facilitando su despliegue y pruebas mediante `docker compose` con version Docker Compose version v2.27.1. 

### Comandos Principales

1. **`make all`** (Recomendado para corrección):
   Este comando ejecuta el ciclo completo de prueba:
   1. Detiene cualquier contenedor previo (`make down`).
   2. Levanta el entorno limpio (`make up`).
   3. Espera a que la base de datos esté lista.
   4. Ejecuta los tests de integración con Hurl (`make test`).
   5. Muestra el estado final de la base de datos (`make db-data`).

2. **`make up`**:
   Levanta todo el entorno (Postgres + App Go) en segundo plano.
   *   **Resultado:** El servidor queda corriendo en `http://localhost:8080`.

3. **`make down`**:
   Detiene y elimina los contenedores para liberar recursos.

4. **`make test`**:
   Ejecuta los scripts de prueba de integración (`cliente.hurl` y `resena.hurl`) para validar el funcionamiento de los endpoints.

5. **`make db-data`**:
   Muestra en consola el contenido actual de las tablas `cliente` y `resena` para verificar la persistencia de datos.

### Cómo probarlo

**Opción A (Automática):**
Ejecutar simplemente:
```bash
make all