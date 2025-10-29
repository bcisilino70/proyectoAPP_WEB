# proyectoAPP_WEB

# NIVEL DE DATOS

## Ejecucion con MAKEFILE

* Para el correcto uso de los Make sugerimos hacer un **make all**, lo que dejara la pagina web activa y creara el contenedor con la base de datos. 
* Luego ejecutar **solo una vez** el **make hurl-req** para crear entidades y verificar que funcione (para poder ejecutarlo de nuevo se deben borrar las filas de las tablas Cliente y Resena con **make db-clean-tablas**). 
* En tercer lugar probar los delete con **make hurl-req-del**. 
* Posterior al uso de los HURL, recomendamos ir a la pagina web y utilizarla donde se veran las entidades creadas por los HURL. La pagina mostrara la opcion de registarse donde puede inventar un usuario y luego accedera al panel de usuario con las funcionalidades. 
* En cualquier momento se puede utilizar **MAKE HURL-CLI:** para listar todos los clientes y **MAKE HURL-RES:** para listar todas las resenas.

 1. **MAKE ALL:** Ejecuta en orden: make destroy (Borra el contenedor y la base de datos) -> make setup (Crea contenedor, otorga privilegios, ejecuta el schema) -> make run_logica (Ejecuta la capa de logica de negocio). 
 
    **Resultado por consola:** Levanta el servidor y se puede abrir en el navegador la pagina web. 

2. **MAKE HURL-REQ:** Crea entidades de clientes y resenas

3. **MAKE HURL-REQ-DEL:** Elimina entidades, lo decidimos separar para que no quede en una sola ejecucion y la salida por consola sea muy larga. 

4. **MAKE DB-CLEAN-TABLAS:** Elimina la informacion de las tablas ya que si por ejemplo se ejecutara dos veces seguidas make hurl-req ocurriria un error en la base de datos por querer crear una entidad que ya esta creada. 

5. **MAKE HURL-CLI:** Lista todos los clientes

6. **MAKE HURL-RES:** Lista todas las resenas



### Dudas sobre queries.sql

* Solo se dejan las consultas con todas las restricciones que se necesiten, por ejemplo una duda era que pasa con Update Resena donde uno como cliente solo da el titulo de la resena pero la base de datos debe buscar ( con el where ) el ID_CLIENTE correspondiente al cliente que quiere cambiar la resena. Como se busca este id_cliente? 
    * Autenticación y autorización: Implementa un sistema de autenticación que permita identificar al cliente y obtener su id en el backend.
    * ¿Cómo se relaciona la autenticación con las consultas SQL? Se ocupa el backend pero en GO 
        
        Autenticación:

        El sistema de autenticación (por ejemplo, manejo de sesiones o tokens JWT) se encarga de identificar al cliente cuando inicia sesión.
        Una vez autenticado, el backend puede obtener el cliente_id del usuario autenticado y usarlo para las consultas SQL.
        Uso del cliente_id en las consultas:

        En tus consultas SQL (queries.sql), el cliente_id debe ser un parámetro obligatorio para las operaciones relacionadas con reseñas.
        Sin embargo, el valor de este parámetro no lo proporciona directamente el cliente, sino que el backend lo inyecta automáticamente al ejecutar la consulta.