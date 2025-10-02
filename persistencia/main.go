package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	sqlc "persistencia/db/sqlc"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "host=localhost port=5433 user=app_user password=app_pass dbname=app_db sslmode=disable" // A tratar el tema de seguridad.
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	defer db.Close()

	queries := sqlc.New(db)
	ctx := context.Background()

	// --- CREAR USUARIO
	crearUsuario, err := queries.CreateCliente(ctx, // Create
		sqlc.CreateClienteParams{
			Nombre:   "Tomas",
			Apellido: "Di Carlo",
			Usuario:  "tomasdicarlo14",
			Pass:     "123",
			Email:    "tomas@gmail.com",
		})
	if err != nil {
		log.Fatalf("Error al crear usuario: %v", err)
	}
	fmt.Printf("Creado correctamente: %+v\n", crearUsuario)

	// --- GET CLIENTE
	usuario, err := queries.GetCliente(ctx, crearUsuario.ID) // Read One
	if err != nil {
		log.Fatalf("Error al buscar usuario: %v", err)
	}
	fmt.Printf("Usuario encontrado: %+v\n", usuario)

	// --- UPDATE CLIENTE
	err = queries.UpdateCliente(ctx, sqlc.UpdateClienteParams{ // Update
		Email:   "johnny.doe@example.com",
		Usuario: crearUsuario.Usuario,
		Pass:    crearUsuario.Pass,
	})
	if err != nil {
		log.Fatalf("Error al modificar cliente: %v", err)
	}

	// --- CREAR RESENA
	crearResena, err := queries.CreateResena(ctx,
		sqlc.CreateResenaParams{
			Titulo:      "Centinela",
			Descripcion: "El mejor lugar de tandil",
			Nota:        2,
			ClienteID:   crearUsuario.ID,
		})
	if err != nil {
		log.Fatalf("Error al crear resena: %v", err)
	}
	fmt.Printf("Creado correctamente: %+v\n", crearResena)

	// --- GET RESENAS
	resenas, err := queries.ListResenas(ctx, crearUsuario.ID)
	if err != nil {
		log.Fatalf("Error al crear resena: %v", err)
	}
	fmt.Printf("Resenas de un usuario: %+v\n", resenas)

}
