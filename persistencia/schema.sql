-- Created by Vertabelo (http://vertabelo.com)
-- Last modification date: 2025-09-18 13:57:38.503

-- tables
-- Table: CLIENTE
CREATE TABLE CLIENTE (
    id int  NOT NULL,
    nombre varchar(20)  NOT NULL,
    apellido varchar(20)  NOT NULL,
    CONSTRAINT CLIENTE_pk PRIMARY KEY (id)
);

-- Table: RESENA
CREATE TABLE RESENA (
    id int  NOT NULL,
    titulo varchar(25)  NOT NULL,
    descripcion varchar(50)  NOT NULL,
    nota int  NOT NULL,
    fecha timestamp  NOT NULL,
    CLIENTE_id int  NOT NULL,
    CONSTRAINT RESENA_pk PRIMARY KEY (id,CLIENTE_id)
);

-- foreign keys
-- Reference: RESENA_CLIENTE (table: RESENA)
ALTER TABLE RESENA ADD CONSTRAINT RESENA_CLIENTE
    FOREIGN KEY (CLIENTE_id)
    REFERENCES CLIENTE (id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- End of file.