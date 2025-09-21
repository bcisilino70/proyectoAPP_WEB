# proyectoAPP_WEB

# NIVEL DE DATOS

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