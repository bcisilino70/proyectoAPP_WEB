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
