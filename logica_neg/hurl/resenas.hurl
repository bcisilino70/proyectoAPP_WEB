# Obtener todas las reseñas de un cliente
GET http://localhost:8080/resenas?cliente_id=1
HTTP/1.1 200

# Crear una reseña asociada a un cliente
POST http://localhost:8080/resenas?cliente_id=1
Content-Type: application/json

{
    "titulo" : "Plaza del tanque",
    "descripcion" : "Un lugar tranquilo",
    "nota" : 5
}

HTTP/1.1 201

# Actualizar una reseña existente del cliente
PUT http://localhost:8080/resenas?id=1&cliente_id=1
Content-Type: application/json

{
    "titulo" : "Plaza del tanque",
    "descripcion" : "Sigue siendo un lugar tranquilo",
    "nota" : 4
}

HTTP/1.1 200

# Eliminar una reseña propia
DELETE http://localhost:8080/resenas?id=1&cliente_id=1
HTTP/1.1 204

# Listar las reseñas más recientes de todos los usuarios (máximo 10)
GET http://localhost:8080/resenas/recientes?limit=10
HTTP/1.1 200


