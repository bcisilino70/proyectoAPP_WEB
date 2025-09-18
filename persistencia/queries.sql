-- name CreateCliente :one
INSERT INTO CLIENTE (nombre,apellido,usuario,pass,email)
VALUES ($1,$2,$3,$4,$5,$6)
RETURNING id,nombre,apellido,usuario,pass,email;

-- name UpdateCliente :exec
-- Permite al cliente cambiar el email. 
UPDATE CLIENTE
SET email = $1
WHERE (usuario = $2) AND (pass = $3)

-- name CreateResena :one
-- Pensar como se coloca el ID del cliente que hace la resena. 
INSERT INTO RESENA (titulo,descripcion,nota,fecha)
VALUES ($1,$2,$3,$4,$5,$6)
RETURNING id,titulo,descripcion,nota,fecha,cliente_id

-- name UpdateResena :exec
UPDATE RESENA                           -- A Chequearrrr
SET (titulo = $1), (descripcion = $2)
WHERE (id = $3) AND (cliente_id = $4)

-- Consulta para listar todas resenas de un cliente.

-- Consulta para listar una resena de un cliente.

-- Consulta para borrar una resena. 


