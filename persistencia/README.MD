# Persistencia de datos. 

### Desarrollo 
* Organizamos la estructura de carpetas, creando una carpeta db.
    * Dentro de la misma segmentamos queries, schema y sqlc
    * queries.sql -> desarrollamos las consultas con el formato de comentarios que requiere sqlc
    * schema.sql -> con la herramienta Vertabelo realizamos un esquema Cliente - Resena y trajimos el sql generado. 
    * sqlc -> fue generada por el comando sqlc generate y nosotros decidimos en el sqlc.yaml que sea esta la ubicacion. 

* Creamos un contenedor docker para poder correr postgresql y crear la base de datos.
* Luego mediante el main.go pudimos conectarnos a la base de datos y desarrollar en go algunas funcionalidades como crear usuario, crear resena. Mediante las queries sql en go que genera el sqlc generate pudimos crear entradas para la base de datos.
* Finalmente agregamos manejo de errores para saber si se estaba realizando o no las acciones CRUD en la base de datos. 

### Para probarlo. 

1. Se necesita una base de datos postgresql donde poder ejecutar y generar el schema.sql 
2. Modificar el connStr poniendo usuario contrasena y base de datos propias. 
3. Luego correr go run main.go
4. La respuesta esperada por consola es que se crearon correctamente el usuario y resena de prueba. 
    * Creado correctamente: {ID:1 Nombre:Tomas Apellido:Di Carlo Usuario:tomasdicarlo14 Pass:123 Email:tomas@gmail.com}
    * Usuario encontrado: {Nombre:Tomas Apellido:Di Carlo Usuario:tomasdicarlo14 Email:tomas@gmail.com}
    * Creado correctamente: {ID:1 Titulo:Centinela Descripcion:El mejor lugar de tandil Nota:2 Fecha:2025-10-02 16:12:31.657275 +0000 +0000 ClienteID:1}
    * Resenas de un usuario: [{Titulo:Centinela Descripcion:El mejor lugar de tandil Nota:2 Fecha:2025-10-02 16:12:31.657275 +0000 +0000}]

### Ruta de comandos.

* Datos:
    * Nombre del contenedor: app_postgres
    * Base de datos: app_db
    * Usuario: app_user
    * Contraseña: app_pass

1. Crear el contenedor de PostgresSQL: 
    * docker run --name app_postgres -e POSTGRES_DB=app_db -e POSTGRES_USER=app_user -e POSTGRES_PASSWORD=app_pass -p 5433:5432 -d postgres:latest
    * Elegir un puerto que no hayas usado antes
2. Acceder al contenedor: 
    * docker exec -it app_postgres psql -U app_user -d app_db
3. Una vez dentro del contenedor, creamos un usuario de postgres y otorgamos privilegios. 
    * CREATE USER app_admin WITH PASSWORD 'admin_pass';
    * GRANT ALL PRIVILEGES ON DATABASE app_db TO app_admin;
    * \q  Salimos del contenedor. 
4. Desde la terminal dentro de persistencia podemos ejecutar este comando para ejecutar el schemas.sql
    * docker exec -i app_postgres psql -U app_user -d app_db < db/schema/schema.sql
    * Salida esperada   CREATE TABLE
                        CREATE TABLE
                        ALTER TABLE 
5. Finalmente desde persistencia en la terminal se puede correr el main. 
    * go run main.go
    * Detalles: en la linea 14 del main se encuentran los datos para entrar a la base de datos y poder ejecutar el main.go. De ser necesario deben ser adaptador al contenedor que se haya decidido crear 
    * Salida esperada: 

    Creado correctamente: {ID:1 Nombre:Tomas Apellido:Di Carlo Usuario:tomasdicarlo14 Pass:123 Email:tomas@gmail.com}
    
    Usuario encontrado: {Nombre:Tomas Apellido:Di Carlo Usuario:tomasdicarlo14 Email:tomas@gmail.com}
    
    Creado correctamente: {ID:1 Titulo:Centinela Descripcion:El mejor lugar de tandil Nota:2 Fecha:2025-10-02 17:41:25.368992 +0000 +0000 ClienteID:1}
    
    Resenas de un usuario: [{Titulo:Centinela Descripcion:El mejor lugar de tandil Nota:2 Fecha:2025-10-02 17:41:25.368992 +0000 +0000}]
