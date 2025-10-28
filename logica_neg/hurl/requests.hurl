# ============================================
# PRUEBAS COMPLETAS DE LA API - CLIENTES Y RESEÑAS
# ============================================

# 0. Verificar que el servidor está funcionando
GET http://localhost:8080/
HTTP/1.1 200
[Asserts]
header "Content-Type" contains "text/html"
body contains "<!DOCTYPE html>"

# ============================================
# SECCIÓN 1: CLIENTES
# ============================================

# 1. Crear Cliente #1
POST http://localhost:8080/clientes
Content-Type: application/json

{
  "nombre": "Juan",
  "apellido": "Pérez",
  "usuario": "juanperez",
  "pass": "password123",
  "email": "juan.perez@mail.com"
}
HTTP/1.1 201

# 2. Crear Cliente #2
POST http://localhost:8080/clientes
Content-Type: application/json

{
  "nombre": "María",
  "apellido": "García",
  "usuario": "mariagarcia",
  "pass": "password456",
  "email": "maria.garcia@mail.com"
}
HTTP/1.1 201

# 3. Crear Cliente #3
POST http://localhost:8080/clientes
Content-Type: application/json

{
  "nombre": "Carlos",
  "apellido": "López",
  "usuario": "carloslopez",
  "pass": "password789",
  "email": "carlos.lopez@mail.com"
}
HTTP/1.1 201

# 4. Listar todos los clientes
GET http://localhost:8080/clientes
HTTP/1.1 200
[Asserts]
header "Content-Type" contains "application/json"

# 5. Actualizar email del Cliente #1
PUT http://localhost:8080/clientes?id=1
Content-Type: application/json

{
  "email": "juan.perez.nuevo@mail.com"
}
HTTP/1.1 200

# 6. Verificar que se actualizó listando todos los clientes
GET http://localhost:8080/clientes
HTTP/1.1 200

# ============================================
# SECCIÓN 2: RESEÑAS
# ============================================

# 7. Crear Reseña #1 para Cliente #1
POST http://localhost:8080/resenas?cliente_id=1
Content-Type: application/json

{
  "titulo": "Plaza San Martín",
  "descripcion": "Hermosa plaza en el centro de la ciudad",
  "nota": 5
}
HTTP/1.1 201

# 8. Crear Reseña #2 para Cliente #1
POST http://localhost:8080/resenas?cliente_id=1
Content-Type: application/json

{
  "titulo": "Restaurante El Buen Sabor",
  "descripcion": "Comida excelente el servicio es un poco lento",
  "nota": 4
}
HTTP/1.1 201

# 9. Crear Reseña #3 para Cliente #2
POST http://localhost:8080/resenas?cliente_id=2
Content-Type: application/json

{
  "titulo": "Cine Central",
  "descripcion": "Buenas instalaciones y pantallas de calidad",
  "nota": 4
}
HTTP/1.1 201

# 10. Crear Reseña #4 para Cliente #2
POST http://localhost:8080/resenas?cliente_id=2
Content-Type: application/json

{
  "titulo": "Parque Nacional",
  "descripcion": "Lugar increíble para hacer trekking",
  "nota": 5
}
HTTP/1.1 201

# 11. Listar todas las reseñas del Cliente #1
GET http://localhost:8080/resenas?cliente_id=1
HTTP/1.1 200
[Asserts]
header "Content-Type" contains "application/json"

# 12. Listar todas las reseñas del Cliente #2
GET http://localhost:8080/resenas?cliente_id=2
HTTP/1.1 200
[Asserts]
header "Content-Type" contains "application/json"

# 13. Actualizar Reseña #1 (del Cliente #1)
PUT http://localhost:8080/resenas?id=1&cliente_id=1
Content-Type: application/json

{
  "titulo": "Plaza San Martín",
  "descripcion": "Hermosa plaza recientemente",
  "nota": 5
}
HTTP/1.1 200

# 14. Verificar actualización listando reseñas del Cliente #1
GET http://localhost:8080/resenas?cliente_id=1
HTTP/1.1 200

# 15. Actualizar Reseña #3 (del Cliente #2)
PUT http://localhost:8080/resenas?id=3&cliente_id=2
Content-Type: application/json

{
  "titulo": "Cine Central",
  "descripcion": "Han mejorado mucho",
  "nota": 4
}
HTTP/1.1 200

# 16. Verificar actualización listando reseñas del Cliente #2
GET http://localhost:8080/resenas?cliente_id=2
HTTP/1.1 200

# 17. Listar reseñas recientes de todos los usuarios (máx. 5)
GET http://localhost:8080/resenas/recientes?limit=5
HTTP/1.1 200
[Asserts]
header "Content-Type" contains "application/json"


