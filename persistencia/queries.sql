-- name CreateCliente :one
INSERT INTO CLIENTE (nombre,apellido,email)
VALUES ($1,$2,$3)
RETURNING id,nombre,apellido,email;

-- name UpdateCliente :excec
UPDATE CLIENTE
SET email = $1

