package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq" // Driver de PostgreSQL
	// Asegúrate de que la ruta de importación coincida con el nombre de tu módulo en go.mod
	// seguido de la ruta al paquete sqlc.
	// Por ejemplo: "proyectoAPP_WEB/persistencia/db/sqlc"
	sqlc "proyectoAPP_WEB/persistencia/db/sqlc"
)

func main() {
	// Reemplaza esta cadena con los datos de tu conexión a PostgreSQL
	connStr := "host=localhost port=5432 user=app_user password=app_pass dbname=app_db sslmode=disable"

	// 1. Conexión con la Base de Datos
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("No se pudo conectar a la base de datos: %v", err)
	}
	defer db.Close()

	// Ping a la base de datos para ver si funciona
	if err := db.Ping(); err != nil {
		log.Fatalf("No se pudo hacer ping a la base de datos: %v", err)
	}

	log.Println("Conexión a la base de datos establecida exitosamente.")

	// Crea una instancia del repositorio generado por sqlc
	queries := sqlc.New(db)

	// Algo para que queries no tire error de variable no usada
	_ = queries

	// 2. Configuración del Servidor HTTP
	/*
		POST /<entidades>: Debe recibir datos en formato JSON, validarlos (ej. que el título no esté vacío), usar el
		método Create... de sqlc para guardarlos en la base de datos y devolver el nuevo objeto como JSON con
		estado 201.
	*/

	/*
		GET /<entidades>: Debe usar el método List... de sqlc para obtener todos los registros y devolverlos
		como un array JSON.
	*/

	/*
		GET /<entidades>/{id}: Debe obtener el ID de la URL, buscar el registro con Get... y devolverlo. Si no
		existe, debe devolver un 404.
	*/

	/*
		PUT /<entidades>/{id}: Debe recibir datos JSON, validarlos, y actualizar el registro correspondiente usando
		Update....
	*/

	/*
		DELETE /<entidades>/{id}: Debe eliminar el registro usando Delete... y devolver un estado 204 (No
		Content).
	*/
	log.Println("Iniciando servidor en http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("No se pudo iniciar el servidor: %v", err)
	}
}
