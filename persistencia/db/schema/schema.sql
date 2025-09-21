-- Created by Vertabelo (http://vertabelo.com)
-- Last modification date: 2025-09-18 13:57:38.503

-- tables
-- Table: CLIENTE
CREATE TABLE CLIENTE (
    id SERIAL PRIMARY KEY,
    nombre varchar(20)  NOT NULL,
    apellido varchar(20)  NOT NULL,
    usuario varchar(20) NOT NULL,
    pass varchar(20) NOT NULL,
    email varchar(50) NOT NULL UNIQUE
);

-- Table: RESENA
CREATE TABLE RESENA (
    id SERIAL PRIMARY KEY,
    titulo VARCHAR(25) NOT NULL UNIQUE,
    descripcion VARCHAR(50) NOT NULL,
    nota INT NOT NULL,
    fecha TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    cliente_id INT NOT NULL
);

ALTER TABLE RESENA ADD CONSTRAINT RESENA_CLIENTE
    FOREIGN KEY (cliente_id)
    REFERENCES CLIENTE (id)
    NOT DEFERRABLE
    INITIALLY IMMEDIATE

-- End of file.