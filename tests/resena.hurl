# 1. Loguearse primero para tener sesión
POST http://localhost:8080/login
[FormParams]
usuario: bauto
contrasena: password123
HTTP 303

# 2. Crear una reseña
POST http://localhost:8080/crear-resena
[FormParams]
titulo: Reseña de prueba Hurl
descripcion: Esta es una reseña automática
nota: 5
HTTP 303

# 3. Crear una reseña para eliminar
POST http://localhost:8080/crear-resena
[FormParams]
titulo: Reseña para borrar
descripcion: Esta reseña será eliminada
nota: 1
HTTP 303

# 4. Eliminar la reseña (usa el ID correcto)
POST http://localhost:8080/eliminar-resena
[FormParams]
resena_id: 2
HTTP 303

# 5. Editar email del usuario (Perfil)
POST http://localhost:8080/actualizar-email
[FormParams]
email: nuevo_email@test.com
HTTP 303