# Obtener todos los clientes
GET http://localhost:8080/clientes
HTTP/1.1 200

# Crear un cliente
POST http://localhost:8080/clientes
Content-Type: application/json

{
  "nombre": "Fabricio",
  "apellido": "Britos",
  "usuario": "chinocao",
  "pass": "onichan",
  "email": "chinito@mail.com"
}

HTTP/1.1 201

# Actualizar el email de un cliente
PUT http://localhost:8080/clientes
Content-Type: application/json

{
  "email": "nuevo_email@mail.com",
  "usuario": "chinocao",
  "pass": "onichan"
}

HTTP/1.1 200

# Eliminar un cliente por id
DELETE http://localhost:8080/clientes?id=2

HTTP/1.1 204
