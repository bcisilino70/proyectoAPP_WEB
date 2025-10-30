# proyectoAPP_WEB

## Prerrequisitos

Antes de ejecutar el proyecto asegurate de tener instalado:

- Docker (Engine y Docker Compose)
- Make
- Hurl (para ejecutar los requests de prueba)
- Git (solo para la opción de clonar el repositorio)


## Opción A: Clonar el repositorio y ejecutarlo

1) Clonar el repo y entrar a la carpeta del proyecto:
```bash
git clone <URL_DEL_REPOSITORIO>
cd <NOMBRE_DE_LA_CARPETA_DEL_PROYECTO>
```

2) Seguí la sección "Ejecucion con MAKEFILE" (más abajo):


## Opción B: Ejecutar desde un ZIP del proyecto

1) Descomprimir el ZIP en una carpeta local y entrar a la carpeta:
```bash
unzip <NOMBRE_DEL_ARCHIVO>.zip -d <CARPETA_DESTINO>
cd <CARPETA_DESTINO>/<NOMBRE_DE_LA_CARPETA_DEL_PROYECTO>
```

2) Seguí la sección "Ejecucion con MAKEFILE" (más abajo):


## Ejecucion con MAKEFILE

* Primer paso : ejecutar **"make all"** en la terminal de comandos.

    Lo que dejara la pagina web activa y creara el contenedor con la base de datos. Importante dejar la terminal abierta para mantener el servidor activo

* Luego ejecutar **solo una vez y en otra terminal** el comando **"make hurl-req"** para crear las entidades cliente y resena de ejemplo.
    
    para poder ejecutarlo de nuevo se deben borrar las filas de las tablas Cliente y Resena con **"make db-clean-tablas"**

* En tercer lugar probar los delete con **"make hurl-req-del"**. 

---TP4---

* Posterior al uso de los HURL, y con el servidor ya activo **"make all"**, para acceder al front-end se debe acceder al siguiente link en un navegador web : http://localhost:8080/ 
La pagina mostrara la opcion de registarse donde puede inventar un usuario y luego acceder al panel de usuario con las funcionalidades.

* En cualquier momento se puede utilizar **"make hurl-cli"** para listar todos los clientes y **"make hurl-res"** para listar todas las resenas.

---algunos comentarios de que ejecuta cada make---

1. **MAKE ALL:** Ejecuta en orden: make destroy (Borra el contenedor y la base de datos) -> make setup (Crea contenedor, otorga privilegios, ejecuta el schema) -> make run_logica (Ejecuta la capa de logica de negocio). 
 
    **Resultado por consola:** Levanta el servidor y se puede abrir en el navegador la pagina web. 

2. **MAKE HURL-REQ:** Crea entidades de clientes y resenas

3. **MAKE HURL-REQ-DEL:** Elimina entidades, lo decidimos separar para que no quede en una sola ejecucion y la salida por consola sea muy larga. 

4. **MAKE DB-CLEAN-TABLAS:** Elimina la informacion de las tablas ya que si por ejemplo se ejecutara dos veces seguidas make hurl-req ocurriria un error en la base de datos por querer crear una entidad que ya esta creada. 

5. **MAKE HURL-CLI:** Lista todos los clientes

6. **MAKE HURL-RES:** Lista todas las resenas

## Notas y solucion de problemas

- Si el puerto 8080 ya está en uso, cerrá el proceso que lo esté usando o cambiá la configuración del puerto si el proyecto lo permite.
- Si quedaron contenedores o datos previos y hay errores al iniciar, probá:
```bash
make destroy
make all
```
- Para volver a correr los Hurl de creación sin errores por duplicados:
```bash
make db-clean-tablas
make hurl-req
```
- Verificá que Docker esté corriendo antes de ejecutar `make all`.
