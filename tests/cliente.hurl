# 1. Verificar que la Home carga
GET http://localhost:8080/
HTTP 200

# 2. Registrar un usuario nuevo
POST http://localhost:8080/register
[FormParams]
nombre : Juan
apellido : Rodriguez
usuario : juanro
mail: test@email.com
contrasena: password123
HTTP 303

# 3. Iniciar sesion (Login)
POST http://localhost:8080/login
[FormParams]
usuario: juanro
contrasena: password123
HTTP 303
