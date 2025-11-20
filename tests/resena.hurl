# 1. Loguearse primero para tener sesion
POST http://localhost:8080/login
[FormParams]
usuario: juanro
contrasena: password123
HTTP 303

# 2. Crear una resena
POST http://localhost:8080/crear-resena
[FormParams]
titulo: Resena de prueba Hurl
descripcion: Esta es una resena automatica
nota: 5
HTTP 303

# 3. Crear una resena para eliminar
POST http://localhost:8080/crear-resena
[FormParams]
titulo: Resena para borrar
descripcion: Esta reseña sera eliminada
nota: 1
HTTP 303

# 4. Eliminar la resena (usa el ID correcto)
POST http://localhost:8080/eliminar-resena
[FormParams]
resena_id: 2
HTTP 303

# 5. Editar email del usuario (Perfil)
POST http://localhost:8080/actualizar-email
[FormParams]
email: nuevo_email@test.com
HTTP 303