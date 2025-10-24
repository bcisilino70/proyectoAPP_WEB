# Obtener todas las resenas de un cliente
GET http://localhost:8080/resenas?cliente_id=1
HTTP/1.1 200

# Crear una resena, asociada a un cliente.
POST http://localhost:8080/resenas?cliente_id=1
Content-Type: application/json

{
    "titulo" : "Plaza del tanque",
    "descripcion" : "Un lugar tranquilo",
    "nota" : 5
}

HTTP/1.1 201


# Actualizar la resena


# Eliminar una resena


