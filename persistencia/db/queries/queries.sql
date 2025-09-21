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
-- Pensar como se coloca el ID del cliente que hace la resena. No parece ser un problema para las queries. 
INSERT INTO RESENA (titulo,descripcion,nota,fecha)
VALUES ($1,$2,$3,$4,$5,$6)
RETURNING id,titulo,descripcion,nota,fecha,cliente_id

-- name UpdateResena :exec
UPDATE RESENA                           
SET titulo = $1, descripcion = $2
WHERE (titulo = $3) AND (cliente_id = $4)    -- El cliente puede actualizar una resena pero no sabe el id de la misma. A chequear

-- Consulta para listar TODAS resenas de un cliente.
-- name ListResenas :many
SELECT titulo, descripcion, nota, fecha 
FROM RESENA
WHERE (cliente_id = $1)

-- Consulta para listar UNA resena de un cliente.
-- name ListResena :one
SELECT titulo, descripcion, nota, fecha 
FROM RESENA
WHERE (cliente_id = $1) and (titulo = $1)

-- Consulta para borrar una resena. 
-- name DeleteResena :exec
DELETE FROM RESENA 
WHERE (titulo = $1) 

-- Consulta para borrar un cliente
-- name DeleteCliente :exec
DELETE FROM CLIENTE
WHERE (usuario = $1) 

